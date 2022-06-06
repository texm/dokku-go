package dokku

import (
	"strings"
)

const (
	disabledEventLoggerMsg = "Disabling dokku events logger"
	enabledEventLoggerMsg  = "Enabling dokku events logger"
)

func (c *DefaultClient) SetEventLoggingEnabled(enabled bool) error {
	var err error
	var output string
	if !enabled {
		output, err = c.exec("events:off")
		if output != disabledEventLoggerMsg {
			return UnexpectedMessageError
		}
	} else {
		output, err = c.exec("events:on")
		if output != enabledEventLoggerMsg {
			return UnexpectedMessageError
		}
	}
	return err
}

func (c *DefaultClient) GetEventLogs() (string, error) {
	return c.exec("events")
}

func (c *DefaultClient) ListLoggedEvents() ([]string, error) {
	var events []string
	sEvents, err := c.exec("events:list --quiet")
	if err != nil {
		return events, err
	}
	events = strings.Split(sEvents, "\n")
	return events, nil
}
