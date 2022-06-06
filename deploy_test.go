package dokku

func (s *DokkuTestSuite) TestDockerImageDeploy() {
	r := s.Require()
	var err error

	testAppName := "test-logs-app"
	err = s.Client.CreateApp(testAppName)
	r.Nil(err, "failed to create app")

	image := "crccheck/hello-world"
	_, err = s.Client.DeployAppFromDockerImage(testAppName, image)
	r.Error(err, "docker socket isnt mounted, this should fail")
}
