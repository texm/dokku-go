package dokku

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type appManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunAppManagerTestSuite(t *testing.T) {
	suite.Run(t, new(appManagerTestSuite))
}

func (s *appManagerTestSuite) TestManagementOptionsFlags() {
	r := s.Require()

	opts := AppManagementOptions{}
	r.Equal("", opts.asFlags())
	opts.SkipDeploy = true
	r.Equal("--skip-deploy", opts.asFlags())
	opts.IgnoreExisting = true
	r.Equal("--skip-deploy --ignore-existing", opts.asFlags())
}

func (s *appManagerTestSuite) TestCreate() {
	s.Require().NoError(
		s.Client.CreateApp("test-create-app"))
}

func (s *appManagerTestSuite) TestDestroy() {
	r := s.Require()

	testAppName := "test-manage-app"

	r.NoError(
		s.Client.CreateApp(testAppName), "failed to create app")

	r.NoError(
		s.Client.DestroyApp(testAppName), "failed to destroy app")

	exists, err := s.Client.CheckAppExists(testAppName)
	r.False(exists, "app was not correctly destroyed")
	r.NoError(err, "failed to check if app exists")
}

func (s *appManagerTestSuite) TestDuplicateName() {
	r := s.Require()

	testAppName := "test-duplicate-app"
	err := s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	err = s.Client.CreateApp(testAppName)
	r.ErrorIs(err, NameTakenError)
}

func (s *appManagerTestSuite) TestNoAppsError() {
	r := s.Require()
	var err error

	_, err = s.Client.GetAllAppReport()
	r.Error(err, "didnt error with no apps?")
	r.ErrorIs(err, NoDeployedAppsError)
}

func (s *appManagerTestSuite) TestGetAppReport() {
	r := s.Require()
	var err error

	testAppName := "test-app-info"
	testAppName2 := "test-app-info2"

	exists, err := s.Client.CheckAppExists(testAppName)
	r.NoError(err, "failed to check if app exists")
	r.False(exists, "incorrect result from exists check")

	r.NoError(s.Client.CreateApp(testAppName))

	r.NoError(s.Client.CreateApp(testAppName2))

	appReport, err := s.Client.GetAppReport(testAppName)
	r.NoError(err, "Failed to get app info")
	r.NotNil(appReport)

	nilReport, err := s.Client.GetAppReport(testAppName + "-doesnt-exist")
	r.Error(err, "Failed to get app info")
	r.Nil(nilReport, "returned app was not nil on error")

	report, err := s.Client.GetAllAppReport()
	r.NoError(err, "Failed to get app info")
	r.Contains(report, testAppName)
	r.Contains(report, testAppName2)
}
