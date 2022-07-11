package dokku

type sshKeysManager interface {
	AddSSHKey(name string, key []byte) error
	ListSSHKeys() error
	ListSSHKeysForName(name string) error
	RemoveSSHKeyByName(name string) error
	RemoveSSHKeyByFingerprint(fingerprint string) error
}

const (
	sshKeysAddCmd               = "ssh-keys:add %s"
	sshKeysListCmd              = "ssh-keys:list --format json"
	sshKeysListForNameCmd       = "ssh-keys:list --format json %s"
	sshKeysRemoveFingerprintCmd = "ssh-keys:remove --fingerprint %s"
	sshKeysRemoveNameCmd        = "ssh-keys:remove %s"
)

func (c *DefaultClient) AddSSHKey(name string, key []byte) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ListSSHKeys() error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ListSSHKeysForName(name string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) RemoveSSHKeyByName(name string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) RemoveSSHKeyByFingerprint(fingerprint string) error {
	//TODO implement me
	panic("implement me")
}
