package dokku

type configManager interface {
	GetDokkuVersion() (string, error)

	SetAppJsonProperty(appName string, property string, value string) error
	GetAppJsonReport(appName string) (*AppJsonReport, error)
	GetAllAppJsonReport() (AppJsonReports, error)

	GetGlobalConfig() (AppConfig, error)
	GetAppConfig(appName string) (AppConfig, error)

	ClearAppConfig(appName string, restart bool) error
	ClearGlobalConfig(restart bool) error

	ExportAppConfig(appName string, format string, merged bool) (string, error)
	ExportGlobalConfig(format string, merged bool) (string, error)

	GetAppConfigValue(appName string, key string) (string, error)
	GetGlobalConfigValue(key string) (string, error)

	SetAppConfigValue(appName string, key string, value string, encoded bool, restart bool) error
	UnsetAppConfigValue(appName string, key string, restart bool) error
	SetGlobalConfigValue(key string, value string, encoded bool, restart bool) error
	UnsetGlobalConfigValue(key string, restart bool) error

	SetAppConfigValues(appName string, config map[string]string, encoded bool, restart bool) error
	UnsetAppConfigValues(appName string, keys []string, restart bool) error
	SetGlobalConfigValues(config map[string]string, encoded bool, restart bool) error
	UnsetGlobalConfigValues(keys []string, restart bool) error

	GetAppConfigKeys(appName string, merged bool) ([]string, error)
	GetGlobalConfigKeys(merged bool) ([]string, error)

	ShowAppConfig(appName string, merged bool) ([]string, error)
	ShowGlobalConfig(merged bool) ([]string, error)
}

type AppJsonReport struct{}
type AppJsonReports map[string]AppJsonReport

type AppConfig struct{}

const (
	versionCmd = "version"

	appJsonReportCmd      = "app-json:report [<app>] [<flag>]"
	appJsonSetPropertyCmd = "app-json:set <app> <property> (<value>)"

	configPrintCmd  = "config (<app>|--global)"
	configBundleCmd = "config:bundle [--merged] (<app>|--global)"
	configClearCmd  = "config:clear [--no-restart] (<app>|--global)"
	configExportCmd = "config:export [--format=FORMAT] [--merged] (<app>|--global)"
	configGetCmd    = "config:get [--quoted] (<app>|--global) KEY"
	configKeysCmd   = "config:keys [--merged] (<app>|--global)"
	configSetCmd    = "config:set [--encoded] [--no-restart] (<app>|--global) KEY1=VALUE1 [KEY2=VALUE2 ...]"
	configShowCmd   = "config:show [--merged] (<app>|--global)"
	configUnsetCmd  = "config:unset [--no-restart] (<app>|--global) KEY1 [KEY2 ...]"
)

func (c *DefaultClient) GetDokkuVersion() (string, error) {
	return c.Exec(versionCmd)
}

func (c *DefaultClient) SetAppJsonProperty(appName string, property string, value string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppJsonReport(appName string) (*AppJsonReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAllAppJsonReport() (AppJsonReports, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetGlobalConfig() (AppConfig, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppConfig(appName string) (AppConfig, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ClearAppConfig(appName string, restart bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ClearGlobalConfig(restart bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ExportAppConfig(appName string, format string, merged bool) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ExportGlobalConfig(format string, merged bool) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppConfigValue(appName string, key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetGlobalConfigValue(key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetAppConfigValue(appName string, key string, value string, encoded bool, restart bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) UnsetAppConfigValue(appName string, key string, restart bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetGlobalConfigValue(key string, value string, encoded bool, restart bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) UnsetGlobalConfigValue(key string, restart bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetAppConfigValues(appName string, config map[string]string, encoded bool, restart bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) UnsetAppConfigValues(appName string, keys []string, restart bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetGlobalConfigValues(config map[string]string, encoded bool, restart bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) UnsetGlobalConfigValues(keys []string, restart bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppConfigKeys(appName string, merged bool) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetGlobalConfigKeys(merged bool) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ShowAppConfig(appName string, merged bool) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ShowGlobalConfig(merged bool) ([]string, error) {
	//TODO implement me
	panic("implement me")
}
