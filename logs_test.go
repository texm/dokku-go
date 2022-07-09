package dokku

import (
	"io/ioutil"
)

func (s *DokkuTestSuite) TestGetEventLogs() {
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

func (s *DokkuTestSuite) TestGetAppLogs() {
	r := s.Require()
	var err error

	testAppName := "test-logs-app"
	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	_, err = s.Client.GetAppLogs(testAppName)
	r.ErrorIs(err, AppNotDeployedError)
}

func (s *DokkuTestSuite) TestTailAppLogs() {
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
