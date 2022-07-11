package dokku

func (s *DokkuTestSuite) TestListPlugins() {
	r := s.Require()

	plugins, err := s.Client.ListPlugins()
	r.NoError(err)
	r.NotEmpty(plugins)
	r.NotEmpty(plugins[0])
}
