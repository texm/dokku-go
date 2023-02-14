package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
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

	SetAppLetsEncryptProperty(appName string, property LetsEncryptProperty, value string) error
	SetGlobalLetsEncryptProperty(property LetsEncryptProperty, value string) error
	ClearAppLetsEncryptProperty(appName string, property LetsEncryptProperty) error
	ClearGlobalLetsEncryptProperty(property LetsEncryptProperty) error

	RevokeAppLetsEncryptCertificate(appName string) error

	GetLetsEncryptAppList() ([]LetsEncryptAppInfo, error)

	GetLetsEncryptAppReport(appName string) (*LetsEncryptAppReport, error)
}

type LetsEncryptProperty string

const (
	LetsEncryptPropertyDnsProvider    = LetsEncryptProperty("dns-provider")
	LetsEncryptPropertyEmail          = LetsEncryptProperty("email")
	LetsEncryptPropertyGracePeriod    = LetsEncryptProperty("graceperiod")
	LetsEncryptPropertyLegoDockerArgs = LetsEncryptProperty("lego-docker-args")
	LetsEncryptPropertyServer         = LetsEncryptProperty("server")
)

type LetsEncryptAppReport struct {
	Active                 bool   `json:"active" dokku:"Letsencrypt active"`
	AutoRenew              bool   `json:"autorenew" dokku:"Letsencrypt autorenew"`
	ComputedDnsProvider    string `json:"computed_dns_provider" dokku:"Letsencrypt computed dns provider"`
	GlobalDnsProvider      string `json:"global_dns_provider" dokku:"Letsencrypt global dns provider"`
	DnsProvider            string `json:"dns_provider" dokku:"Letsencrypt dns provider"`
	ComputedEmail          string `json:"computed_email" dokku:"Letsencrypt computed email"`
	GlobalEmail            string `json:"global_email" dokku:"Letsencrypt global email"`
	Email                  string `json:"email" dokku:"Letsencrypt email"`
	Expiration             int    `json:"expiration" dokku:"Letsencrypt expiration"`
	ComputedGracePeriod    int    `json:"computed_grace_period" dokku:"Letsencrypt computed graceperiod"`
	GlobalGracePeriod      int    `json:"global_grace_period" dokku:"Letsencrypt global graceperiod"`
	GracePeriod            int    `json:"grace_period" dokku:"Letsencrypt graceperiod"`
	ComputedLegoDockerArgs string `json:"computed_lego_docker_args" dokku:"Letsencrypt computed lego docker args"`
	GlobalLegoDockerArgs   string `json:"global_lego_docker_args" dokku:"Letsencrypt global lego docker args"`
	LegoDockerArgs         string `json:"lego_docker_args" dokku:"Letsencrypt lego docker args"`
	ComputedServer         string `json:"computed_server" dokku:"Letsencrypt computed server"`
	GlobalServer           string `json:"global_server" dokku:"Letsencrypt global server"`
	Server                 string `json:"server" dokku:"Letsencrypt server"`
}

type LetsEncryptAppInfo struct{}

const (
	letsEncryptActiveCmd        = "letsencrypt:active %s"
	letsEncryptAutoRenewCmd     = "letsencrypt:auto-renew %s"
	letsEncryptCleanupCmd       = "letsencrypt:cleanup %s"
	letsEncryptCronCmd          = "letsencrypt:cron-job %s"
	letsEncryptDisableAppCmd    = "letsencrypt:disable %s"
	letsEncryptEnableAppCmd     = "letsencrypt:enable %s"
	letsEncryptListCmd          = "letsencrypt:list"
	letsEncryptRevokeAppCmd     = "letsencrypt:revoke %s"
	letsEncryptAppReportCmd     = "letsencrypt:report %s"
	letsEncryptSetPropertyCmd   = "letsencrypt:set %s %s %s"
	letsEncryptClearPropertyCmd = "letsencrypt:set %s %s"
)

func exitCodeReturn(err error) (bool, error) {
	exitErr, ok := err.(*ExitCodeError)
	if ok && exitErr.ExitStatus() == 1 {
		return false, nil
	}
	return false, err
}

func (c *BaseClient) LetsEncryptAutoRenewApp(appName string) error {
	cmd := fmt.Sprintf(letsEncryptAutoRenewCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) LetsEncryptAutoRenew() error {
	cmd := fmt.Sprintf(letsEncryptAutoRenewCmd, "")
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) LetsEncryptCleanup(appName string) error {
	cmd := fmt.Sprintf(letsEncryptCleanupCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GetLetsEncryptCronJobEnabled() (bool, error) {
	// https://github.com/dokku/dokku-letsencrypt/issues/221
	cmd := fmt.Sprintf(letsEncryptCronCmd, "")
	out, err := c.Exec(cmd)
	fmt.Println(out)
	return false, err
}

func (c *BaseClient) AddLetsEncryptCronJob() error {
	cmd := fmt.Sprintf(letsEncryptCronCmd, "--add")
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) RemoveLetsEncryptCronJob() error {
	cmd := fmt.Sprintf(letsEncryptCronCmd, "--remove")
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GetAppLetsEncryptEnabled(appName string) (bool, error) {
	cmd := fmt.Sprintf(letsEncryptActiveCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil && out == "" {
		return exitCodeReturn(err)
	}
	return true, nil
}

func (c *BaseClient) EnableAppLetsEncrypt(appName string) error {
	cmd := fmt.Sprintf(letsEncryptEnableAppCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) DisableAppLetsEncrypt(appName string) error {
	cmd := fmt.Sprintf(letsEncryptDisableAppCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) RevokeAppLetsEncryptCertificate(appName string) error {
	cmd := fmt.Sprintf(letsEncryptRevokeAppCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GetLetsEncryptAppList() ([]LetsEncryptAppInfo, error) {
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

func (c *BaseClient) GetLetsEncryptAppReport(appName string) (*LetsEncryptAppReport, error) {
	cmd := fmt.Sprintf(letsEncryptAppReportCmd, appName)
	output, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report LetsEncryptAppReport
	if err := reports.ParseInto(output, &report); err != nil {
		return nil, fmt.Errorf("failed to parse report: %w", err)
	}

	return &report, nil
}

func (c *BaseClient) SetAppLetsEncryptProperty(appName string, property LetsEncryptProperty, value string) error {
	cmd := fmt.Sprintf(letsEncryptSetPropertyCmd, appName, property, value)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) ClearAppLetsEncryptProperty(appName string, property LetsEncryptProperty) error {
	return c.SetAppLetsEncryptProperty(appName, property, "")
}

func (c *BaseClient) SetGlobalLetsEncryptProperty(property LetsEncryptProperty, value string) error {
	return c.SetAppLetsEncryptProperty("--global", property, value)
}

func (c *BaseClient) ClearGlobalLetsEncryptProperty(property LetsEncryptProperty) error {
	return c.SetGlobalLetsEncryptProperty(property, "")
}
