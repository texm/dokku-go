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
	Host string
	// optional, defaults to 22
	Port                 string
	PrivateKey           *rsa.PrivateKey
	PrivateKeyBytes      []byte
	PrivateKeyPassphrase []byte
	// optional
	HostKeyCallback ssh.HostKeyCallback
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
	var signerError error
	if cfg.PrivateKey != nil {
		signer, signerError = ssh.NewSignerFromKey(cfg.PrivateKey)

	} else if len(cfg.PrivateKeyBytes) > 0 {
		if len(cfg.PrivateKeyPassphrase) > 0 {
			signer, signerError = ssh.ParsePrivateKeyWithPassphrase(cfg.PrivateKeyBytes, cfg.PrivateKeyPassphrase)
		} else {
			signer, signerError = ssh.ParsePrivateKey(cfg.PrivateKeyBytes)
		}
	} else {
		return nil, InvalidPrivateKeyError
	}

	if signerError != nil {
		return nil, signerError
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

func isInvalidAppError(out string) bool {
	return strings.HasPrefix(out, "!     App") &&
		strings.HasSuffix(out, "does not exist")
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
		if isInvalidAppError(cleaned) {
			return cleaned, InvalidAppError
		}

		var exitCodeErr *ssh.ExitError
		if errors.As(err, &exitCodeErr) {
			return cleaned, newDokkuError(cleaned)
		}
		return cleaned, err
	}

	return cleaned, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
