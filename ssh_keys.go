package dokku

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type sshKeysManager interface {
	AddSSHKey(name string, key []byte) error
	ListSSHKeys() ([]SSHKey, error)
	ListSSHKeysForName(name string) ([]SSHKey, error)
	RemoveSSHKeyByName(name string) error
	RemoveSSHKeyByFingerprint(fingerprint string) error
}

type SSHKey struct {
	Name              string `json:"name"`
	Fingerprint       string `json:"fingerprint"`
	AllowedSSHOptions string `json:"SSHCOMMAND_ALLOWED_KEYS"`
}

const (
	sshKeysAddCmd               = "ssh-keys:add %s"
	sshKeysListCmd              = "ssh-keys:list --format json %s"
	sshKeysRemoveFingerprintCmd = "ssh-keys:remove --fingerprint %s"
	sshKeysRemoveNameCmd        = "ssh-keys:remove %s"
)

// https://dokku.com/docs/deployment/user-management/#granting-other-unix-user-accounts-dokku-access

func (c *BaseClient) AddSSHKey(name string, key []byte) error {
	cmd := fmt.Sprintf(sshKeysAddCmd, name)
	reader := bytes.NewReader(key)
	_, err := c.ExecWithInput(cmd, reader)
	return err
}

func (c *BaseClient) ListSSHKeys() ([]SSHKey, error) {
	return c.ListSSHKeysForName("")
}

func (c *BaseClient) ListSSHKeysForName(name string) ([]SSHKey, error) {
	cmd := fmt.Sprintf(sshKeysListCmd, name)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var keys []SSHKey
	if err := json.Unmarshal([]byte(out), &keys); err != nil {
		return nil, err
	}
	return keys, nil
}

func (c *BaseClient) RemoveSSHKeyByName(name string) error {
	cmd := fmt.Sprintf(sshKeysRemoveNameCmd, name)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) RemoveSSHKeyByFingerprint(fingerprint string) error {
	cmd := fmt.Sprintf(sshKeysRemoveFingerprintCmd, fingerprint)
	_, err := c.Exec(cmd)
	return err
}
