package dokku

import "fmt"

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
	cmd := fmt.Sprintf(cronAppListCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var crons []CronTask
	fmt.Println(out)

	return crons, nil
}

func (c *DefaultClient) GetAppCronReport(appName string) (*AppCronReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAllAppCronReport() (CronReport, error) {
	//TODO implement me
	panic("implement me")
}
