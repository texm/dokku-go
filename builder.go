package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
	"strings"
)

type builderManager interface {
	GetAppBuilderDockerfileReport(appName string) (*AppBuilderDockerfileReport, error)
	SetAppBuilderDockerfileProperty(appName string, property DockerfileProperty, value string) error

	GetAppBuilderPackReport(appName string) (*AppBuilderPackReport, error)
	SetAppBuilderPackProperty(appName string, property BuildpackProperty, value string) error

	GetAppBuilderReport(appName string) (*AppBuilderReport, error)
	SetAppBuilderProperty(appName string, property BuilderProperty, value string) error

	AddAppBuildpack(appName string, buildpack string) error
	ClearAppBuildpacks(appName string) error
	ListAppBuildpacks(appName string) ([]string, error)
	RemoveAppBuildpack(appName string, buildpack string) error
	GetAppBuildpacksReport(appName string) (*AppBuildpacksReport, error)
	SetAppBuildpack(appName string, buildpack string) error
	SetAppBuildpacksProperty(appName string, property BuildpackProperty, value string) error
	SetGlobalBuildpacksProperty(property BuildpackProperty, value string) error
}

type AppBuilderDockerfileReport struct {
	DockerfilePath         string `dokku:"Builder dockerfile dockerfile path"`
	ComputedDockerfilePath string `dokku:"Builder dockerfile computed dockerfile path"`
	GlobalDockerfilePath   string `dokku:"Builder dockerfile global dockerfile path"`
}
type AppBuilderPackReport struct {
	ProjectTOMLPath         string `dokku:"Builder pack projecttoml path"`
	ComputedProjectTOMLPath string `dokku:"Builder pack computed projecttoml path"`
	GlobalProjectTOMLPath   string `dokku:"Builder pack global projecttoml path"`
}
type AppBuilderReport struct {
	BuildDir         string `dokku:"Builder build dir"`
	ComputedBuildDir string `dokku:"Builder computed build dir"`
	GlobalBuildDir   string `dokku:"Builder global build dir"`

	SelectedBuilder         string `dokku:"Builder selected"`
	ComputedSelectedBuilder string `dokku:"Builder computed selected"`
	GlobalSelectedBuilder   string `dokku:"Builder global selected"`
}
type AppBuildpacksReport struct {
	Stack         string `dokku:"Buildpacks stack"`
	ComputedStack string `dokku:"Buildpacks computed stack"`
	GlobalStack   string `dokku:"Buildpacks global stack"`

	List string `dokku:"Buildpacks list"`
}

type BuilderProperty string
type BuildpackProperty string
type DockerfileProperty string

const (
	BuilderPropertySelected = BuilderProperty("selected")
	BuilderPropertyBuildDir = BuilderProperty("build-dir")

	BuildpackPropertyProjectTomlPath = BuildpackProperty("projecttoml-path")
	BuildpackPropertyStackBuilder    = BuildpackProperty("stack")

	DockerfilePropertyPath = DockerfileProperty("dockerfile-path")
)

const (
	builderDockerfileReportCmd      = "builder-dockerfile:report %s"
	builderDockerfileSetPropertyCmd = "builder-dockerfile:set %s %s %s"

	builderPackReportCmd      = "builder-pack:report %s"
	builderPackSetPropertyCmd = "builder-pack:set %s %s %s"

	builderReportCmd      = "builder:report %s"
	builderSetPropertyCmd = "builder:set %s %s %s"

	buildpacksAddCmd         = "buildpacks:add --index %d %s %s"
	buildpacksClearCmd       = "buildpacks:clear %s"
	buildpacksListCmd        = "buildpacks:list %s"
	buildpacksRemoveCmd      = "buildpacks:remove %s %s"
	buildpacksReportCmd      = "buildpacks:report %s"
	buildpacksSetCmd         = "buildpacks:set --index %d %s %s"
	buildpacksSetPropertyCmd = "buildpacks:set-property %s %s %s"
)

func (c *BaseClient) GetAppBuilderDockerfileReport(appName string) (*AppBuilderDockerfileReport, error) {
	cmd := fmt.Sprintf(builderDockerfileReportCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report AppBuilderDockerfileReport
	if err := reports.ParseInto(out, &report); err != nil {
		return nil, err
	}

	return &report, err
}

func (c *BaseClient) SetAppBuilderDockerfileProperty(appName string, property DockerfileProperty, value string) error {
	cmd := fmt.Sprintf(builderDockerfileSetPropertyCmd, appName, property, value)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GetAppBuilderPackReport(appName string) (*AppBuilderPackReport, error) {
	cmd := fmt.Sprintf(builderPackReportCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report AppBuilderPackReport
	if err := reports.ParseInto(out, &report); err != nil {
		return nil, err
	}

	return &report, err
}

func (c *BaseClient) SetAppBuilderPackProperty(appName string, property BuildpackProperty, value string) error {
	cmd := fmt.Sprintf(builderPackSetPropertyCmd, appName, property, value)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GetAppBuilderReport(appName string) (*AppBuilderReport, error) {
	cmd := fmt.Sprintf(builderReportCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report AppBuilderReport
	if err := reports.ParseInto(out, &report); err != nil {
		return nil, err
	}

	return &report, err
}

func (c *BaseClient) SetAppBuilderProperty(appName string, property BuilderProperty, value string) error {
	cmd := fmt.Sprintf(builderSetPropertyCmd, appName, property, value)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) AddAppBuildpackAtIndex(appName string, buildpack string, index int) error {
	cmd := fmt.Sprintf(buildpacksAddCmd, index, appName, buildpack)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) AddAppBuildpack(appName string, buildpack string) error {
	return c.AddAppBuildpackAtIndex(appName, buildpack, 1)
}

func (c *BaseClient) ClearAppBuildpacks(appName string) error {
	cmd := fmt.Sprintf(buildpacksClearCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) ListAppBuildpacks(appName string) ([]string, error) {
	cmd := fmt.Sprintf(buildpacksListCmd, appName)
	out, err := c.Exec(cmd)

	var packs []string
	for i, line := range strings.Split(out, "\n") {
		if i == 0 {
			continue
		}
		packs = append(packs, line)
	}

	return packs, err
}

func (c *BaseClient) RemoveAppBuildpack(appName string, buildpack string) error {
	cmd := fmt.Sprintf(buildpacksRemoveCmd, appName, buildpack)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GetAppBuildpacksReport(appName string) (*AppBuildpacksReport, error) {
	cmd := fmt.Sprintf(buildpacksReportCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report AppBuildpacksReport
	if err := reports.ParseInto(out, &report); err != nil {
		return nil, err
	}

	return &report, err
}

func (c *BaseClient) SetAppBuildpackIndex(appName string, buildpack string, index int) error {
	cmd := fmt.Sprintf(buildpacksSetCmd, index, appName, buildpack)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) SetAppBuildpack(appName string, buildpack string) error {
	return c.SetAppBuildpackIndex(appName, buildpack, 1)
}

func (c *BaseClient) SetAppBuildpacksProperty(appName string, property BuildpackProperty, value string) error {
	cmd := fmt.Sprintf(buildpacksSetPropertyCmd, appName, property, value)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) SetGlobalBuildpacksProperty(property BuildpackProperty, value string) error {
	return c.SetAppBuildpacksProperty("--global", property, value)
}
