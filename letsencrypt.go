package dokku

import "fmt"

type letsEncryptManager interface {
	LetsEncryptAutoRenewApp(appName string) error
	LetsEncryptAutoRenew() error

	LetsEncryptCleanup(appName string) error
	GetLetsEncryptCronJobEnabled() (bool, error)

	AddLetsEncryptCronJob() error
	RemoveLetsEncryptCronJob() error

	GetAppLetsEncryptActive(appName string) (bool, error)
	GetAppLetsEncryptEnabled(appName string) (bool, error)
	SetAppLetsEncryptEnabled(appName string, enabled bool) error

	RevokeAppLetsEncryptCertificate(appName string) error

	GetLetsEncryptAppList() ([]LetsEncryptAppInfo, error)
}

type LetsEncryptAppInfo struct{}

const (
	letsEncryptActiveCmd     = "letsencrypt:active %s"
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
	cmd := fmt.Sprintf(letsEncryptAutoRenewCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) LetsEncryptAutoRenew() error {
	cmd := fmt.Sprintf(letsEncryptAutoRenewCmd, "")
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) LetsEncryptCleanup(appName string) error {
	cmd := fmt.Sprintf(letsEncryptCleanupCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) GetLetsEncryptCronJobEnabled() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) AddLetsEncryptCronJob() error {
	_, err := c.Exec(letsEncryptAddCronCmd)
	return err
}

func (c *DefaultClient) RemoveLetsEncryptCronJob() error {
	_, err := c.Exec(letsEncryptRemoveCronCmd)
	return err
}

func (c *DefaultClient) GetAppLetsEncryptActive(appName string) (bool, error) {
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
