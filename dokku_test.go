package dokku

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/suite"
	"github.com/texm/dokku-go/internal/testutils"
	"os/exec"
	"strings"
)

type dokkuTestSuite struct {
	suite.Suite
	Dokku                     *testutils.DokkuContainer
	AttachContainerTestLogger bool
	DefaultAppName            string
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

	if s.DefaultAppName != "" {
		if err := s.Client.CreateApp(s.DefaultAppName); err != nil {
			s.T().Fatal("failed to create default app")
		}
	}
}

func (s *dokkuTestSuite) TearDownSuite() {
	ctx := context.Background()

	if err := s.cleanupAppDockerContainers(); err != nil {
		s.T().Errorf("failed to cleanup app containers: %s", err.Error())
	}

	if s.Dokku != nil {
		s.Dokku.Cleanup(ctx)
	}

	if s.Client != nil {
		if err := s.Client.Close(); err != nil {
			// ignore err https://github.com/golang/go/issues/32453
			// s.T().Errorf("failed to close client %s", err.Error())
		}
	}
}

func (s *dokkuTestSuite) cleanupAppDockerContainers() error {
	apps, err := s.Client.ListApps()
	if err != nil {
		return fmt.Errorf("failed apps list: %w", err)
	}

	for _, appName := range apps {
		filter := fmt.Sprintf("label=com.dokku.app-name=%s", appName)
		out, err := exec.Command("docker", "ps", "-a", "-f", filter).Output()
		if err != nil {
			return fmt.Errorf("failed docker ps for app '%s': %w", appName, err)
		}

		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		for _, line := range lines[1:] {
			containerID := strings.Split(line, " ")[0]
			cmd := exec.Command("docker", "rm", "-f", containerID)
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("failed docker rm (id=%s) for app '%s': %w", containerID, appName, err)
			}
		}
	}

	return nil
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
