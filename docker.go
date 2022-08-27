package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
	"strings"
)

type dockerManager interface {
	DockerCleanup(appName string) error
	DockerCleanupAll() error

	GetAppDockerOptionsReport(appName string) (*AppDockerOptionsReport, error)
	GetGlobalDockerOptionsReport() (DockerOptionsReport, error)

	AddAppPhaseDockerOption(appName string, phase string, option string) error
	ClearAppPhaseDockerOptions(appName string, phase string) error
	RemoveAppPhaseDockerOption(appName string, phase string, option string) error

	LoginDockerRegistry(server string, username string, password string) error
	GetAppDockerRegistryReport(appName string) (*AppDockerRegistryReport, error)
	GetDockerRegistryReport() (DockerRegistryReport, error)
	SetAppDockerRegistryProperty(appName string, property DockerRegistryProperty, value string) error
	ClearAppDockerRegistryProperty(appName string, property DockerRegistryProperty) error

	RunAppCommand(appName string, cmd string, options *DockerRunOptions) (string, error)
	ListAppRunContainers(appName string) ([]string, error)
}

type AppDockerOptionsReport struct {
	BuildOptions  string `dokku:"Docker options build"`
	DeployOptions string `dokku:"Docker options deploy"`
	RunOptions    string `dokku:"Docker options run"`
}
type DockerOptionsReport map[string]*AppDockerOptionsReport

type AppDockerRegistryReport struct {
	ImageRepo         string `dokku:"Registry image repo"`
	ComputedImageRepo string `dokku:"Registry computed image repo"`

	PushOnRelease         string `dokku:"Registry push on release"`
	GlobalPushOnRelease   bool   `dokku:"Registry global push on release"`
	ComputedPushOnRelease bool   `dokku:"Registry computed push on release"`

	Server         string `dokku:"Registry server"`
	GlobalServer   string `dokku:"Registry global server"`
	ComputedServer string `dokku:"Registry computed server"`

	TagVersion string `dokku:"Registry tag version"`
}
type DockerRegistryReport map[string]AppDockerRegistryReport

type DockerRunOptions struct {
	Detached    bool
	Environment map[string]string
}

type DockerRegistryProperty string

const (
	DockerRegistryPropertyServer        = DockerRegistryProperty("server")
	DockerRegistryPropertyImageRepo     = DockerRegistryProperty("image-repo")
	DockerRegistryPropertyPushOnRelease = DockerRegistryProperty("push-on-release")
)

const (
	cleanupCmd = "cleanup %s"

	dockerOptionsAddCmd    = "docker-options:add %s %s %s"
	dockerOptionsClearCmd  = "docker-options:clear %s %s"
	dockerOptionsRemoveCmd = "docker-options:remove %s %s %s"
	dockerOptionsReportCmd = "docker-options:report %s"

	dockerRegistryLoginCmd       = "registry:login %s %s %s"
	dockerRegistryReportCmd      = "registry:report %s"
	dockerRegistrySetPropertyCmd = "registry:set %s %s %s"

	dockerRunCmd         = "run %s --no-tty %s %s"
	dockerRunDetachedCmd = "run:detached %s --no-tty %s %s"
	dockerRunListCmd     = "run:list %s"
)

func (c *BaseClient) DockerCleanup(appName string) error {
	cmd := fmt.Sprintf(cleanupCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) DockerCleanupAll() error {
	return c.DockerCleanup("")
}

func (c *BaseClient) GetAppDockerOptionsReport(appName string) (*AppDockerOptionsReport, error) {
	cmd := fmt.Sprintf(dockerOptionsReportCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report AppDockerOptionsReport
	if err := reports.ParseInto(out, &report); err != nil {
		return nil, err
	}

	return &report, nil
}

func (c *BaseClient) GetGlobalDockerOptionsReport() (DockerOptionsReport, error) {
	cmd := fmt.Sprintf(dockerOptionsReportCmd, "")
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report DockerOptionsReport
	if err := reports.ParseIntoMap(out, &report); err != nil {
		return nil, err
	}

	return report, nil
}

func (c *BaseClient) AddAppPhaseDockerOption(appName string, phase string, option string) error {
	cmd := fmt.Sprintf(dockerOptionsAddCmd, appName, phase, option)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) ClearAppPhaseDockerOptions(appName string, phase string) error {
	cmd := fmt.Sprintf(dockerOptionsClearCmd, appName, phase)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) RemoveAppPhaseDockerOption(appName string, phase string, option string) error {
	cmd := fmt.Sprintf(dockerOptionsRemoveCmd, appName, phase, option)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) LoginDockerRegistry(server string, username string, password string) error {
	cmd := fmt.Sprintf(dockerRegistryLoginCmd, server, username, password)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GetAppDockerRegistryReport(appName string) (*AppDockerRegistryReport, error) {
	cmd := fmt.Sprintf(dockerRegistryReportCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report *AppDockerRegistryReport
	if err := reports.ParseInto(out, &report); err != nil {
		return nil, err
	}

	return report, nil
}

func (c *BaseClient) GetDockerRegistryReport() (DockerRegistryReport, error) {
	cmd := fmt.Sprintf(dockerRegistryReportCmd, "")
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report DockerRegistryReport
	if err := reports.ParseIntoMap(out, &report); err != nil {
		return nil, err
	}

	return report, nil
}

func (c *BaseClient) SetAppDockerRegistryProperty(appName string, property DockerRegistryProperty, value string) error {
	cmd := fmt.Sprintf(dockerRegistrySetPropertyCmd, appName, property, value)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) ClearAppDockerRegistryProperty(appName string, property DockerRegistryProperty) error {
	return c.SetAppDockerRegistryProperty(appName, property, "")
}

func (c *BaseClient) RunAppCommand(appName string, runCmd string, options *DockerRunOptions) (string, error) {
	tpl := dockerRunCmd
	envArg := ""
	if options != nil {
		if options.Detached {
			tpl = dockerRunDetachedCmd
		}
		if options.Environment != nil {
			var env []string
			for key, val := range options.Environment {
				env = append(env, fmt.Sprintf("--env '%s=%s'", key, val))
			}
			envArg = strings.Join(env, " ")
		}
	}
	cmd := fmt.Sprintf(tpl, envArg, appName, runCmd)
	return c.Exec(cmd)
}

// TODO: implement
func (c *BaseClient) ListAppRunContainers(appName string) ([]string, error) {
	cmd := fmt.Sprintf(dockerRunListCmd, appName)
	_, err := c.Exec(cmd)

	var containers []string
	// =====> node-js-app run containers
	// NAMES                   COMMAND            CREATED
	// node-js-app.run.28689   "/exec sleep 15"   2 seconds ago

	return containers, err
}
