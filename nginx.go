package dokku

import (
	"errors"
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
	"strings"
)

type nginxManager interface {
	GetAppNginxConfig(appName string) (string, error)

	GetAppNginxAccessLogs(appName string) (string, error)
	GetAppNginxErrorLogs(appName string) (string, error)

	GetAppNginxReport(appName string) (*AppNginxReport, error)
	GetGlobalNginxReport() (NginxReport, error)

	ValidateAllNginxConfig(clean bool) error
	ValidateAppNginxConfig(appName string, clean bool) error

	SetAppNginxProperty(appName string, property NginxProperty, value string) error
}

type AppNginxReport struct {
	AccessLogFormat       string `dokku:"Nginx access log format"`
	AccessLogPath         string `dokku:"Nginx access log path"`
	BindAddressIPv4       string `dokku:"Nginx bind address ipv4"`
	BindAddressIPv6       string `dokku:"Nginx bind address ipv6"`
	ClientMaxBodySize     int    `dokku:"Nginx client max body size"`
	DisableCustomConfig   bool   `dokku:"Nginx disable custom config"`
	ErrorLogPath          string `dokku:"Nginx error log path"`
	GlobalHSTS            bool   `dokku:"Nginx global hsts"`
	ComputedHSTS          bool   `dokku:"Nginx computed hsts"`
	HSTS                  bool   `dokku:"Nginx hsts"`
	HSTSIncludeSubdomains bool   `dokku:"Nginx hsts include subdomains"`
	HSTSMaxAge            int    `dokku:"Nginx hsts max age"`
	HSTSPreload           bool   `dokku:"Nginx hsts preload"`
	ProxyBufferSize       int    `dokku:"Nginx proxy buffer size"`
	ProxyBuffering        string `dokku:"Nginx proxy buffering"`
	ProxyBuffers          string `dokku:"Nginx proxy buffers"`
	ProxyBusyBuffersSize  int    `dokku:"Nginx proxy busy buffers size"`
	ProxyReadTimeout      string `dokku:"Nginx proxy read timeout"`
	LastVisitedAt         string `dokku:"Nginx last visited at"`
	XForwardedForValue    string `dokku:"Nginx x forwarded for value"`
	XForwardedPortValue   string `dokku:"Nginx x forwarded port value"`
	XForwardedProtoValue  string `dokku:"Nginx x forwarded proto value"`
	XForwardedSSL         bool   `dokku:"Nginx x forwarded ssl"`
}
type NginxReport map[string]*AppNginxReport

type NginxProperty string

const (
	NginxPropertyBindAddressIPv4 = NginxProperty("bind-address-ipv4")
	NginxPropertyBindAddressIPv6 = NginxProperty("bind-address-ipv6")
	NginxPropertyHSTSHeader      = NginxProperty("hsts")
	NginxPropertyAccessLogPath   = NginxProperty("access-log-path")
	// check /etc/nginx/conf.d/00-log-formats.conf
	NginxPropertyAccessLogFormat   = NginxProperty("access-log-format")
	NginxPropertyErrorLogPath      = NginxProperty("error-log-path")
	NginxPropertyProxyReadTimeout  = NginxProperty("proxy-read-timeout")
	NginxPropertyClientMaxBodySize = NginxProperty("client-max-body-size")
	NginxDisableCustomConfig       = NginxProperty("disable-custom-config")
)

var (
	nginxNoConfigMsgPrefix = "!     No nginx.conf exists for"
	NginxNoConfigErr       = errors.New("no nginx.conf exists for app")
)

const (
	nginxAccessLogsCmd     = "nginx:access-logs %s"
	nginxErrorLogsCmd      = "nginx:error-logs %s"
	nginxReportCmd         = "nginx:report %s"
	nginxSetPropertyCmd    = "nginx:set %s %s %s"
	nginxShowConfigCmd     = "nginx:show-config %s"
	nginxValidateConfigCmd = "nginx:validate-config %s"
)

func (c *DefaultClient) GetAppNginxConfig(appName string) (string, error) {
	cmd := fmt.Sprintf(nginxShowConfigCmd, appName)
	out, err := c.Exec(cmd)
	if strings.HasPrefix(out, nginxNoConfigMsgPrefix) {
		return "", NginxNoConfigErr
	}
	if err != nil {
		return "", err
	}
	return out, nil
}

func (c *DefaultClient) GetAppNginxAccessLogs(appName string) (string, error) {
	cmd := fmt.Sprintf(nginxAccessLogsCmd, appName)
	return c.Exec(cmd)
}

func (c *DefaultClient) GetAppNginxErrorLogs(appName string) (string, error) {
	cmd := fmt.Sprintf(nginxErrorLogsCmd, appName)
	return c.Exec(cmd)
}

func (c *DefaultClient) GetAppNginxReport(appName string) (*AppNginxReport, error) {
	cmd := fmt.Sprintf(nginxReportCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report AppNginxReport
	if err := reports.ParseInto(out, &report); err != nil {
		return nil, err
	}

	return &report, nil
}

func (c *DefaultClient) GetGlobalNginxReport() (NginxReport, error) {
	cmd := fmt.Sprintf(nginxReportCmd, "")
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report NginxReport
	if err := reports.ParseIntoMap(out, &report); err != nil {
		return nil, err
	}

	return report, nil
}

func (c *DefaultClient) ValidateAllNginxConfig(clean bool) error {
	return c.ValidateAppNginxConfig("", clean)
}

func (c *DefaultClient) ValidateAppNginxConfig(appName string, clean bool) error {
	cmd := fmt.Sprintf(nginxValidateConfigCmd, appName)
	if clean {
		cmd += " --clean"
	}
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) SetAppNginxProperty(appName string, property NginxProperty, value string) error {
	cmd := fmt.Sprintf(nginxSetPropertyCmd, appName, property, value)
	_, err := c.Exec(cmd)
	return err
}
