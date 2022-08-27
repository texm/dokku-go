package dokku

import (
	"bytes"
	"io"
)

type BaseClient struct {
	stdout bytes.Buffer
	stderr bytes.Buffer

	executor commandExecutor
}

type commandExecutor interface {
	Exec(command string) (string, error)
	ExecStreaming(command string) (*CommandOutputStream, error)
	ExecWithStdin(command string, input io.Reader) (string, error)
}

type CommandOutputStream struct {
	Stdout io.Reader
	Stderr io.Reader
}

func (c *BaseClient) Exec(command string) (string, error) {
	return c.executor.Exec(command)
}

func (c *BaseClient) ExecStreaming(command string) (*CommandOutputStream, error) {
	return c.executor.ExecStreaming(command)
}

func (c *BaseClient) ExecWithStdin(command string, input io.Reader) (string, error) {
	return c.executor.ExecWithStdin(command, input)
}
