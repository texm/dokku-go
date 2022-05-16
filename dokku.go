package dokku

import "errors"

const (
	noAppsDokkuMessage = "!     You haven't deployed any applications yet"
)

var (
	InvalidAppError     = errors.New("App does not exist")
	NoDeployedAppsError = errors.New("No apps have been deployed")
	NotImplementedError = errors.New("Method not implemented")
)

func newDokkuError(msg string) error {
	return errors.New("dokku error: '" + msg + "'")
}
