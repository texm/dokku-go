package dokku

import (
	"bytes"
	"io"
)

type BaseClient struct {
	executor commandExecutor
}

type commandExecutor interface {
	exec(command string, input *bytes.Reader) (string, error)
	execStreaming(command string, input *bytes.Reader) (*CommandOutputStream, error)
}

type CommandOutputStream struct {
	Stdout io.Reader
	Stderr io.Reader
	Error  error
}

func (c *BaseClient) Exec(command string) (string, error) {
	return c.executor.exec(command, nil)
}

func (c *BaseClient) ExecStreaming(command string) (*CommandOutputStream, error) {
	return c.executor.execStreaming(command, nil)
}

func (c *BaseClient) ExecWithInput(command string, input *bytes.Reader) (string, error) {
	return c.executor.exec(command, input)
}

func (c *BaseClient) ExecWithInputStreaming(command string, input *bytes.Reader) (*CommandOutputStream, error) {
	return c.executor.execStreaming(command, input)
}
