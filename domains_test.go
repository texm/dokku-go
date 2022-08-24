package dokku

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type domainsManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunDomainsManagerTestSuite(t *testing.T) {
	suite.Run(t, new(domainsManagerTestSuite))
}

func (s *domainsManagerTestSuite) TestGetAppDomains() {
	r := s.Require()

	testAppName := "test-domains-app"
	r.NoError(s.Client.CreateApp(testAppName))

	appDomain := "foo.example.com"
	globalDomain := "bar.example.com"

	r.NoError(s.Client.AddAppDomain(testAppName, appDomain))
	r.NoError(s.Client.AddGlobalDomain(globalDomain))

	report, err := s.Client.GetAppDomainsReport(testAppName)
	r.NoError(err)

	r.Len(report.AppDomains, 1)
	r.Equal(report.AppDomains[0], appDomain)

	r.Len(report.GlobalDomains, 1)
	r.Equal(report.GlobalDomains[0], globalDomain)
}
