package dokku

func (s *DokkuTestSuite) TestGetAppConfig() {
	r := s.Require()

	testApp := "test-nginx-app"
	r.NoError(s.Client.CreateApp(testApp))

	_, err := s.Client.GetAppNginxConfig(testApp)
	r.ErrorIs(err, NginxNoConfigErr)
}
