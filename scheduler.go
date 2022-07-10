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

/*
scheduler-docker-local:report [<app>] [<flag>]                                               Displays a scheduler-docker-local report for one or more apps
scheduler-docker-local:set <app> <property> (<value>)                                        Set or clear a scheduler-docker-local property for an app

scheduler:report [<app>] [<flag>]                                                            Displays a scheduler report for one or more apps
scheduler:set <app> <property> (<value>)                                                     Set or clear a scheduler property for an app
*/
