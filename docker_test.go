package dokku

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type dockerManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunDockerManagerTestSuite(t *testing.T) {
	suite.Run(t, new(dockerManagerTestSuite))
}
