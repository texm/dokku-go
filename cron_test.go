package dokku

import "fmt"

func (s *DokkuTestSuite) TestListAppCrons() {
	r := s.Require()

	testApp := "test-cron-app"
	r.NoError(s.Client.CreateApp(testApp))

	// appJsonCron := "{\"cron\":[{\"command\":\"echo hello\",\"schedule\":\"@daily\"}]}"

	crons, err := s.Client.ListAppCronTasks(testApp)
	r.NoError(err)
	fmt.Println(crons)
}
