package dokku

type configManager interface {
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

/*
app-json:report [<app>] [<flag>]                                                             Displays a app-json report for one or more apps
app-json:set <app> <property> (<value>)                                                      Set or clear a app-json property for an app

config (<app>|--global)                                                                      Pretty-print an app or global environment
config:bundle [--merged] (<app>|--global)                                                    Bundle environment into tarfile
config:clear [--no-restart] (<app>|--global)                                                 Clears environment variables
config:export [--format=FORMAT] [--merged] (<app>|--global)                                  Export a global or app environment
config:get [--quoted] (<app>|--global) KEY                                                   Display a global or app-specific config value
config:keys [--merged] (<app>|--global)                                                      Show keys set in environment
config:set [--encoded] [--no-restart] (<app>|--global) KEY1=VALUE1 [KEY2=VALUE2 ...]         Set one or more config vars
config:show [--merged] (<app>|--global)                                                      Show keys set in environment
config:unset [--no-restart] (<app>|--global) KEY1 [KEY2 ...]                                 Unset one or more config vars
*/
