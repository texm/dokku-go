package dokku

func (s *DokkuTestSuite) TestBuilderReport() {
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
