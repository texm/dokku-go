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

type Client interface {
	// client connection methods

	Dial() error
	DialWithTimeout(timeout time.Duration) error
	Close() error

	Exec(command string) (string, error)

	// apps

	CloneApp(currentAppName string, newAppName string) error
	CreateApp(appName string) error
	DestroyApp(appName string) error
	CheckAppExists(appName string) (bool, error)
	ListApps() ([]string, error)
	LockApp(appName string) error
	IsLocked(appName string) (bool, error)
	RenameApp(currentAppName string, newAppName string) error
	GetAppReport(appName string) (*AppReport, error)
	GetAllAppReport() (AppsReport, error)
	UnlockApp(appName string) error

	// ps

	GetProcessInfo(appName string) error
	GetAppProcessReport(appName string) (*ProcessReport, error)
	GetAllProcessReport() (ProcessesReport, error)
	GetAppProcessScale(appName string) (map[string]int, error)
	SetAppProcessScale(appName string, processName string, scale int, skipDeploy bool) error
	StartApp(appName string, p *ParallelismOptions) error
	StartAllApps(p *ParallelismOptions) error
	StopApp(appName string, p *ParallelismOptions) error
	StopAllApps(p *ParallelismOptions) error
	RebuildApp(appName string, p *ParallelismOptions) error
	RebuildAllApps(p *ParallelismOptions) error
	RestartApp(appName string, p *ParallelismOptions) error
	RestartAppProcess(appName string, process string, p *ParallelismOptions) error
	RestartAllApps(p *ParallelismOptions) error
	SetAppProcessProperty(appName string, key string, value string) error

	GetAppResourceReport(appName string) (*ResourceReport, error)
	GetAllAppResourceReport() (ResourcesReport, error)

	SetAppDefaultResourceLimit(appName string, resource ResourceSpec, limit int) error
	ClearAppDefaultResourceLimit(appName string, resource ResourceSpec) error
	ClearAppDefaultResourceLimits(appName string) error

	SetAppProcessResourceLimit(appName string, process string, resource ResourceSpec, limit int) error
	ClearAppProcessResourceLimit(appName string, process string, resource ResourceSpec) error
	ClearAppProcessResourceLimits(appName string, process string) error

	SetAppResourceReservation(appName string, resource ResourceSpec, reserve int) error
	ClearAppResourceReservation(appName string, resource ResourceSpec) error
	ClearAppResourceReservations(appName string) error

	SetAppProcessResourceReservation(appName string, process string, resource ResourceSpec, reserve int) error
	ClearAppProcessResourceReservation(appName string, process string, resource ResourceSpec) error
	ClearAppProcessResourceReservations(appName string, process string) error

	// GetAllResourceReport() (ProcessesReport, error)

	GitSyncAppRepo(appName string, repo string, opt *GitSyncOptions) error

	SetEventLoggingEnabled(enabled bool) error
	GetEventLogs() (string, error)
	ListLoggedEvents() ([]string, error)

	TailAppLogs(appName string) (io.Reader, error)
	GetAppLogs(string) (string, error)
	GetNAppLogs(appName string, numLines int) (string, error)
	GetAppProcessLogs(appName, process string) (string, error)
	GetAppFailedDeployLogs(appName string) (string, error)
	GetAllFailedDeployLogs() (string, error)

	SetAppDeployChecksEnabled(appName string, enabled bool) error
	DeployAppFromDockerImage(appName, image string) (string, error)
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
