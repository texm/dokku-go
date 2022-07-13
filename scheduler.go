package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
)

type schedulerManager interface {
	GetAppSchedulerDockerLocalReport(appName string) (*AppSchedulerDockerLocalReport, error)
	GetSchedulerDockerLocalReport() (SchedulerDockerLocalReport, error)
	SetSchedulerDockerLocalProperty(appName string, property DockerLocalSchedulerProperty, value string) error

	GetAppSchedulerReport(appName string) (*AppSchedulerReport, error)
	GetSchedulerReport() (SchedulerReport, error)
	SetAppSchedulerProperty(appName string, property SchedulerProperty, value string) error
}

type AppSchedulerDockerLocalReport struct {
	DisableChown          bool `dokku:"Scheduler docker local disable chown"`
	ParallelScheduleCount int  `dokku:"Scheduler docker local parallel schedule count"`
}
type SchedulerDockerLocalReport map[string]*AppSchedulerDockerLocalReport

type AppSchedulerReport struct {
	ComputedSelectedScheduler string `dokku:"Scheduler computed selected"`
	GlobalSelectedScheduler   string `dokku:"Scheduler global selected"`
	SelectedScheduler         string `dokku:"Scheduler selected"`
}
type SchedulerReport map[string]*AppSchedulerReport

type SchedulerProperty string
type DockerLocalSchedulerProperty string

const (
	SchedulerPropertySelected = SchedulerProperty("selected")

	DockerLocalSchedulerPropertyDisableChown          = DockerLocalSchedulerProperty("disable-chown")
	DockerLocalSchedulerPropertyParallelScheduleCount = DockerLocalSchedulerProperty("parallel-schedule-count")
)

const (
	schedulerDockerLocalReportCmd      = "scheduler-docker-local:report %s"
	schedulerDockerLocalSetPropertyCmd = "scheduler-docker-local:set %s %s %s"

	schedulerReportCmd      = "scheduler:report %s"
	schedulerSetPropertyCmd = "scheduler:set %s %s %s"
)

func (c *DefaultClient) GetAppSchedulerDockerLocalReport(appName string) (*AppSchedulerDockerLocalReport, error) {
	cmd := fmt.Sprintf(schedulerDockerLocalReportCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report *AppSchedulerDockerLocalReport
	if err := reports.ParseInto(out, &report); err != nil {
		return nil, err
	}

	return report, nil
}

func (c *DefaultClient) GetSchedulerDockerLocalReport() (SchedulerDockerLocalReport, error) {
	cmd := fmt.Sprintf(schedulerDockerLocalReportCmd, "")
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report SchedulerDockerLocalReport
	if err := reports.ParseIntoMap(out, &report); err != nil {
		return nil, err
	}

	return report, nil
}

func (c *DefaultClient) SetSchedulerDockerLocalProperty(appName string, property DockerLocalSchedulerProperty, value string) error {
	cmd := fmt.Sprintf(schedulerDockerLocalSetPropertyCmd, appName, property, value)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) GetAppSchedulerReport(appName string) (*AppSchedulerReport, error) {
	cmd := fmt.Sprintf(schedulerReportCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report *AppSchedulerReport
	if err := reports.ParseInto(out, &report); err != nil {
		return nil, err
	}

	return report, nil
}

func (c *DefaultClient) GetSchedulerReport() (SchedulerReport, error) {
	cmd := fmt.Sprintf(schedulerReportCmd, "")
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report SchedulerReport
	if err := reports.ParseIntoMap(out, &report); err != nil {
		return nil, err
	}

	return report, nil
}

func (c *DefaultClient) SetAppSchedulerProperty(appName string, property SchedulerProperty, value string) error {
	cmd := fmt.Sprintf(schedulerSetPropertyCmd, appName, property, value)
	_, err := c.Exec(cmd)
	return err
}
