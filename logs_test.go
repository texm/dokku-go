package dokku

import (
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"testing"
)

type logsManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunLogsManagerTestSuite(t *testing.T) {
	suite.Run(t, new(logsManagerTestSuite))
}

func (s *logsManagerTestSuite) TestGetEventLogs() {
	r := s.Require()
	var err error

	testAppName := "test-process-app"

	err = s.Client.SetEventLoggingEnabled(false)
	r.NoError(err)

	err = s.Client.SetEventLoggingEnabled(true)
	r.NoError(err)

	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	_, err = s.Client.GetEventLogs()
	r.NoError(err)
	// TODO: dokku logs doesn't seem to work here?
	// r.NotEmpty(logs)

	events, err := s.Client.ListLoggedEvents()
	r.NoError(err)
	r.NotEmpty(events)
}

func (s *logsManagerTestSuite) TestGetAppLogs() {
	r := s.Require()
	var err error

	testAppName := "test-logs-app"
	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	_, err = s.Client.GetAppLogs(testAppName)
	r.ErrorIs(err, AppNotDeployedError)
}

func (s *logsManagerTestSuite) TestTailAppLogs() {
	r := s.Require()
	var err error

	testAppName := "test-logs-app"
	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	reader, err := s.Client.TailAppLogs(testAppName)
	r.NoError(err)

	logs, err := ioutil.ReadAll(reader)
	r.NoError(err)
	r.Empty(logs)
}
