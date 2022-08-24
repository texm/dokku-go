package dokku

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type gitManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunGitManagerTestSuite(t *testing.T) {
	suite.Run(t, new(gitManagerTestSuite))
}

func (s *gitManagerTestSuite) TestGitReport() {
	r := s.Require()
	var err error

	testAppName := "test-git-app"

	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	report, err := s.Client.GitGetAppReport(testAppName)
	r.NoError(err)
	r.Equal("master", report.DeployBranch)
}

func (s *gitManagerTestSuite) TestSyncGitRepo() {
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
