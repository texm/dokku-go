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

const (
	proxyBuildConfigCmd  = "proxy:build-config [--parallel count] [--all|<app>]"
	proxyClearConfigCmd  = "proxy:clear-config [--all|<app>]"
	proxyDisableAppCmd   = "proxy:disable <app>"
	proxyEnableAppCmd    = "proxy:enable <app>"
	proxyPortsCmd        = "proxy:ports <app>"
	proxyPortsAddCmd     = "proxy:ports-add <app> [<scheme>:<host-port>:<container-port>...]"
	proxyPortsClearCmd   = "proxy:ports-clear <app>"
	proxyPortsRemoveCmd  = "proxy:ports-remove <app> [<host-port>|<scheme>:<host-port>:<container-port>...]"
	proxyPortsSetCmd     = "proxy:ports-set <app> [<scheme>:<host-port>:<container-port>...]"
	proxyPortsReportCmd  = "proxy:report [<app>] [<flag>]"
	proxyPortsSetTypeCmd = "proxy:set <app> <proxy-type>"
)

func (c *DefaultClient) BuildAllProxyConfig(parallel *ParallelismOptions) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) BuildAppProxyConfig(appName string, parallel *ParallelismOptions) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ClearAllProxyConfig() error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ClearAppProxyConfig(appName string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppProxyReport(appName string) (*AppProxyReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAllAppProxyReport() (ProxyReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetAppProxyEnabled(appName string, enabled bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppProxyPortMappings(appName string) (AppProxyPortMappings, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) AddAppProxyPort(appName string, port ProxyPortMapping) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) AddAppProxyPorts(appName string, ports []ProxyPortMapping) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ClearAppProxyPorts(appName string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) RemoveAppProxyPort(appName string, port ProxyPortMapping) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) RemoveAppProxyPorts(appName string, ports []ProxyPortMapping) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetAppProxyPorts(appName string, ports []ProxyPortMapping) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetAppProxyType(appName string, proxyType string) error {
	//TODO implement me
	panic("implement me")
}
