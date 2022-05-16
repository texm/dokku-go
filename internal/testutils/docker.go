package testutils

import (
	"context"
	"log"
	"os"
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

	// this is gross
	// mounting the host socket is easier than running some kind of VM setup, although it is somewhat insecure as
	// the test container now needs to run privileged
	// socketMount := testcontainers.BindMount(dockerSocketFile, defaultDockerSocketFile)

	req := testcontainers.ContainerRequest{
		Image:        testingImage,
		Privileged:   false,
		ExposedPorts: []string{"22/tcp"},
		// Mounts:       testcontainers.ContainerMounts{socketMount},
		WaitingFor: wait.ForListeningPort("22").WithStartupTimeout(startupTimeout),
	}
	gReq := testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	}

	log.Println("requesting dokku container")
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

	dockerSocketFile = path.Join(home, ".colima/docker.sock")
	os.Setenv("DOCKER_HOST", "unix://"+dockerSocketFile)
	os.Setenv("TESTCONTAINERS_DOCKER_SOCKET_OVERRIDE", dockerSocketFile)
	return nil
}
