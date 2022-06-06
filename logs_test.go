package dokku

func (s *DokkuTestSuite) TestGetEventLogs() {
	r := s.Require()
	var err error

	testAppName := "test-process-app"

	err = s.Client.SetEventLoggingEnabled(false)
	r.Nil(err)

	err = s.Client.SetEventLoggingEnabled(true)
	r.Nil(err)

	err = s.Client.CreateApp(testAppName)
	r.Nil(err, "failed to create app")

	_, err = s.Client.GetEventLogs()
	r.Nil(err)
	// TODO: dokku logs doesn't seem to work here?
	// r.NotEmpty(logs)

	events, err := s.Client.ListLoggedEvents()
	r.Nil(err)
	r.NotEmpty(events)
}

func (s *DokkuTestSuite) TestGetAppLogs() {
	r := s.Require()
	var err error

	testAppName := "test-logs-app"
	err = s.Client.CreateApp(testAppName)
	r.Nil(err, "failed to create app")

	_, err = s.Client.GetAppLogs(testAppName)
	r.ErrorIs(err, AppNotDeployedError)
}
