package dokku

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"io"
	"net"
	"os"
	"path"
	"strings"
	"time"
)

type SSHClient struct {
	BaseClient

	cfg    *SSHClientConfig
	sshCfg *ssh.ClientConfig
}

type SSHClientConfig struct {
	Host string
	// optional, defaults to 22
	Port                 string
	PrivateKey           *rsa.PrivateKey
	PrivateKeyBytes      []byte
	PrivateKeyPassphrase []byte

	// optional, defaults to 5 seconds
	ConnectionTimeout *time.Duration

	// optional, defaults to using $HOME/.ssh/known_hosts
	HostKeyCallback ssh.HostKeyCallback
}

type sshExecutor struct {
	conn *ssh.Client
}

var (
	InvalidPrivateKeyError = errors.New("invalid private key")
)

const (
	sshDokkuUser      = "dokku"
	defaultSSHTimeout = time.Second * 5
	knownHostsFile    = ".ssh/known_hosts"
)

func NewSSHClient(cfg *SSHClientConfig) (*SSHClient, error) {
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

	timeout := defaultSSHTimeout
	if cfg.ConnectionTimeout != nil {
		timeout = *cfg.ConnectionTimeout
	}

	sshConfig := &ssh.ClientConfig{
		User: sshDokkuUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		Timeout:         timeout,
		HostKeyCallback: hostKeyCallback,
	}

	addr := net.JoinHostPort(cfg.Host, cfg.Port)
	sshConn, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, err
	}

	client := &SSHClient{
		cfg:    cfg,
		sshCfg: sshConfig,
		BaseClient: BaseClient{
			executor: &sshExecutor{
				conn: sshConn,
			},
		},
	}

	return client, nil
}

func (c *SSHClient) Close() error {
	sshExec, ok := c.executor.(*sshExecutor)
	if ok {
		return sshExec.conn.Close()
	}
	return nil
}

func (e *sshExecutor) exec(cmd string, input io.Reader) (string, error) {
	session, err := e.conn.NewSession()
	if err != nil {
		return "", err
	}

	if input != nil {
		stdin, err := session.StdinPipe()
		if err != nil {
			return "", err
		}
		go io.Copy(stdin, input)
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
		var sshExitErr *ssh.ExitError
		if errors.As(cmdErr, &sshExitErr) {
			exitErr := &ExitCodeError{
				output:     cleaned,
				err:        err,
				exitStatus: sshExitErr.ExitStatus(),
			}
			return "", exitErr
		}
		return cleaned, err
	}

	return cleaned, nil
}

func (e *sshExecutor) execStreaming(cmd string, input io.Reader) (*CommandOutputStream, error) {
	session, err := e.conn.NewSession()
	if err != nil {
		return nil, err
	}

	if input != nil {
		stdin, err := session.StdinPipe()
		if err != nil {
			return nil, err
		}
		go io.Copy(stdin, input)
	}

	stream := &CommandOutputStream{}

	stream.Stdout, err = session.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stream.Stderr, err = session.StderrPipe()
	if err != nil {
		return nil, err
	}

	go func(stream *CommandOutputStream) {
		cmdErr := session.Run(cmd)
		if stream != nil {
			stream.Error = cmdErr
		}
		if sshErr := closeSession(session); sshErr != nil && stream != nil {
			if cmdErr != nil {
				stream.Error = fmt.Errorf("ssh close err '%s' after command error: %w", sshErr.Error(), cmdErr)
			} else {
				stream.Error = fmt.Errorf("ssh close err: %w", sshErr)
			}
		}
	}(stream)

	return stream, nil
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
