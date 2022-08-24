package dokku

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type networkManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunNetworkManagerTestSuite(t *testing.T) {
	suite.Run(t, new(networkManagerTestSuite))
}
