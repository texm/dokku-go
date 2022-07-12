package testutils

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/testcontainers/testcontainers-go"
	"golang.org/x/crypto/ssh"
)

const (
	testKeyPath     = "/home/dokku/test_key.pub"
	testKeyFileMode = 0666
)

type DokkuContainer struct {
	testcontainers.Container
	Host        string
	SSHPort     string
	logConsumer *testLogConsumer
}

type testLogConsumer struct {
	Msgs []string
}

func (g *testLogConsumer) Accept(l testcontainers.Log) {
	g.Msgs = append(g.Msgs, string(l.Content))
}

func (dc *DokkuContainer) Cleanup(ctx context.Context) {
	dc.Terminate(ctx)
	if dc.logConsumer != nil {
		dc.StopLogProducer()
	}
}

func (dc *DokkuContainer) GetLogs() []string {
	if dc.logConsumer != nil {
		return dc.logConsumer.Msgs
	}
	return []string{}
}

func (dc *DokkuContainer) AttachLogConsumer(ctx context.Context) error {
	dc.logConsumer = &testLogConsumer{
		Msgs: []string{},
	}

	if err := dc.StartLogProducer(ctx); err != nil {
		return err
	}

	dc.FollowOutput(dc.logConsumer)

	return nil
}

func (dc *DokkuContainer) HostKeyFunc() func(string, net.Addr, ssh.PublicKey) error {
	return func(host string, remote net.Addr, key ssh.PublicKey) error {
		if net.JoinHostPort(dc.Host, dc.SSHPort) != host {
			return errors.New("invalid host supplied for handshake?")
		}
		return nil
	}
}

func (dc *DokkuContainer) RegisterPublicKey(ctx context.Context, key []byte, name string) error {
	err := dc.CopyToContainer(ctx, key, testKeyPath, testKeyFileMode)
	if err != nil {
		return err
	}

	chownCmd := []string{"/usr/bin/dokku", "ssh-keys:add", name, testKeyPath}
	retCode, err := dc.Exec(ctx, chownCmd)
	if err != nil {
		return fmt.Errorf("failed to add ssh key: %w", err)
	} else if retCode != 0 {
		return fmt.Errorf("failed to add ssh key: got exit code %d", retCode)
	}

	return nil
}
