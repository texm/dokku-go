package dokku

func (s *DokkuTestSuite) TestListAppCrons() {
	r := s.Require()

	testApp := "test-cron-app"
	r.NoError(s.Client.CreateApp(testApp))

	// exampleCronOutput := "{\"cron\":[{\"command\":\"echo hello\",\"schedule\":\"@daily\"}]}"

	_, err := s.Client.ListAppCronTasks(testApp)
	r.NoError(err)
}
