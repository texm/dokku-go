package dokku

import (
	"bytes"
	"crypto/rsa"
	"errors"
	"net"
	"os"
	"path"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

const (
	dokkuUser      = "dokku"
	defaultTimeout = time.Second * 5
	knownHostsFile = ".ssh/known_hosts"
)

var (
	InvalidPrivateKeyError = errors.New("No valid private key supplied")
)

type ClientConfig struct {
	Host                 string
	Port                 string
	PrivateKey           *rsa.PrivateKey
	PrivateKeyBytes      []byte
	PrivateKeyPassphrase string
	HostKeyCallback      ssh.HostKeyCallback
}

type Client struct {
	cfg    *ClientConfig
	sshCfg *ssh.ClientConfig
	conn   *ssh.Client
	stdout bytes.Buffer
	stderr bytes.Buffer
}

func NewClient(cfg *ClientConfig) (*Client, error) {
	if cfg.Port == "" {
		cfg.Port = "22"
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	hostKeyCallback := cfg.HostKeyCallback
	if hostKeyCallback == nil {
		cb, err := knownhosts.New(path.Join(homeDir, knownHostsFile))
		if err != nil {
			return nil, err
		}
		hostKeyCallback = cb
	}

	var signer ssh.Signer
	if cfg.PrivateKey != nil {
		signer, err = ssh.NewSignerFromKey(cfg.PrivateKey)
		if err != nil {
			return nil, err
		}
	} else if len(cfg.PrivateKeyBytes) > 0 {
		signer, err = ssh.ParsePrivateKey(cfg.PrivateKeyBytes)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, InvalidPrivateKeyError
	}

	sshConfig := &ssh.ClientConfig{
		User: dokkuUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		Timeout:         defaultTimeout,
		HostKeyCallback: hostKeyCallback,
	}

	client := &Client{
		cfg:    cfg,
		sshCfg: sshConfig,
	}

	return client, nil
}

func (c *Client) Dial() error {
	addr := net.JoinHostPort(c.cfg.Host, c.cfg.Port)
	sshConn, err := ssh.Dial("tcp", addr, c.sshCfg)
	if err != nil {
		return err
	}

	c.conn = sshConn

	return nil
}

func (c *Client) DialWithTimeout(timeout time.Duration) error {
	c.sshCfg.Timeout = timeout
	return c.Dial()
}

func (c *Client) exec(cmd string) (string, error) {
	session, err := c.conn.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	cleaned := strings.TrimSpace(string(output))
	if err != nil {
		var exitCodeErr *ssh.ExitError
		if errors.As(err, &exitCodeErr) {
			return cleaned, newDokkuError(cleaned)
		}
		return "", err
	}

	return cleaned, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
