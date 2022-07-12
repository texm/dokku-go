package dokku

import (
	"fmt"
)

/*
func (s *DokkuTestSuite) TestAddSSHKey() {
	r := s.Require()

	r.NoError(s.GrantAdminPriveleges())

	key, err := testutils.GenerateRSAKeyPair()
	r.NoError(err)

	r.NoError(s.Client.AddSSHKey("bleh", key.PublicKey))
}
*/

func (s *DokkuTestSuite) TestListSSHKeys() {
	r := s.Require()

	keys, err := s.Client.ListSSHKeys()
	fmt.Println(keys)
	r.NoError(err)
	r.NotEmpty(keys)
	r.Equal("test", keys[0].Name)
}
