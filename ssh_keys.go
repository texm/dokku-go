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
