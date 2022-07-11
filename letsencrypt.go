package dokku

type letsEncryptManager interface {
	LetsEncryptAutoRenewApp(appName string) error
	LetsEncryptAutoRenew() error

	LetsEncryptCleanup(appName string) error
	GetLetsEncryptCronJobEnabled() (bool, error)
	SetLetsEncryptCronJobEnabled(enabled bool) error

	GetAppLetsEncryptEnabled(appName string) (bool, error)
	SetAppLetsEncryptEnabled(appName string, enabled bool) error

	RevokeAppLetsEncryptCertificate(appName string) error

	GetLetsEncryptAppList() ([]LetsEncryptAppInfo, error)
}

type LetsEncryptAppInfo struct{}

const (
	letsEncryptAutoRenewCmd  = "letsencrypt:auto-renew %s"
	letsEncryptCleanupCmd    = "letsencrypt:cleanup %s"
	letsEncryptAddCronCmd    = "letsencrypt:cron-job --add"
	letsEncryptRemoveCronCmd = "letsencrypt:cron-job --remove"
	letsEncryptDisableAppCmd = "letsencrypt:disable %s"
	letsEncryptEnableAppCmd  = "letsencrypt:enable %s"
	letsEncryptListCmd       = "letsencrypt:list"
	letsEncryptRevokeAppCmd  = "letsencrypt:revoke %s"
)

func (c *DefaultClient) LetsEncryptAutoRenewApp(appName string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) LetsEncryptAutoRenew() error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) LetsEncryptCleanup(appName string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetLetsEncryptCronJobEnabled() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetLetsEncryptCronJobEnabled(enabled bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppLetsEncryptEnabled(appName string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetAppLetsEncryptEnabled(appName string, enabled bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) RevokeAppLetsEncryptCertificate(appName string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetLetsEncryptAppList() ([]LetsEncryptAppInfo, error) {
	//TODO implement me
	panic("implement me")
}
