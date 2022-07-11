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
type NginxReport map[string]*AppNginxReport

const (
	nginxAccessLogsCmd          = "nginx:access-logs %s"
	nginxErrorLogsCmd           = "nginx:error-logs %s"
	nginxReportCmd              = "nginx:report %s"
	nginxSetPropertyCmd         = "nginx:set %s %s %s"
	nginxShowConfigCmd          = "nginx:show-config %s"
	nginxValidateConfigCmd      = "nginx:validate-config %s"
	nginxValidateCleanConfigCmd = "nginx:validate-config %s --clean"
)

func (c *DefaultClient) GetAppNginxConfig(appName string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppNginxAccessLogs(appName string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppNginxErrorLogs(appName string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppNginxReport(appName string) (*AppNginxReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetGlobalNginxReport() (NginxReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ValidateAllNginxConfig(clean bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ValidateAppNginxConfig(appName string, clean bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetAppNginxProperty(appName string, property string, value string) error {
	//TODO implement me
	panic("implement me")
}
