package dokku

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"testing"
)

type sshKeysManagerTestSuite struct {
	dokkuTestSuite
}

func TestRunSSHKeysManagerTestSuite(t *testing.T) {
	suite.Run(t, new(sshKeysManagerTestSuite))
}

/*
func (s *sshKeysManagerTestSuite) GrantAdminPrivileges() error {
	if err := s.Client.Close(); err != nil {
		return err
	}

	ctx := context.Background()

	chownCmd := []string{"/usr/bin/dokku", "ssh-keys:remove", "test"}
	retCode, err := s.Dokku.Exec(ctx, chownCmd)
	if err != nil {
		return fmt.Errorf("failed to remove ssh key: %w", err)
	} else if retCode != 0 {
		return fmt.Errorf("failed to remove ssh key: got exit code %d", retCode)
	}

	return s.CreateTestClient(ctx, true)
}

func (s *sshKeysManagerTestSuite) TestAddSSHKey() {
	r := s.Require()

	r.NoError(s.GrantAdminPrivileges())

	key, err := testutils.GenerateRSAKeyPair()
	r.NoError(err)

	r.NoError(s.Client.AddSSHKey("bleh", key.PublicKey))
}
*/

func (s *sshKeysManagerTestSuite) TestListSSHKeys() {
	r := s.Require()

	keys, err := s.Client.ListSSHKeys()
	fmt.Println(keys)
	r.NoError(err)
	r.NotEmpty(keys)
	r.Equal("test", keys[0].Name)
}
