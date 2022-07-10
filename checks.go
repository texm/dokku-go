package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
	"strings"
)

const (
	checksEnableCmd         = "checks:enable %s"
	checksEnableProcessCmd  = "checks:enable %s %s"
	checksDisableCmd        = "checks:disable %s"
	checksDisableProcessCmd = "checks:disable %s %s"
	checksSkipCmd           = "checks:skip %s"
	checksSkipProcessCmd    = "checks:skip %s %s"
	checksReportCmd         = "checks:report %s"
)

func (c *DefaultClient) EnableAppDeployChecks(appName string) error {
	cmd := fmt.Sprintf(checksEnableCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) EnableAppProcessesDeployChecks(appName string, processes []string) error {
	cmd := fmt.Sprintf(checksEnableProcessCmd, appName, strings.Join(processes, ","))
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) DisableAppDeployChecks(appName string) error {
	cmd := fmt.Sprintf(checksDisableCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) DisableAppProcessesDeployChecks(appName string, processes []string) error {
	cmd := fmt.Sprintf(checksDisableProcessCmd, appName, strings.Join(processes, ","))
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) SetAppDeployChecksSkipped(appName string) error {
	cmd := fmt.Sprintf(checksSkipCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) SetAppProcessesDeployChecksSkipped(appName string, processes []string) error {
	cmd := fmt.Sprintf(checksSkipProcessCmd, appName, strings.Join(processes, ","))
	_, err := c.Exec(cmd)
	return err
}

type appChecksRawReport struct {
	Disabled string `dokku:"Checks disabled list"`
	Skipped  string `dokku:"Checks skipped list"`
}

type AppChecksReport struct {
	AllDisabled       bool     `json:"all_disabled"`
	AllSkipped        bool     `json:"all_skipped"`
	DisabledProcesses []string `json:"disabled_processes"`
	SkippedProcesses  []string `json:"skipped_processes"`
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

func (c *DefaultClient) GetAppDeployChecksReport(appName string) (*AppChecksReport, error) {
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

type ChecksReport map[string]*AppChecksReport

func (c *DefaultClient) GetDeployChecksReport() (ChecksReport, error) {
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
