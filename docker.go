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

/*
cleanup [<app>]                                                                              Cleans up exited/dead Docker containers and removes dangling images

docker-options:add <app> <phase(s)> OPTION                                                   Add Docker option to app for phase (comma separated phase list)
docker-options:clear <app> [phase(s)]                                                        Clear a docker options from application with an optional phase (comma separated phase list)
docker-options:remove <app> <phase(s)> OPTION                                                Remove Docker option from app for phase (comma separated phase list)
docker-options:report [<app>] [<flag>]                                                       Displays a docker options report for one or more apps

registry:login [--password-stdin] <server> <username> [<password>]                           Login to a docker registry
registry:report [<app>] [<flag>]                                                             Displays a registry report for one or more apps
registry:set <app> <property> (<value>)                                                      Set or clear a registry property for an app

run [-e|--env KEY=VALUE] [--no-tty] <app> <cmd>                                              Run a command in a new container using the current app image
run:detached [-e|-env KEY=VALUE] [--no-tty] <app> <cmd>                                      Run a command in a new detached container using the current app image
run:list <app>                                                                               List all run containers for an app
*/
