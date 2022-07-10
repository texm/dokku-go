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

	if err := s.createTestContainer(ctx); err != nil {
		s.T().Fatal("Failed to create dokku container: ", err)
	}

	if err := s.createTestClient(ctx); err != nil {
		s.T().Fatal("Failed to create default dokku client: ", err)
	}
}

func (s *DokkuTestSuite) TearDownTest() {
	ctx := context.Background()

	apps, err := s.Client.ListApps()
	if err != nil {
		fmt.Println("failed to list apps")
	}
	for _, app := range apps {
		if err := s.Client.DestroyApp(app); err != nil {
			fmt.Printf("failed to destroy app %s: %s\n", app, err.Error())
		}
	}

	if s.Dokku != nil {
		s.Dokku.Cleanup(ctx)
	}

	if s.Client != nil {
		s.Client.Close()
	}
}

func (s *DokkuTestSuite) createTestContainer(ctx context.Context) error {
	dc, err := testutils.CreateDokkuContainer(ctx)
	if err != nil {
		return err
	}
	dc.AttachLogConsumer(ctx)

	s.Dokku = dc
	return nil
}

func (s *DokkuTestSuite) createTestClient(ctx context.Context) error {
	keyPair, err := testutils.GenerateRSAKeyPair()
	if err != nil {
		return err
	}

	if err := s.Dokku.RegisterPublicKey(ctx, keyPair.PublicKey); err != nil {
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
