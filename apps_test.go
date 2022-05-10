package dokku

func (s *DokkuTestSuite) TestCanManageApp() {
	r := s.Require()
	var err error

	testAppName := "test-manage-app"

	err = s.Client.CreateApp(testAppName)
	r.Nil(err, "failed to create app")

	err = s.Client.DestroyApp(testAppName)
	r.Nil(err, "failed to destroy app")
}

func (s *DokkuTestSuite) TestCanGetAppInfo() {
	r := s.Require()
	var err error

	testAppName := "test-app-info"

	exists, err := s.Client.CheckAppExists(testAppName)
	r.Nil(err, "failed to check if app exists")
	r.False(exists, "incorrect result from exists check")

	err = s.Client.CreateApp(testAppName)
	r.Nil(err, "failed to create app")

	info, err := s.Client.GetAppInfo(testAppName)
	r.NoError(err, "Failed to get app info")
	r.Equal(testAppName, info.Name, "AppInfo.Name does not match")

	nilInfo, err := s.Client.GetAppInfo(testAppName + "-doesnt-exist")
	r.Error(err, "Failed to get app info")
	r.Nil(nilInfo, "returned app was not nil on error")
}
