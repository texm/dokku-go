package dokku

func (s *DokkuTestSuite) TestGetProcessInfo() {
	r := s.Require()
	var err error

	testAppName := "test-process-app"

	err = s.Client.CreateApp(testAppName)
	r.Nil(err, "failed to create app")

	err = s.Client.GetProcessInfo(testAppName)
	r.Nil(err, "failed to get info")
}

func (s *DokkuTestSuite) TestGetProcessReport() {
	r := s.Require()
	var err error

	testAppName := "test-process-app"

	err = s.Client.CreateApp(testAppName)
	r.Nil(err, "failed to create app")

	report, err := s.Client.GetAllProcessReport()
	r.Nil(err, "failed to get report")
	r.Contains(report, testAppName)
}
