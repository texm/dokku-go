package dokku

import (
	"bytes"
	"errors"
	"io/ioutil"
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

type ClientConfig struct {
	Host                string
	Port                string
	PrivateKey          []byte
	PrivateKeyFileName  string
	PrivateKeyFilePath  string
	PrivateKeyPassphare string
	HostKeyCallback     ssh.HostKeyCallback
}

type Client struct {
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

	var privateKey []byte
	if len(privateKey) > 0 {
		privateKey = cfg.PrivateKey
	} else if cfg.PrivateKeyFilePath != "" {
		privateKey, err = ioutil.ReadFile(cfg.PrivateKeyFilePath)
		if err != nil {
			return nil, err
		}
	} else if cfg.PrivateKeyFileName != "" {
		filePath := path.Join(homeDir, ".ssh", cfg.PrivateKeyFileName)
		privateKey, err = ioutil.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("no valid private key")
	}
	signer, err := ssh.ParsePrivateKey(privateKey)

	if err != nil {
		return nil, err
	}
	pubKeyAuth := ssh.PublicKeys(signer)

	sshConfig := &ssh.ClientConfig{
		User:            dokkuUser,
		Auth:            []ssh.AuthMethod{pubKeyAuth},
		HostKeyCallback: hostKeyCallback,
		Timeout:         defaultTimeout,
	}

	addr := net.JoinHostPort(cfg.Host, cfg.Port)
	sshConn, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, err
	}

	client := &Client{
		conn: sshConn,
	}

	return client, nil
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
