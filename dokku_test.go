package dokku

import (
	"context"
	"fmt"
	"github.com/texm/dokku-go/internal/testutils"

	"github.com/stretchr/testify/suite"
)

type dokkuTestSuite struct {
	suite.Suite
	Dokku                     *testutils.DokkuContainer
	AttachContainerTestLogger bool
	Client                    Client
}

func (s *dokkuTestSuite) SetupTest() {

}

func (s *dokkuTestSuite) SetupSuite() {
	ctx := context.Background()

	if err := s.CreateTestContainer(ctx); err != nil {
		s.T().Fatal("Failed to create dokku container: ", err)
	}

	if err := s.CreateTestClient(ctx, false); err != nil {
		s.T().Fatal("Failed to create default dokku client: ", err)
	}
}

func (s *dokkuTestSuite) TearDownSuite() {
	ctx := context.Background()

	apps, err := s.Client.ListApps()
	if err != nil {
		fmt.Println("failed to list apps")
	}
	for _, app := range apps {
		containers, err := s.Client.ListAppRunContainers(app)
		if err != nil {
			fmt.Println("failed to get containers for app")
		}
		fmt.Println("app", app, "containers", containers)
	}

	if s.Dokku != nil {
		s.Dokku.Cleanup(ctx)
	}

	if s.Client != nil {
		s.Client.Close()
	}
}

func (s *dokkuTestSuite) CreateTestContainer(ctx context.Context) error {
	dc, err := testutils.CreateDokkuContainer(ctx, s.AttachContainerTestLogger)
	if err != nil {
		return err
	}
	s.Dokku = dc

	if s.AttachContainerTestLogger {
		return dc.AttachTestLogger(ctx, s.T())
	}

	return nil
}

func (s *dokkuTestSuite) CreateTestClient(ctx context.Context, admin bool) error {
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
