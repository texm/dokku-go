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

/*
builder-dockerfile:report [<app>] [<flag>]                                                   Displays a builder-dockerfile report for one or more apps
builder-dockerfile:set <app> <property> (<value>)                                            Set or clear a builder-dockerfile property for an app

builder-pack:report [<app>] [<flag>]                                                         Displays a builder-pack report for one or more apps
builder-pack:set <app> <property> (<value>)                                                  Set or clear a builder-pack property for an app

builder:report [<app>] [<flag>]                                                              Displays a builder report for one or more apps
builder:set <app> <property> (<value>)                                                       Set or clear a builder property for an app

buildpacks:add [--index 1] <app> <buildpack>                                                 Add new app buildpack while inserting into list of buildpacks if necessary
buildpacks:clear <app>                                                                       Clear all buildpacks set on the app
buildpacks:list <app>                                                                        List all buildpacks for an app
buildpacks:remove <app> <buildpack>                                                          Remove a buildpack set on the app
buildpacks:report [<app>] [<flag>]                                                           Displays a buildpack report for one or more apps
buildpacks:set [--index 1] <app> <buildpack>                                                 Set new app buildpack at a given position defaulting to the first buildpack if no index is specified
buildpacks:set-property [--global|<app>] <key> <value>                                       Set or clear a buildpacks property for an app
*/
