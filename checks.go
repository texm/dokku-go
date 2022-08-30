package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
	"strings"
)

type checksManager interface {
	GetDeployChecksReport() (ChecksReport, error)
	GetAppDeployChecksReport(appName string) (*AppChecksReport, error)
	EnableAppDeployChecks(appName string) error
	EnableAppProcessesDeployChecks(appName string, processes []string) error
	DisableAppDeployChecks(appName string) error
	DisableAppProcessesDeployChecks(appName string, processes []string) error
	SetAppDeployChecksSkipped(appName string) error
	SetAppProcessesDeployChecksSkipped(appName string, processes []string) error
}

type AppChecksReport struct {
	AllDisabled       bool     `json:"all_disabled"`
	AllSkipped        bool     `json:"all_skipped"`
	DisabledProcesses []string `json:"disabled_processes"`
	SkippedProcesses  []string `json:"skipped_processes"`
}

type ChecksReport map[string]*AppChecksReport

const (
	checksEnableCmd         = "checks:enable %s"
	checksEnableProcessCmd  = "checks:enable %s %s"
	checksDisableCmd        = "checks:disable %s"
	checksDisableProcessCmd = "checks:disable %s %s"
	checksSkipCmd           = "checks:skip %s"
	checksSkipProcessCmd    = "checks:skip %s %s"
	checksReportCmd         = "checks:report %s"
)

func (c *BaseClient) EnableAppDeployChecks(appName string) error {
	cmd := fmt.Sprintf(checksEnableCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) EnableAppProcessesDeployChecks(appName string, processes []string) error {
	cmd := fmt.Sprintf(checksEnableProcessCmd, appName, strings.Join(processes, ","))
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) DisableAppDeployChecks(appName string) error {
	cmd := fmt.Sprintf(checksDisableCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) DisableAppProcessesDeployChecks(appName string, processes []string) error {
	cmd := fmt.Sprintf(checksDisableProcessCmd, appName, strings.Join(processes, ","))
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) SetAppDeployChecksSkipped(appName string) error {
	cmd := fmt.Sprintf(checksSkipCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) SetAppProcessesDeployChecksSkipped(appName string, processes []string) error {
	cmd := fmt.Sprintf(checksSkipProcessCmd, appName, strings.Join(processes, ","))
	_, err := c.Exec(cmd)
	return err
}

type appChecksRawReport struct {
	Disabled string `dokku:"Checks disabled list"`
	Skipped  string `dokku:"Checks skipped list"`
}

func parseRawReport(r appChecksRawReport) *AppChecksReport {
	report := &AppChecksReport{
		AllDisabled:       r.Disabled == "_all_",
		AllSkipped:        r.Skipped == "_all_",
		DisabledProcesses: []string{},
		SkippedProcesses:  []string{},
	}
	if !report.AllDisabled && r.Disabled != "none" {
		report.DisabledProcesses = strings.Split(r.Disabled, ",")
	}
	if !report.AllSkipped && r.Skipped != "none" {
		report.SkippedProcesses = strings.Split(r.Skipped, ",")
	}

	return report
}

func (c *BaseClient) GetAppDeployChecksReport(appName string) (*AppChecksReport, error) {
	cmd := fmt.Sprintf(checksReportCmd, appName)
	output, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var rawReport appChecksRawReport
	if err := reports.ParseInto(output, &rawReport); err != nil {
		return nil, err
	}

	return parseRawReport(rawReport), nil
}

func (c *BaseClient) GetDeployChecksReport() (ChecksReport, error) {
	cmd := fmt.Sprintf(checksReportCmd, "")
	output, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	rawReports := map[string]appChecksRawReport{}
	if err := reports.ParseIntoMap(output, &rawReports); err != nil {
		return nil, err
	}

	report := ChecksReport{}
	for name, individualReport := range rawReports {
		report[name] = parseRawReport(individualReport)
	}

	return report, nil
}
