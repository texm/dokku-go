package dokku

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type pluginManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunPluginManagerTestSuite(t *testing.T) {
	suite.Run(t, new(pluginManagerTestSuite))
}

func (s *pluginManagerTestSuite) TestListPlugins() {
	r := s.Require()

	plugins, err := s.Client.ListPlugins()
	r.NoError(err)
	r.NotEmpty(plugins)
	r.NotEmpty(plugins[0])
}
