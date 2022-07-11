package dokku

type builderManager interface {
	GetAppBuilderDockerfileReport(appName string) (*AppBuilderDockerfileReport, error)
	SetAppBuilderDockerfileProperty(appName string, property string, value string) error

	GetAppBuilderPackReport(appName string) (*AppBuilderPackReport, error)
	SetAppBuilderPackProperty(appName string, property string, value string) error

	GetAppBuilderReport(appName string) (*AppBuilderReport, error)
	SetAppBuilderProperty(appName string, property string, value string) error

	AddAppBuildpack(appName string, buildpack string) error
	ClearAppBuildpacks(appName string) error
	ListAppBuildpacks(appName string) ([]string, error)
	RemoveAppBuildpack(appName string, buildpack string) error
	GetAppBuildpacksReport(appName string) (*AppBuildpacksReport, error)
	SetAppBuildpack(appName string, buildpack string) error
	SetAppBuildpacksProperty(appName string, property string, value string) error

	SetGlobalBuildpacksProperty(appName string, property string, value string) error
}

type AppBuilderDockerfileReport struct{}
type AppBuilderPackReport struct{}
type AppBuilderReport struct{}
type AppBuildpacksReport struct{}

const (
	builderDockerfileReportCmd      = "builder-dockerfile:report [<app>] [<flag>]"
	builderDockerfileSetPropertyCmd = "builder-dockerfile:set <app> <property> (<value>)"

	builderPackReportCmd      = "builder-pack:report [<app>] [<flag>]"
	builderPackSetPropertyCmd = "builder-pack:set <app> <property> (<value>)"

	builderReportCmd      = "builder:report [<app>] [<flag>]"
	builderSetPropertyCmd = "builder:set <app> <property> (<value>)"

	buildpacksAddCmd         = "buildpacks:add [--index 1] <app> <buildpack>"
	buildpacksClearCmd       = "buildpacks:clear <app>"
	buildpacksListCmd        = "buildpacks:list <app>"
	buildpacksRemoveCmd      = "buildpacks:remove <app> <buildpack>"
	buildpacksReportCmd      = "buildpacks:report [<app>] [<flag>]"
	buildpacksSetCmd         = "buildpacks:set [--index 1] <app> <buildpack>"
	buildpacksSetPropertyCmd = "buildpacks:set-property [--global|<app>] <key> <value>"
)

func (c *DefaultClient) GetAppBuilderDockerfileReport(appName string) (*AppBuilderDockerfileReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetAppBuilderDockerfileProperty(appName string, property string, value string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppBuilderPackReport(appName string) (*AppBuilderPackReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetAppBuilderPackProperty(appName string, property string, value string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppBuilderReport(appName string) (*AppBuilderReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetAppBuilderProperty(appName string, property string, value string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) AddAppBuildpack(appName string, buildpack string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ClearAppBuildpacks(appName string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ListAppBuildpacks(appName string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) RemoveAppBuildpack(appName string, buildpack string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppBuildpacksReport(appName string) (*AppBuildpacksReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetAppBuildpack(appName string, buildpack string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetAppBuildpacksProperty(appName string, property string, value string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetGlobalBuildpacksProperty(appName string, property string, value string) error {
	//TODO implement me
	panic("implement me")
}
