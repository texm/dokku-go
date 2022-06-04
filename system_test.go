package dokku

import "fmt"

func (s *DokkuTestSuite) TestGetEventLogs() {
	r := s.Require()
	var err error

	testAppName := "test-process-app"

	err = s.Client.SetEventLoggingEnabled(false)
	r.Nil(err)

	err = s.Client.SetEventLoggingEnabled(true)
	r.Nil(err)

	err = s.Client.CreateApp(testAppName)
	r.Nil(err, "failed to create app")

	logs, err := s.Client.GetEventLogs()
	r.Nil(err)
	fmt.Println(logs)
}
