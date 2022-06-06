package testutils

import (
	"context"
	"fmt"
	"os"
	"os/user"
	"path"
	"runtime"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	testingImage            = "ghcr.io/texm/dokku-go:test-env"
	startupTimeout          = time.Second * 5
	defaultDockerSocketFile = "/var/run/docker.sock"
)

var (
	dockerSocketFile = defaultDockerSocketFile
)

func CreateDokkuContainer(ctx context.Context) (*DokkuContainer, error) {
	if runtime.GOOS == "darwin" {
		if err := setupColimaEnv(); err != nil {
			return nil, err
		}
	}

	// mounting the docker socket into the container is insecure, but nobody else should run this
	socketMount := testcontainers.BindMount(dockerSocketFile, defaultDockerSocketFile)

	req := testcontainers.ContainerRequest{
		Image:        testingImage,
		Privileged:   true,
		ExposedPorts: []string{"22/tcp"},
		Mounts:       testcontainers.ContainerMounts{socketMount},
		WaitingFor:   wait.ForListeningPort("22").WithStartupTimeout(startupTimeout),
	}
	gReq := testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	}

	container, err := testcontainers.GenericContainer(ctx, gReq)
	if err != nil {
		return nil, err
	}

	if err := ensureMatchingDockerGroupId(ctx, container); err != nil {
		return nil, maybeTerminateContainerAfterError(ctx, container, err)
	}

	host, err := container.Host(ctx)
	if err != nil {
		return nil, maybeTerminateContainerAfterError(ctx, container, err)
	}

	mappedPort, err := container.MappedPort(ctx, "22")
	if err != nil {
		return nil, maybeTerminateContainerAfterError(ctx, container, err)
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

	dockerSocketFile = path.Join(home, ".colima/docker.sock")
	if err := os.Setenv("DOCKER_HOST", "unix://"+dockerSocketFile); err != nil {
		return err
	}

	if err := os.Setenv("TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE", dockerSocketFile); err != nil {
		return err
	}

	return nil
}

func ensureMatchingDockerGroupId(ctx context.Context, container testcontainers.Container) error {
	dockerGroup, err := user.LookupGroup("docker")
	if err != nil {
		return err
	}

	exitCode, err := container.Exec(ctx, []string{"groupmod", "-g", dockerGroup.Gid, "docker"})
	if exitCode != 0 {
		return fmt.Errorf("failed to change gid of containerized docker group, got exit code %d\n", exitCode)
	}

	return err
}

func maybeTerminateContainerAfterError(ctx context.Context, container testcontainers.Container, err error) error {
	if termErr := container.Terminate(ctx); termErr != nil {
		return fmt.Errorf("failed to terminate container: %s after failing to handle error: %w", termErr.Error(), err)
	}
	return err
}
