package dokku

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
)

const (
	noAppsDokkuMessage         = "You haven't deployed any applications yet"
	nameTakenMessage           = "!     Name is already taken"
	lockCreatedMessage         = "-----> Deploy lock created"
	deployLockNotExistsMessage = "!     Deploy lock does not exist"
)

var (
	InvalidAppError        = errors.New("app does not exist")
	AppNotDeployedError    = errors.New("app is not deployed")
	NoDeployedAppsError    = errors.New("no apps have been deployed")
	NotImplementedError    = errors.New("method not implemented")
	UnexpectedMessageError = errors.New("unexpected confirmation message")
	NameTakenError         = errors.New("app name already in use")
)

type ExitCodeError struct {
	error
	Output string
	sshErr *ssh.ExitError
}

func newExitCodeErr(output string, err *ssh.ExitError) *ExitCodeError {
	return &ExitCodeError{
		Output: output,
		sshErr: err,
	}
}

func (xe *ExitCodeError) Error() string {
	return fmt.Sprintf("dokku error: %s", xe.sshErr.Error())
}

func (xe *ExitCodeError) Unwrap() error {
	return xe.sshErr
}

func (xe *ExitCodeError) CommandOutput() string {
	return xe.Output
}

type Client interface {
	dokkuSSHClient

	appManager
	builderManager
	certsManager
	checksManager
	configManager
	cronManager
	dockerManager
	domainsManager
	gitManager
	letsEncryptManager
	logsManager
	networkManager
	nginxManager
	pluginManager
	processManager
	proxyManager
	resourceManager
	schedulerManager
	sshKeysManager
	storageManager
}
