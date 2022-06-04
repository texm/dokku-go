package dokku

import "errors"

const (
	noAppsDokkuMessage = "You haven't deployed any applications yet"
)

var (
	InvalidAppError     = errors.New("app does not exist")
	NoDeployedAppsError = errors.New("no apps have been deployed")
	NotImplementedError = errors.New("method not implemented")
)
