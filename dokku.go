package dokku

import "errors"

var (
	InvalidAppError     = errors.New("App does not exist")
	NotImplementedError = errors.New("Method not implemented")
)

func newDokkuError(msg string) error {
	return errors.New("dokku error: '" + msg + "'")
}
