package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
	"regexp"
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

func (c *DefaultClient) ListAppCronTasks(appName string) ([]CronTask, error) {
	cmd := fmt.Sprintf(cronAppListCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var multipleWhitespaceRe = regexp.MustCompile("\\s\\s+")
	var crons []CronTask
	for i, line := range strings.Split(out, "\n") {
		if i == 0 {
			continue
		}
		cols := multipleWhitespaceRe.Split(line, 3)
		if len(cols) < 3 {
			return nil, fmt.Errorf("failed to parse cron line: '%s'", line)
		}
		crons = append(crons, CronTask{
			ID:       cols[0],
			Schedule: cols[1],
			Command:  cols[2],
		})
	}

	return crons, nil
}

func (c *DefaultClient) GetAppCronReport(appName string) (*AppCronReport, error) {
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

func (c *DefaultClient) GetAllAppCronReport() (CronReport, error) {
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
