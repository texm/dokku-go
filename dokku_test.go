package dokku

import (
	"context"
	"fmt"
	"testing"

	"github.com/texm/dokku-go/internal/testutils"

	"github.com/stretchr/testify/suite"
)

func TestRunDokkuTestSuite(t *testing.T) {
	suite.Run(t, new(DokkuTestSuite))
}

type DokkuTestSuite struct {
	suite.Suite
	Dokku  *testutils.DokkuContainer
	Client Client
}

func (s *DokkuTestSuite) SetupTest() {
	ctx := context.Background()

	if err := s.CreateTestContainer(ctx); err != nil {
		s.T().Fatal("Failed to create dokku container: ", err)
	}

	if err := s.CreateTestClient(ctx, false); err != nil {
		s.T().Fatal("Failed to create default dokku client: ", err)
	}
}

func (s *DokkuTestSuite) TearDownTest() {
	ctx := context.Background()

	/*apps, err := s.Client.ListApps()
	if err != nil {
		fmt.Println("failed to list apps")
	}
	for _, app := range apps {
		if err := s.Client.DestroyApp(app); err != nil {
			fmt.Printf("failed to destroy app %s: %s\n", app, err.Error())
		}
	}*/

	if s.Dokku != nil {
		s.Dokku.Cleanup(ctx)
	}

	if s.Client != nil {
		s.Client.Close()
	}
}

func (s *DokkuTestSuite) CreateTestContainer(ctx context.Context) error {
	dc, err := testutils.CreateDokkuContainer(ctx)
	if err != nil {
		return err
	}
	if err := dc.AttachLogConsumer(ctx); err != nil {
		return err
	}

	s.Dokku = dc
	return nil
}

func (s *DokkuTestSuite) CreateTestClient(ctx context.Context, admin bool) error {
	keyPair, err := testutils.GenerateRSAKeyPair()
	if err != nil {
		return err
	}

	keyName := "test"
	if admin {
		keyName = "admin"
	}

	if err := s.Dokku.RegisterPublicKey(ctx, keyPair.PublicKey, keyName); err != nil {
		return err
	}

	cfg := &ClientConfig{
		Host:            s.Dokku.Host,
		Port:            s.Dokku.SSHPort,
		PrivateKey:      keyPair.PrivateKey,
		HostKeyCallback: s.Dokku.HostKeyFunc(),
	}
	client, err := NewClient(cfg)
	if err != nil {
		return err
	}

	if err := client.Dial(); err != nil {
		return err
	}

	s.Client = client
	return nil
}

func (s *DokkuTestSuite) GrantAdminPriveleges() error {
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
