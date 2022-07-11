package dokku

type cronManager interface {
	ListAppCronTasks(appName string) ([]CronTask, error)
	GetAppCronReport(appName string) (*AppCronReport, error)
	GetAllAppCronReport() (CronReport, error)
}

type CronTask struct{}
type AppCronReport struct{}
type CronReport map[string]AppCronReport

const (
	cronAppListCmd   = "cron:list %s"
	cronReportCmd    = "cron:report"
	cronAppReportCmd = "cron:report %s"
)

func (c *DefaultClient) ListAppCronTasks(appName string) ([]CronTask, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppCronReport(appName string) (*AppCronReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAllAppCronReport() (CronReport, error) {
	//TODO implement me
	panic("implement me")
}
