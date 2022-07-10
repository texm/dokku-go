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

/*
letsencrypt:auto-renew <app>            Auto-renew app if renewal is necessary
letsencrypt:auto-renew                  Auto-renew all apps secured by letsencrypt if renewal is necessary
letsencrypt:cleanup <app>               Remove stale certificate directories for app
letsencrypt:cron-job [--add --remove]   Add or remove a cron job that periodically calls auto-renew.
letsencrypt:disable <app>               Disable letsencrypt for an app
letsencrypt:enable <app>                Enable or renew letsencrypt for an app
letsencrypt:help                        Display letsencrypt help
letsencrypt:list                        List letsencrypt-secured apps with certificate expiry times
letsencrypt:revoke <app>                Revoke letsencrypt certificate for app
*/
