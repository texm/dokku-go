package dokku

import (
	"errors"
	"fmt"
	"strings"
)

const (
	appNotExistsMsg        = "does not exist"
	appNotDeployedMsg      = "has not been deployed"
	noAppsDokkuMsg         = "You haven't deployed any applications yet"
	nameTakenMsg           = "!     Name is already taken"
	lockCreatedMsg         = "-----> Deploy lock created"
	deployLockNotExistsMsg = "!     Deploy lock does not exist"
)

var (
	InvalidAppError        = errors.New("app does not exist")
	AppNotDeployedError    = errors.New("app is not deployed")
	NoDeployedAppsError    = errors.New("no apps have been deployed")
	NotImplementedError    = errors.New("method not implemented")
	UnexpectedMessageError = errors.New("unexpected confirmation message")
	NameTakenError         = errors.New("app name already in use")
)

func checkGenericErrors(output string) error {
	if strings.HasSuffix(output, appNotExistsMsg) {
		return InvalidAppError
	}
	if strings.HasSuffix(output, appNotDeployedMsg) {
		return AppNotDeployedError
	}
	if strings.Contains(output, noAppsDokkuMsg) {
		return NoDeployedAppsError
	}
	if strings.Contains(output, nameTakenMsg) {
		return NameTakenError
	}
	return nil
}

type ExitCodeError struct {
	error
	output     string
	exitStatus int
	err        error
}

func (xe *ExitCodeError) Error() string {
	return fmt.Sprintf("dokku error: '%s'", xe.Output())
}

func (xe *ExitCodeError) Unwrap() error {
	return xe.err
}

func (xe *ExitCodeError) Output() string {
	return xe.output
}

func (xe *ExitCodeError) ExitStatus() int {
	return xe.exitStatus
}
