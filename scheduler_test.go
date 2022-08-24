package dokku

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type schedulerManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunSchedulerManagerTestSuite(t *testing.T) {
	suite.Run(t, new(schedulerManagerTestSuite))
}
