package dokku

import "fmt"

func (s *DokkuTestSuite) TestGetResourceReport() {
	r := s.Require()
	var err error

	testAppName := "test-resource-app"

	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	report, err := s.Client.GetAppResourceReport(testAppName)
	r.NoError(err)
	fmt.Println(report)
}
