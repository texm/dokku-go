package dokku

import (
	"bytes"
	"crypto/rsa"
	"errors"
	"fmt"
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

type Client interface {
	Dial() error
	DialWithTimeout(time.Duration) error
	Close() error

	exec(string) (string, error)

	CloneApp(string, string) error
	CreateApp(string) error
	DestroyApp(string) error
	CheckAppExists(string) (bool, error)
	ListApps() ([]string, error)
	LockApp(string) error
	IsLocked(string) (bool, error)
	RenameApp(string, string) error
	GetAppReport(string) (*AppReport, error)
	GetAllAppReport() (AppsReport, error)
	UnlockApp(string) error

	GetAllProcessReport() (ProcessesReport, error)
	GetProcessInfo(string) error

	SetEventLoggingEnabled(bool) error
	GetEventLogs() (string, error)
	ListLoggedEvents() ([]string, error)
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

func isInvalidAppError(out string) bool {
	return strings.HasSuffix(out, "does not exist")
}

func isNoDeployedAppsError(out string) bool {
	return strings.Contains(out, noAppsDokkuMessage)
}

// TODO: generalise
func checkGenericError(output string) error {
	if isInvalidAppError(output) {
		return InvalidAppError
	}
	if isNoDeployedAppsError(output) {
		return NoDeployedAppsError
	}
	return nil
}

func (c *DefaultClient) exec(cmd string) (string, error) {
	session, err := c.conn.NewSession()
	if err != nil {
		return "", err
	}

	output, cmdErr := session.CombinedOutput(cmd)
	cleaned := strings.TrimSpace(string(output))

	if sessErr := session.Close(); sessErr != nil {
		return cleaned, sessErr
	}

	if err := checkGenericError(cleaned); err != nil {
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

func (c *DefaultClient) Close() error {
	return c.conn.Close()
}
