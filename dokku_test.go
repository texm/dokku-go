package dokku

import (
	"context"
	"log"
	"testing"

	"github.com/texm/dokku-go/internal/testutils"

	"github.com/stretchr/testify/suite"
)

func TestDokkuTestSuite(t *testing.T) {
	suite.Run(t, new(DokkuTestSuite))
}

type DokkuTestSuite struct {
	suite.Suite
	Dokku  *testutils.DokkuContainer
	Client *Client
}

func (s *DokkuTestSuite) SetupSuite() {
	ctx := context.Background()

	dc, err := testutils.CreateDokkuContainer(ctx)
	if err != nil {
		log.Fatal("Failed to create dokku container", err)
	}

	keyPair, err := testutils.GenerateRSAKeyPair()
	if err != nil {
		log.Fatal("Failed to create keypair", err)
	}
	dc.AttachLogConsumer(ctx)
	dc.RegisterDokkuPublicKey(ctx, keyPair.PublicKey)

	cfg := &ClientConfig{
		Host:            dc.Host,
		Port:            dc.SSHPort,
		PrivateKey:      keyPair.PrivateKey,
		HostKeyCallback: dc.HostKeyFunc(),
	}
	client, err := NewClient(cfg)
	if err != nil {
		log.Fatal("Failed to create dokku client", err)
	}
	if err := client.Dial(); err != nil {
		log.Fatal("Failed to create dokku client", err)
	}

	s.Dokku = dc
	s.Client = client
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
