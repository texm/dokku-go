package dokku

type nginxManager interface {
	GetAppNginxConfig(appName string) (string, error)

	GetAppNginxAccessLogs(appName string) (string, error)
	GetAppNginxErrorLogs(appName string) (string, error)

	GetAppNginxReport(appName string) (*AppNginxReport, error)
	GetGlobalNginxReport() (NginxReport, error)

	ValidateAllNginxConfig(clean bool) error
	ValidateAppNginxConfig(appName string, clean bool) error

	SetAppNginxProperty(appName string, property string, value string) error
}

type AppNginxReport struct{}
type NginxReport map[string]AppNginxReport

/*
nginx:access-logs <app> [-t]                                                                 Show the nginx access logs for an application (-t follows)
nginx:error-logs <app> [-t]                                                                  Show the nginx error logs for an application (-t follows)
nginx:report [<app>] [<flag>]                                                                Displays an nginx report for one or more apps
nginx:set <app> <property> (<value>)                                                         Set or clear an nginx property for an app
nginx:show-config <app>                                                                      Display app nginx config
nginx:validate-config [<app>] [--clean]                                                      Validates and optionally cleans up invalid nginx configurations
*/
