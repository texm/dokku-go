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
