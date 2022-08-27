package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
	"regexp"
	"strings"
)

type proxyManager interface {
	BuildAllProxyConfig(parallel *ParallelismOptions) error
	BuildAppProxyConfig(appName string, parallel *ParallelismOptions) error
	ClearAllProxyConfig() error
	ClearAppProxyConfig(appName string) error

	GetAppProxyReport(appName string) (*AppProxyReport, error)
	GetAllAppProxyReport() (ProxyReport, error)

	SetAppProxyEnabled(appName string) error
	SetAppProxyDisabled(appName string) error

	GetAppProxyPortMappings(appName string) ([]ProxyPortMapping, error)

	AddAppProxyPort(appName string, port ProxyPortMapping) error
	AddAppProxyPorts(appName string, ports []ProxyPortMapping) error

	ClearAppProxyPorts(appName string) error
	RemoveAppProxyPort(appName string, port ProxyPortMapping) error
	RemoveAppProxyPorts(appName string, ports []ProxyPortMapping) error

	SetAppProxyPorts(appName string, ports []ProxyPortMapping) error
	SetAppProxyType(appName string, proxyType ProxyType) error
}

type ProxyPortMapping struct {
	Scheme        string
	HostPort      string
	ContainerPort string
}

func (m *ProxyPortMapping) String() string {
	return fmt.Sprintf("%s:%s:%s", m.Scheme, m.HostPort, m.ContainerPort)
}

func concatPortList(ports []ProxyPortMapping) string {
	portStrings := make([]string, len(ports))
	for i, port := range ports {
		portStrings[i] = port.String()
	}
	return strings.Join(portStrings, " ")
}

type AppProxyReport struct {
	Enabled bool   `dokku:"Proxy enabled"`
	PortMap string `dokku:"Proxy port map"`
	Type    string `dokku:"Proxy type"`
}
type ProxyReport map[string]*AppProxyReport

var (
	proxyNoPortMappingsMsg = "!     No port mappings configured for app"
)

type ProxyType string

const (
	ProxyTypeNginx = ProxyType("nginx")
)

const (
	proxyBuildConfigCmd = "proxy:build-config --parallel %d %s"
	proxyClearConfigCmd = "proxy:clear-config %s"
	proxyDisableAppCmd  = "proxy:disable %s"
	proxyEnableAppCmd   = "proxy:enable %s"
	proxyPortsCmd       = "proxy:ports %s"
	proxyPortsAddCmd    = "proxy:ports-add %s %s"
	proxyPortsClearCmd  = "proxy:ports-clear %s"
	proxyPortsRemoveCmd = "proxy:ports-remove %s %s"
	proxyPortsSetCmd    = "proxy:ports-set %s %s"
	proxyReportCmd      = "proxy:report %s"
	proxySetTypeCmd     = "proxy:set %s %s"
)

func (c *BaseClient) BuildAllProxyConfig(parallel *ParallelismOptions) error {
	return c.BuildAppProxyConfig("--all", parallel)
}

func (c *BaseClient) BuildAppProxyConfig(appName string, parallel *ParallelismOptions) error {
	cmd := fmt.Sprintf(proxyBuildConfigCmd, getParallelism(parallel), appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) ClearAllProxyConfig() error {
	return c.ClearAppProxyConfig("--all")
}

func (c *BaseClient) ClearAppProxyConfig(appName string) error {
	cmd := fmt.Sprintf(proxyClearConfigCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GetAppProxyReport(appName string) (*AppProxyReport, error) {
	cmd := fmt.Sprintf(proxyReportCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report AppProxyReport
	if err := reports.ParseInto(out, &report); err != nil {
		return nil, err
	}

	return &report, nil
}

func (c *BaseClient) GetAllAppProxyReport() (ProxyReport, error) {
	cmd := fmt.Sprintf(proxyReportCmd, "")
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report ProxyReport
	if err := reports.ParseInto(out, &report); err != nil {
		return nil, err
	}

	return report, nil
}

func (c *BaseClient) SetAppProxyEnabled(appName string) error {
	cmd := fmt.Sprintf(proxyEnableAppCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) SetAppProxyDisabled(appName string) error {
	cmd := fmt.Sprintf(proxyDisableAppCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GetAppProxyPortMappings(appName string) ([]ProxyPortMapping, error) {
	cmd := fmt.Sprintf(proxyPortsCmd, appName)
	out, err := c.Exec(cmd)
	if out == proxyNoPortMappingsMsg {
		return []ProxyPortMapping{}, nil
	}
	if err != nil {
		return nil, err
	}

	var multipleWhitespaceRe = regexp.MustCompile("\\s\\s+")

	var mappings []ProxyPortMapping
	for i, line := range strings.Split(out, "\n") {
		if i <= 1 {
			continue
		}
		cols := multipleWhitespaceRe.Split(line, 3)
		if len(cols) < 3 {
			return nil, fmt.Errorf("error parsing port map line '%s'", line)
		}
		mappings = append(mappings, ProxyPortMapping{
			Scheme:        cols[0],
			HostPort:      cols[1],
			ContainerPort: cols[2],
		})
	}

	return mappings, nil
}

func (c *BaseClient) AddAppProxyPort(appName string, port ProxyPortMapping) error {
	cmd := fmt.Sprintf(proxyPortsAddCmd, appName, port.String())
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) AddAppProxyPorts(appName string, ports []ProxyPortMapping) error {
	cmd := fmt.Sprintf(proxyPortsAddCmd, appName, concatPortList(ports))
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) ClearAppProxyPorts(appName string) error {
	cmd := fmt.Sprintf(proxyPortsClearCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) RemoveAppProxyPort(appName string, port ProxyPortMapping) error {
	cmd := fmt.Sprintf(proxyPortsRemoveCmd, appName, port.String())
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) RemoveAppProxyPorts(appName string, ports []ProxyPortMapping) error {
	cmd := fmt.Sprintf(proxyPortsRemoveCmd, appName, concatPortList(ports))
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) SetAppProxyPorts(appName string, ports []ProxyPortMapping) error {
	cmd := fmt.Sprintf(proxyPortsSetCmd, appName, concatPortList(ports))
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) SetAppProxyType(appName string, proxyType ProxyType) error {
	cmd := fmt.Sprintf(proxySetTypeCmd, appName, proxyType)
	_, err := c.Exec(cmd)
	return err
}
