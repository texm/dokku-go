package dokku

import (
	"fmt"
	"strings"
)

type letsEncryptManager interface {
	LetsEncryptAutoRenewApp(appName string) error
	LetsEncryptAutoRenew() error

	LetsEncryptCleanup(appName string) error
	GetLetsEncryptCronJobEnabled() (bool, error)

	AddLetsEncryptCronJob() error
	RemoveLetsEncryptCronJob() error

	GetAppLetsEncryptEnabled(appName string) (bool, error)
	EnableAppLetsEncrypt(appName string) error
	DisableAppLetsEncrypt(appName string) error

	RevokeAppLetsEncryptCertificate(appName string) error

	GetLetsEncryptAppList() ([]LetsEncryptAppInfo, error)
}

type LetsEncryptAppInfo struct{}

const (
	letsEncryptActiveCmd     = "letsencrypt:active %s"
	letsEncryptAutoRenewCmd  = "letsencrypt:auto-renew %s"
	letsEncryptCleanupCmd    = "letsencrypt:cleanup %s"
	letsEncryptCronCmd       = "letsencrypt:cron-job %s"
	letsEncryptDisableAppCmd = "letsencrypt:disable %s"
	letsEncryptEnableAppCmd  = "letsencrypt:enable %s"
	letsEncryptListCmd       = "letsencrypt:list"
	letsEncryptRevokeAppCmd  = "letsencrypt:revoke %s"
)

func exitCodeReturn(err error) (bool, error) {
	exitErr, ok := err.(*ExitCodeError)
	if ok && exitErr.sshErr.ExitStatus() == 1 {
		return false, nil
	}
	return false, err
}

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
	// https://github.com/dokku/dokku-letsencrypt/issues/221
	cmd := fmt.Sprintf(letsEncryptCronCmd, "")
	out, err := c.Exec(cmd)
	fmt.Println(out)
	return false, err
}

func (c *DefaultClient) AddLetsEncryptCronJob() error {
	cmd := fmt.Sprintf(letsEncryptCronCmd, "--add")
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) RemoveLetsEncryptCronJob() error {
	cmd := fmt.Sprintf(letsEncryptCronCmd, "--remove")
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) GetAppLetsEncryptEnabled(appName string) (bool, error) {
	cmd := fmt.Sprintf(letsEncryptActiveCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil && out == "" {
		return exitCodeReturn(err)
	}
	return true, nil
}

func (c *DefaultClient) EnableAppLetsEncrypt(appName string) error {
	cmd := fmt.Sprintf(letsEncryptEnableAppCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) DisableAppLetsEncrypt(appName string) error {
	cmd := fmt.Sprintf(letsEncryptDisableAppCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) RevokeAppLetsEncryptCertificate(appName string) error {
	cmd := fmt.Sprintf(letsEncryptRevokeAppCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) GetLetsEncryptAppList() ([]LetsEncryptAppInfo, error) {
	out, err := c.Exec(letsEncryptListCmd)
	if err != nil {
		return nil, err
	}
	// var multipleWhitespaceRe = regexp.MustCompile("\\s+")
	var infoList []LetsEncryptAppInfo
	for i, line := range strings.Split(out, "\n") {
		if i == 0 {
			continue
		}
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		//rows := multipleWhitespaceRe.Split(line, 4)
		//infoList = append()
	}
	return infoList, nil
}
