package dokku

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type cronManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunCronManagerTestSuite(t *testing.T) {
	suite.Run(t, new(configManagerTestSuite))
}

func (s *cronManagerTestSuite) TestListAppCrons() {
	r := s.Require()

	testApp := "test-cron-app"
	r.NoError(s.Client.CreateApp(testApp))

	// exampleCronOutput := "{\"cron\":[{\"command\":\"echo hello\",\"schedule\":\"@daily\"}]}"

	_, err := s.Client.ListAppCronTasks(testApp)
	r.NoError(err)
}
