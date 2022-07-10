package dokku

type proxyManager interface {
	BuildAllProxyConfig(parallel *ParallelismOptions) error
	BuildAppProxyConfig(appName string, parallel *ParallelismOptions) error
	ClearAllProxyConfig() error
	ClearAppProxyConfig(appName string) error

	GetAppProxyReport(appName string) (*AppProxyReport, error)
	GetAllAppProxyReport() (ProxyReport, error)

	SetAppProxyEnabled(appName string, enabled bool) error

	GetAppProxyPortMappings(appName string) (AppProxyPortMappings, error)

	AddAppProxyPort(appName string, port ProxyPortMapping) error
	AddAppProxyPorts(appName string, ports []ProxyPortMapping) error

	ClearAppProxyPorts(appName string) error
	RemoveAppProxyPort(appName string, port ProxyPortMapping) error
	RemoveAppProxyPorts(appName string, ports []ProxyPortMapping) error

	SetAppProxyPorts(appName string, ports []ProxyPortMapping) error
	SetAppProxyType(appName string, proxyType string) error
}

type ProxyPortMapping struct{}
type AppProxyPortMappings struct{}

type AppProxyReport struct{}
type ProxyReport map[string]*AppProxyReport

/*
proxy:build-config [--parallel count] [--all|<app>]                                          (Re)builds config for a given app
proxy:clear-config [--all|<app>]                                                             Clears config for a given app
proxy:disable <app>                                                                          Disable proxy for app
proxy:enable <app>                                                                           Enable proxy for app
proxy:ports <app>                                                                            List proxy port mappings for app
proxy:ports-add <app> [<scheme>:<host-port>:<container-port>...]                             Add proxy port mappings to an app
proxy:ports-clear <app>                                                                      Clear all proxy port mappings for an app
proxy:ports-remove <app> [<host-port>|<scheme>:<host-port>:<container-port>...]              Remove specific proxy port mappings from an app
proxy:ports-set <app> [<scheme>:<host-port>:<container-port>...]                             Set proxy port mappings for an app
proxy:report [<app>] [<flag>]                                                                Displays a proxy report for one or more apps
proxy:set <app> <proxy-type>                                                                 Set proxy type for app
*/
