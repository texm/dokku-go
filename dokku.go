package dokku

import (
	"errors"
	"time"
)

const (
	noAppsDokkuMessage          = "You haven't deployed any applications yet"
	nameTakenMessage            = "!     Name is already taken"
	lockCreatedMessage          = "-----> Deploy lock created"
	deployLockExistsMessage     = "Deploy lock exists"
	deployLockNotExistsMessage  = "!     Deploy lock does not exist"
	appNotExistsMessageTemplate = "!     App %s does not exist"
)

var (
	InvalidAppError        = errors.New("app does not exist")
	AppNotDeployedError    = errors.New("app is not deployed")
	NoDeployedAppsError    = errors.New("no apps have been deployed")
	NotImplementedError    = errors.New("method not implemented")
	UnexpectedMessageError = errors.New("unexpected confirmation message")
	NameTakenError         = errors.New("app name already in use")
)

type Client interface {
	Dial() error
	DialWithTimeout(timeout time.Duration) error
	Close() error

	Exec(command string) (string, error)

	appManager
	processManager
	resourceManager
	gitManager
	logsManager
	checksManager
	networkManager

	//builderManager
	//configManager
	//cronManager
	//dockerManager
	//domainsManager
	//letsEncryptManager
	//nginxManager
	//pluginManager
	//proxyManager
	//schedulerManager
	//sshKeysManager
	//storageManager

	//GetDokkuVersion() (string, error)
}
