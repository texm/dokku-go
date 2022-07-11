package dokku

type dockerManager interface {
	DockerCleanup(appName string) error
	DockerCleanupAll() error

	GetAppDockerOptionsReport(appName string) (*AppDockerOptionsReport, error)
	GetGlobalDockerOptionsReport() (DockerOptionsReport, error)

	AddAppPhaseDockerOption(appName string, phase string, option string) error
	AddAppPhasesDockerOption(appName string, phases []string, option string) error
	ClearAppPhaseDockerOptions(appName string, phase string) error
	RemoveAppPhaseDockerOption(appName string, phase string, option string) error
	RemoveAppPhasesDockerOption(appName string, phases []string, option string) error

	LoginDockerRegistry(server string, username string, password string) error
	GetAppDockerRegistryReport(appName string) (*AppDockerRegistryReport, error)
	GetDockerRegistryReport() (DockerRegistryReport, error)
	SetAppDockerRegistryProperty(appName string, property string, value string) error
	ClearAppDockerRegistryProperty(appName string, property string) error

	RunAppCommand(appName string, cmd string, env map[string]string) (string, error)
	ListAppRunContainers(appName string) ([]string, error)
}

type AppDockerOptionsReport struct{}
type DockerOptionsReport map[string]*AppDockerOptionsReport

type AppDockerRegistryReport struct{}
type DockerRegistryReport map[string]AppDockerRegistryReport

const (
	cleanupCmd = "cleanup [<app>]"

	dockerOptionsAddCmd    = "docker-options:add <app> <phase(s)> OPTION"
	dockerOptionsClearCmd  = "docker-options:clear <app> [phase(s)]"
	dockerOptionsRemoveCmd = "docker-options:remove <app> <phase(s)> OPTION"
	dockerOptionsReportCmd = "docker-options:report [<app>] [<flag>]"

	dockerRegistryLoginCmd       = "registry:login [--password-stdin] <server> <username> [<password>]"
	dockerRegistryReportCmd      = "registry:report [<app>] [<flag>]"
	dockerRegistrySetPropertyCmd = "registry:set <app> <property> (<value>)"

	dockerRunCmd         = "run [-e|--env KEY=VALUE] [--no-tty] <app> <cmd>"
	dockerRunDetachedCmd = "run:detached [-e|-env KEY=VALUE] [--no-tty] <app> <cmd>"
	dockerRunListCmd     = "run:list <app>"
)

func (c *DefaultClient) DockerCleanup(appName string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) DockerCleanupAll() error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppDockerOptionsReport(appName string) (*AppDockerOptionsReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetGlobalDockerOptionsReport() (DockerOptionsReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) AddAppPhaseDockerOption(appName string, phase string, option string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) AddAppPhasesDockerOption(appName string, phases []string, option string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ClearAppPhaseDockerOptions(appName string, phase string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) RemoveAppPhaseDockerOption(appName string, phase string, option string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) RemoveAppPhasesDockerOption(appName string, phases []string, option string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) LoginDockerRegistry(server string, username string, password string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppDockerRegistryReport(appName string) (*AppDockerRegistryReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetDockerRegistryReport() (DockerRegistryReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetAppDockerRegistryProperty(appName string, property string, value string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ClearAppDockerRegistryProperty(appName string, property string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) RunAppCommand(appName string, cmd string, env map[string]string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ListAppRunContainers(appName string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}
