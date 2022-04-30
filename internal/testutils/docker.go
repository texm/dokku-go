package testutils

import (
	"context"
	"os"
	"path"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	colimaSocket   = ".colima/docker.sock"
	containerHash  = "sha256:d5eb1a4a2d713ba28a8afeed6126cebfba6086881101536076dd30f10d28f30e"
	testKeyPath    = "/home/dokku/test_key.pub"
	startupTimeout = time.Second * 8
)

func CreateDokkuContainer(ctx context.Context) (*DokkuContainer, error) {
	// check platform
	if err := setupColimaEnv(); err != nil {
		return nil, err
	}

	waitStrategy := wait.ForListeningPort("22").WithStartupTimeout(startupTimeout)
	req := testcontainers.ContainerRequest{
		Image:        containerHash,
		Privileged:   false,
		SkipReaper:   true,
		ExposedPorts: []string{"22/tcp"},
		WaitingFor:   waitStrategy,
	}
	gReq := testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	}
	container, err := testcontainers.GenericContainer(ctx, gReq)
	if err != nil {
		return nil, err
	}

	host, err := container.Host(ctx)
	if err != nil {
		container.Terminate(ctx)
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, "22")
	if err != nil {
		container.Terminate(ctx)
		return nil, err
	}

	dc := &DokkuContainer{
		Container: container,
		Host:      host,
		SSHPort:   mappedPort.Port(),
	}

	return dc, nil
}

func setupColimaEnv() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	socketFile := path.Join(home, colimaSocket)
	os.Setenv("DOCKER_HOST", "unix://"+socketFile)
	os.Setenv("TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE", socketFile)
	return nil
}
