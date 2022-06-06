package dokku

import (
	"fmt"
	"strings"
)

const (
	disabledEventLoggerMsg = "Disabling dokku events logger"
	enabledEventLoggerMsg  = "Enabling dokku events logger"
)

const (
	appLogsCmd             = "logs %s --quiet"
	appLogsProcessCmd      = "logs %s --quiet --ps %s"
	appFailedDeployLogsCmd = "logs:failed %s"
	allFailedDeployLogsCmd = "logs:failed --all"

	eventsCmd     = "events"
	eventsListCmd = "events:list --quiet"
	eventsOnCmd   = "events:on"
	eventsOffCmd  = "events:off"
)

func (c *DefaultClient) GetAppLogs(appName string) (string, error) {
	cmd := fmt.Sprintf(appLogsCmd, appName)
	return c.exec(cmd)
}

func (c *DefaultClient) GetAppProcessLogs(appName, process string) (string, error) {
	cmd := fmt.Sprintf(appLogsProcessCmd, appName, process)
	return c.exec(cmd)
}

func (c *DefaultClient) GetAppFailedDeployLogs(appName string) (string, error) {
	cmd := fmt.Sprintf(appFailedDeployLogsCmd, appName)
	return c.exec(cmd)
}

func (c *DefaultClient) GetAllFailedDeployLogs() (string, error) {
	return c.exec(allFailedDeployLogsCmd)
}

func (c *DefaultClient) SetEventLoggingEnabled(enabled bool) error {
	var err error
	var output string
	if !enabled {
		output, err = c.exec(eventsOffCmd)
		if output != disabledEventLoggerMsg {
			return UnexpectedMessageError
		}
	} else {
		output, err = c.exec(eventsOnCmd)
		if output != enabledEventLoggerMsg {
			return UnexpectedMessageError
		}
	}
	return err
}

func (c *DefaultClient) GetEventLogs() (string, error) {
	return c.exec(eventsCmd)
}

func (c *DefaultClient) ListLoggedEvents() ([]string, error) {
	var events []string
	sEvents, err := c.exec(eventsListCmd)
	if err != nil {
		return events, err
	}
	events = strings.Split(sEvents, "\n")
	return events, nil
}
