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

func CreateDokkuContainer(ctx context.Context) (*DokkuContainer, error) {
	if runtime.GOOS == "darwin" {
		if err := setupColimaEnv(); err != nil {
			return nil, err
		}
	}

	// mounting the docker socket into the container is insecure, but nobody else should run this
	socketMount := testcontainers.BindMount(defaultDockerSocketFile, defaultDockerSocketFile)

	req := testcontainers.ContainerRequest{
		Image:        testingImage,
		Privileged:   false,
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

	localDockerSocketFile := path.Join(home, ".colima/docker.sock")
	localDockerSocketURI := fmt.Sprintf("unix://%s", localDockerSocketFile)

	if err := os.Setenv("DOCKER_HOST", "unix://"+localDockerSocketURI); err != nil {
		return err
	}
	if err := os.Setenv("TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE", localDockerSocketFile); err != nil {
		return err
	}

	return nil
}

func ensureMatchingDockerGroupId(ctx context.Context, container testcontainers.Container) error {
	var hostGid string
	if runtime.GOOS == "darwin" {
		// need to 'colima ssh' and groupmod ping to 998, groupmod docker 999
		hostGid = "999"
	} else {
		dockerGroup, err := user.LookupGroup("docker")
		if err != nil {
			return err
		}
		hostGid = dockerGroup.Gid
	}

	exitCode, err := container.Exec(ctx, []string{"groupmod", "-g", hostGid, "docker"})
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
