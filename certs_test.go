package dokku

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type certsManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunCertsManagerTestSuite(t *testing.T) {
	suite.Run(t, new(certsManagerTestSuite))
}
