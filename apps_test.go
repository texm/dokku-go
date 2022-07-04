package dokku

func (s *DokkuTestSuite) TestCanManageApp() {
	r := s.Require()
	var err error

	testAppName := "test-manage-app"

	err = s.Client.CreateApp(testAppName)
	r.Nil(err, "failed to create app")

	err = s.Client.DestroyApp(testAppName)
	r.Nil(err, "failed to destroy app")

	exists, err := s.Client.CheckAppExists(testAppName)
	r.Nil(err, "failed to check if app exists")
	r.False(exists, "app was not correctly destroyed")
}

func (s *DokkuTestSuite) TestCanCreateApp() {
	r := s.Require()

	testAppName := "test-manage-app"
	err := s.Client.CreateApp(testAppName)
	r.Nil(err, "failed to create app")
}

func (s *DokkuTestSuite) TestDuplicateAppName() {
	r := s.Require()

	testAppName := "test-duplicate-app"
	err := s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	err = s.Client.CreateApp(testAppName)
	r.ErrorIs(err, NameTakenError)
}

func (s *DokkuTestSuite) TestNoAppsFunctionality() {
	r := s.Require()
	var err error

	_, err = s.Client.GetAllAppReport()
	r.Error(err, "didnt error with no apps?")
	r.ErrorIs(err, NoDeployedAppsError)
}

func (s *DokkuTestSuite) TestCanGetAppInfo() {
	r := s.Require()
	var err error

	testAppName := "test-app-info"
	testAppName2 := "test-app-info2"

	exists, err := s.Client.CheckAppExists(testAppName)
	r.Nil(err, "failed to check if app exists")
	r.False(exists, "incorrect result from exists check")

	err = s.Client.CreateApp(testAppName)
	r.Nil(err, "failed to create app 1")

	err = s.Client.CreateApp(testAppName2)
	r.Nil(err, "failed to create app 2")

	_, err = s.Client.GetAppReport(testAppName)
	r.NoError(err, "Failed to get app info")

	nilInfo, err := s.Client.GetAppReport(testAppName + "-doesnt-exist")
	r.Error(err, "Failed to get app info")
	r.Nil(nilInfo, "returned app was not nil on error")

	report, err := s.Client.GetAllAppReport()
	r.NoError(err, "Failed to get app info")
	r.Contains(report, testAppName)
	r.Contains(report, testAppName2)
}
