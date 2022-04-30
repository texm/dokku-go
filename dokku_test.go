package dokku

import (
	"context"
	"log"
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
	Client *Client
}

func (s *DokkuTestSuite) SetupSuite() {
	ctx := context.Background()

	if err := s.createTestContainer(ctx); err != nil {
		log.Fatal("Failed to create dokku container", err)
	}

	if err := s.createTestClient(ctx); err != nil {
		log.Fatal("Failed to create default dokku client", err)
	}
}

func (s *DokkuTestSuite) TearDownSuite() {
	ctx := context.Background()

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

	s.Dokku.RegisterPublicKey(ctx, keyPair.PublicKey)

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
