package dokku

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type dokkuClientTestSuite struct {
	suite.Suite
}

func TestRunDokkuClientTestSuite(t *testing.T) {
	suite.Run(t, new(dokkuClientTestSuite))
}

func (s *checksManagerTestSuite) TestCreateSSHClient() {

}
