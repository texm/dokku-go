package dokku

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type proxyManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunProxyManagerTestSuite(t *testing.T) {
	suite.Run(t, new(proxyManagerTestSuite))
}
