package dokku

func (s *DokkuTestSuite) TestGitReport() {
	r := s.Require()
	var err error

	testAppName := "test-git-app"

	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	report, err := s.Client.GitGetAppReport(testAppName)
	r.NoError(err)
	r.Equal("master", report.DeployBranch)
}

func (s *DokkuTestSuite) TestSyncGitRepo() {
	r := s.Require()
	var err error

	testAppName := "test-git-app"

	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	r.NoError(s.Client.DisableAppDeployChecks(testAppName))

	testRepo := "https://github.com/texm/go-hello-world-http.git"
	options := &GitSyncOptions{
		Build:  true,
		GitRef: "main",
	}
	err = s.Client.GitSyncAppRepo(testAppName, testRepo, options)
	r.NoError(err)
}
