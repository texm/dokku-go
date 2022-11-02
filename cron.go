package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
	"strings"
)

type cronManager interface {
	ListAppCronTasks(appName string) ([]CronTask, error)
	GetAppCronReport(appName string) (*AppCronReport, error)
	GetAllAppCronReport() (CronReport, error)
}

type CronTask struct {
	ID       string
	Schedule string
	Command  string
}
type AppCronReport struct {
	TaskCount int `dokku:"Cron task count"`
}
type CronReport map[string]*AppCronReport

const (
	cronAppListCmd   = "cron:list %s"
	cronReportCmd    = "cron:report"
	cronAppReportCmd = "cron:report %s"
)

func parseCronOutput(output string) ([]CronTask, error) {
	lines := strings.Split(output, "\n")
	columnLine := lines[0]
	scheduleIndex := strings.Index(columnLine, "Schedule")
	commandIndex := strings.Index(columnLine, "Command")

	crons := make([]CronTask, len(lines)-1)
	for i, line := range lines[1:] {
		id := strings.TrimSpace(line[:scheduleIndex])
		schedule := strings.TrimSpace(line[scheduleIndex:commandIndex])
		command := strings.TrimSpace(line[commandIndex:])

		crons[i] = CronTask{
			ID:       id,
			Schedule: schedule,
			Command:  command,
		}
	}
	return crons, nil
}

func (c *BaseClient) ListAppCronTasks(appName string) ([]CronTask, error) {
	cmd := fmt.Sprintf(cronAppListCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	return parseCronOutput(out)
}

func (c *BaseClient) GetAppCronReport(appName string) (*AppCronReport, error) {
	cmd := fmt.Sprintf(cronAppReportCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report AppCronReport
	if err := reports.ParseInto(out, &report); err != nil {
		return nil, err
	}

	return &report, nil
}

func (c *BaseClient) GetAllAppCronReport() (CronReport, error) {
	out, err := c.Exec(cronReportCmd)
	if err != nil {
		return nil, err
	}

	var report CronReport
	if err := reports.ParseIntoMap(out, &report); err != nil {
		return nil, err
	}

	return report, nil
}
