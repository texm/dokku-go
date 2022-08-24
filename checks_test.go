package dokku

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type checksManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunChecksManagerTestSuite(t *testing.T) {
	suite.Run(t, new(checksManagerTestSuite))
}

func (s *checksManagerTestSuite) TestGetChecksReport() {
	r := s.Require()

	testAppName := "test-deploy-app"
	r.NoError(s.Client.CreateApp(testAppName), "failed to create app")

	r.NoError(s.Client.DisableAppDeployChecks(testAppName))

	report, err := s.Client.GetAppDeployChecksReport(testAppName)
	r.NoError(err)
	r.True(report.AllDisabled)

	fullReport, err := s.Client.GetDeployChecksReport()
	r.NoError(err)
	r.Contains(fullReport, testAppName)
}
