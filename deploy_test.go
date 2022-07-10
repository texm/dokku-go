package dokku

func (s *DokkuTestSuite) TestDockerImageDeploy() {
	r := s.Require()
	var err error

	testAppName := "test-deploy-app"
	r.NoError(s.Client.CreateApp(testAppName), "failed to create app")

	r.NoError(s.Client.SetAppDeployChecksEnabled(testAppName, false))

	_, err = s.Client.DeployAppFromDockerImage(testAppName, "dokku/node-js-getting-started:latest")
	r.NoError(err)
}
