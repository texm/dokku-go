package dokku

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type logsManager interface {
	SetEventLoggingEnabled(enabled bool) error
	GetEventLogs() (string, error)
	ListLoggedEvents() ([]string, error)

	TailAppLogs(appName string) (io.Reader, error)
	GetAppLogs(string) (string, error)
	GetNAppLogs(appName string, numLines int) (string, error)
	GetAppProcessLogs(appName, process string) (string, error)
	GetAppFailedDeployLogs(appName string) (string, error)
	GetAllFailedDeployLogs() (string, error)
}

const (
	disabledEventLoggerMsg = "Disabling dokku events logger"
	enabledEventLoggerMsg  = "Enabling dokku events logger"
)

const (
	appLogsCmd             = "logs %s --quiet"
	appNLogsCmd            = "logs %s -n %d --quiet"
	appTailLogsCmd         = "logs %s --tail --quiet"
	appLogsProcessCmd      = "logs %s --quiet --ps %s"
	appFailedDeployLogsCmd = "logs:failed %s"
	allFailedDeployLogsCmd = "logs:failed --all"

	eventsCmd     = "events"
	eventsListCmd = "events:list --quiet"
	eventsOnCmd   = "events:on"
	eventsOffCmd  = "events:off"
)

func (c *BaseClient) TailAppLogs(appName string) (io.Reader, error) {
	cmd := fmt.Sprintf(appTailLogsCmd, appName)
	stream, err := c.ExecStreaming(cmd)
	if err != nil {
		return nil, err
	}

	pr, pw := io.Pipe()
	go func() {
		errBuf := bufio.NewReader(stream.Stderr)
		outBuf := bufio.NewReader(stream.Stdout)
		for {
			line, _, err := outBuf.ReadLine()
			if err != nil {
				_ = pw.CloseWithError(err)
			}

			if errBuf.Buffered() > 0 {
				stderr, _, err := errBuf.ReadLine()
				if err != nil {
					log.Printf("error while reading stderr: %s", err.Error())
				}
				_ = pw.CloseWithError(fmt.Errorf("stderr: %s", stderr))
			}

			_, err = pw.Write(line)
			if err != nil {
				_ = pw.CloseWithError(err)
			}
		}
	}()

	return pr, nil
}

func (c *BaseClient) GetNAppLogs(appName string, numLines int) (string, error) {
	cmd := fmt.Sprintf(appNLogsCmd, appName, numLines)
	return c.Exec(cmd)
}

func (c *BaseClient) GetAppLogs(appName string) (string, error) {
	return c.GetNAppLogs(appName, 50)
}

func (c *BaseClient) GetAppProcessLogs(appName, process string) (string, error) {
	cmd := fmt.Sprintf(appLogsProcessCmd, appName, process)
	return c.Exec(cmd)
}

func (c *BaseClient) GetAppFailedDeployLogs(appName string) (string, error) {
	cmd := fmt.Sprintf(appFailedDeployLogsCmd, appName)
	return c.Exec(cmd)
}

func (c *BaseClient) GetAllFailedDeployLogs() (string, error) {
	return c.Exec(allFailedDeployLogsCmd)
}

func (c *BaseClient) SetEventLoggingEnabled(enabled bool) error {
	var err error
	var output string
	if !enabled {
		output, err = c.Exec(eventsOffCmd)
		if output != disabledEventLoggerMsg {
			return UnexpectedMessageError
		}
	} else {
		output, err = c.Exec(eventsOnCmd)
		if output != enabledEventLoggerMsg {
			return UnexpectedMessageError
		}
	}
	return err
}

func (c *BaseClient) GetEventLogs() (string, error) {
	return c.Exec(eventsCmd)
}

func (c *BaseClient) ListLoggedEvents() ([]string, error) {
	var events []string
	sEvents, err := c.Exec(eventsListCmd)
	if err != nil {
		return events, err
	}
	events = strings.Split(sEvents, "\n")
	return events, nil
}
