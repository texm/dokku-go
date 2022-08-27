package dokku

import (
	"context"
	"github.com/stretchr/testify/suite"
	"testing"
)

type gitManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunGitManagerTestSuite(t *testing.T) {
	suite.Run(t, &gitManagerTestSuite{
		dokkuTestSuite{
			DefaultAppName:            "test-git-app",
			AttachContainerTestLogger: true,
		},
	})
}

func (s *gitManagerTestSuite) TestGitReport() {
	r := s.Require()
	var err error

	report, err := s.Client.GitGetAppReport(s.DefaultAppName)
	r.NoError(err)
	r.Equal("master", report.DeployBranch)
}

func (s *gitManagerTestSuite) TestSyncGitRepo() {
	r := s.Require()
	var err error

	//ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
	r.NoError(s.Dokku.InstallBuildPacksCLI(context.Background()))

	r.NoError(s.Client.DisableAppDeployChecks(s.DefaultAppName))

	testRepo := "https://github.com/texm/go-hello-world-http.git"
	options := &GitSyncOptions{
		Build:  true,
		GitRef: "main",
	}
	err = s.Client.GitSyncAppRepo(s.DefaultAppName, testRepo, options)
	r.NoError(err)
}
