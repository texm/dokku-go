package dokku

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type builderManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunBuilderManagerTestSuite(t *testing.T) {
	suite.Run(t, new(builderManagerTestSuite))
}

func (s *builderManagerTestSuite) TestReports() {
	r := s.Require()

	testAppName := "test-builder-app"
	r.NoError(s.Client.CreateApp(testAppName), "failed to create app")

	report, err := s.Client.GetAppBuilderDockerfileReport(testAppName)
	r.NoError(err)
	r.Equal("Dockerfile", report.GlobalDockerfilePath)

	report2, err2 := s.Client.GetAppBuilderPackReport(testAppName)
	r.NoError(err2)
	r.Equal("project.toml", report2.GlobalProjectTOMLPath)
}
