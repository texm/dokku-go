package dokku

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type processManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunProcessManagerTestSuite(t *testing.T) {
	suite.Run(t, new(processManagerTestSuite))
}

func (s *processManagerTestSuite) TestGetProcessInfo() {
	r := s.Require()
	var err error

	testAppName := "test-process-app"

	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	err = s.Client.GetProcessInfo(testAppName)
	r.ErrorIs(err, AppNotDeployedError, "did not detect app not being deployed")
}

func (s *processManagerTestSuite) TestGetProcessReport() {
	r := s.Require()
	var err error

	testAppName := "test-process-app"

	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	appReport, err := s.Client.GetAppProcessReport(testAppName)
	r.NoError(err, "failed to get report")
	r.True(appReport.Restore)

	report, err := s.Client.GetAllProcessReport()
	r.NoError(err, "failed to get report")
	r.Contains(report, testAppName)
}

func (s *processManagerTestSuite) TestGetProcessScale() {
	r := s.Require()
	var err error

	testAppName := "test-process-app"

	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	scaleReport, err := s.Client.GetAppProcessScale(testAppName)
	r.NoError(err, "failed to get report")
	r.Contains(scaleReport, "web")
	r.Equal(scaleReport["web"], 1)
}

func (s *processManagerTestSuite) TestSetProcessScale() {
	r := s.Require()
	var err error

	testAppName := "test-process-app"

	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	_, err = s.Client.SetAppProcessScale(testAppName, "web", 2, true)
	r.NoError(err, "failed to set app scale")
}
