package dokku

type schedulerManager interface {
	GetAppSchedulerDockerLocalReport(appName string) (AppDockerLocalSchedulerReport, error)
	GetSchedulerDockerLocalReport() (DockerLocalReport, error)

	GetAppSchedulerReport(appName string) (AppSchedulerReport, error)
	GetSchedulerReport() (SchedulerReport, error)

	SetAppSchedulerProperty(appName string, property string, value string) error
}

type AppDockerLocalSchedulerReport struct{}
type DockerLocalReport map[string]*AppDockerLocalSchedulerReport

type AppSchedulerReport struct{}
type SchedulerReport map[string]*AppSchedulerReport

const (
	schedulerDockerLocalReportCmd      = "scheduler-docker-local:report %s"
	schedulerDockerLocalSetPropertyCmd = "scheduler-docker-local:set %s %s %s"

	schedulerReportCmd      = "scheduler:report %s"
	schedulerSetPropertyCmd = "scheduler:set %s %s %s"
)

func (c *DefaultClient) GetAppSchedulerDockerLocalReport(appName string) (AppDockerLocalSchedulerReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetSchedulerDockerLocalReport() (DockerLocalReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppSchedulerReport(appName string) (AppSchedulerReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetSchedulerReport() (SchedulerReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetAppSchedulerProperty(appName string, property string, value string) error {
	//TODO implement me
	panic("implement me")
}
