package dokku

func (s *DokkuTestSuite) TestGetProcessInfo() {
	r := s.Require()
	var err error

	testAppName := "test-process-app"

	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	err = s.Client.GetProcessInfo(testAppName)
	r.ErrorIs(err, AppNotDeployedError, "did not detect app not being deployed")
}

func (s *DokkuTestSuite) TestGetProcessReport() {
	r := s.Require()
	var err error

	testAppName := "test-process-app"

	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	appReport, err := s.Client.GetAppProcessReport(testAppName)
	r.NoError(err, "failed to get report")
	r.True(appReport.Restore)

	report, err := s.Client.GetAllProcessReport()
	r.NoError(err, "failed to get report")
	r.Contains(report, testAppName)
}
