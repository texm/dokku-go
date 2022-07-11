package dokku

import (
	"bytes"
	"crypto/rsa"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

type dokkuSSHClient interface {
	Dial() error
	DialWithTimeout(timeout time.Duration) error

	Close() error

	Exec(command string) (string, error)
}

const (
	dokkuUser      = "dokku"
	defaultTimeout = time.Second * 5
	knownHostsFile = ".ssh/known_hosts"
)

var (
	InvalidPrivateKeyError = errors.New("invalid private key")
)

type ClientConfig struct {
	Host string
	// optional, defaults to 22
	Port                 string
	PrivateKey           *rsa.PrivateKey
	PrivateKeyBytes      []byte
	PrivateKeyPassphrase []byte
	// optional, defaults to using $HOME/.ssh/known_hosts
	HostKeyCallback ssh.HostKeyCallback
}

type DefaultClient struct {
	cfg    *ClientConfig
	sshCfg *ssh.ClientConfig
	conn   *ssh.Client
	stdout bytes.Buffer
	stderr bytes.Buffer
}

func NewClient(cfg *ClientConfig) (Client, error) {
	if cfg.Port == "" {
		cfg.Port = "22"
	}

	hostKeyCallback := cfg.HostKeyCallback
	if hostKeyCallback == nil {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}

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

	client := &DefaultClient{
		cfg:    cfg,
		sshCfg: sshConfig,
	}

	return client, nil
}

func (c *DefaultClient) Dial() error {
	addr := net.JoinHostPort(c.cfg.Host, c.cfg.Port)
	sshConn, err := ssh.Dial("tcp", addr, c.sshCfg)
	if err != nil {
		return err
	}

	c.conn = sshConn

	return nil
}

func (c *DefaultClient) DialWithTimeout(timeout time.Duration) error {
	c.sshCfg.Timeout = timeout
	return c.Dial()
}

func checkGenericErrors(output string) error {
	if strings.HasSuffix(output, "does not exist") {
		return InvalidAppError
	}
	if strings.HasSuffix(output, "has not been deployed") {
		return AppNotDeployedError
	}
	if strings.Contains(output, noAppsDokkuMessage) {
		return NoDeployedAppsError
	}
	if strings.Contains(output, nameTakenMessage) {
		return NameTakenError
	}
	return nil
}

func closeSession(session *ssh.Session) error {
	// The session can be closed asynchronously at any time by the server,
	// so it's always possible for correctly-written code to get an EOF error
	// from calling Close() - so we ignore it
	err := session.Close()
	if err.Error() != "EOF" {
		return fmt.Errorf("error closing ssh session: %w", err)
	}
	return nil
}

func (c *DefaultClient) Exec(cmd string) (string, error) {
	session, err := c.conn.NewSession()
	if err != nil {
		return "", err
	}

	output, cmdErr := session.CombinedOutput(cmd)
	cleaned := strings.TrimSpace(string(output))

	if sessErr := closeSession(session); sessErr != nil {
		return cleaned, sessErr
	}

	if err := checkGenericErrors(cleaned); err != nil {
		return cleaned, err
	}

	if cmdErr != nil {
		var exitCodeErr *ssh.ExitError
		if errors.As(cmdErr, &exitCodeErr) {
			return cleaned, fmt.Errorf("dokku error: '%w'", cmdErr)
		}
		return cleaned, err
	}

	return cleaned, nil
}

type CommandStream struct {
	Stdout io.Reader
	Stderr io.Reader
}

func (c *DefaultClient) StreamingExec(cmd string) (*CommandStream, error) {
	session, err := c.conn.NewSession()
	if err != nil {
		return nil, err
	}
	stream := &CommandStream{}

	stream.Stdout, err = session.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stream.Stderr, err = session.StderrPipe()
	if err != nil {
		return nil, err
	}

	go func() {
		_ = session.Run(cmd)
		_ = closeSession(session)
	}()

	return stream, nil
}

func (c *DefaultClient) Close() error {
	return c.conn.Close()
}
