package dokku

const (
	testAppName = "test-app"
)

func (s *DokkuTestSuite) TestCanCreateApp() {
	err := s.Client.CreateApp(testAppName)
	s.NoError(err, "Failed to create app")

	info, err := s.Client.GetAppInfo(testAppName)
	s.NoError(err, "Failed to get app info")
	s.Equal(testAppName, info.Name, "AppInfo.Name does not match")
}
