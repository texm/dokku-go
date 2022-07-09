package dokku

func (s *DokkuTestSuite) TestSyncGitRepo() {
	r := s.Require()
	var err error

	testAppName := "test-git-app"

	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	testRepo := "https://github.com/simonvanderveldt/go-hello-world-http.git"
	err = s.Client.GitSyncAppRepo(testAppName, testRepo, nil)
	r.NoError(err)
}
