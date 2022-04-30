package dokku

import (
	"context"

	"github.com/texm/dokku-go/internal/testutils"
)

func (s *DokkuTestSuite) TestCanCreateClient() {
	r := s.Require()
	ctx := context.Background()

	keyPair, err := testutils.GenerateRSAKeyPair()
	r.Nil(err, "failed to create keypair")

	s.Dokku.RegisterPublicKey(ctx, keyPair.PublicKey)

	cfg := &ClientConfig{
		Host:            s.Dokku.Host,
		Port:            s.Dokku.SSHPort,
		PrivateKey:      keyPair.PrivateKey,
		HostKeyCallback: s.Dokku.HostKeyFunc(),
	}
	client, err := NewClient(cfg)
	r.Nil(err, "error while creating client")
	r.NotNil(client, "returned client is nil")
}
