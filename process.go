package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
	"strings"
)

const (
	psInspectCommand            = "ps:inspect %s"
	psRebuildCommand            = "ps:rebuild [--parallel count] [--all|<app>]"
	psReportCommand             = "ps:report"
	psReportAppCommand          = "ps:report %s"
	psReportAppWithFlagsCommand = "ps:report %s %s"
	psRestartCommand            = "ps:restart [--parallel count] [--all|<app>] [<process-name>]"
	psRestoreCommand            = "ps:restore [<app>]"
	psScaleCommand              = "ps:scale [--skip-deploy] <app> <proc>=<count> [<proc>=<count>...]"
	psSetCommand                = "ps:set <app> <key> <value>"
	psStartCommand              = "ps:start [--parallel count] [--all|<app>]"
	psStopCommand               = "ps:stop [--parallel count] [--all|<app>]"
)

func (c *DefaultClient) GetProcessInfo(appName string) error {
	cmd := fmt.Sprintf(psInspectCommand, appName)
	output, err := c.Exec(cmd)
	if err != nil {
		if strings.HasPrefix(output, "\"docker container inspect\" requires at least 1 argument.") {
			return AppNotDeployedError
		}
		return err
	}

	return NotImplementedError
}

type ProcessReport struct {
	Deployed             bool   `json:"deployed" dokku:"Deployed"`
	Processes            int    `json:"processes" dokku:"Processes"`
	CanScale             bool   `json:"can_scale" dokku:"Ps can scale"`
	ComputedProcfilePath string `json:"computed_procfile_path" dokku:"Ps computed procfile path"`
	GlobalProcfilePath   string `json:"global_procfile_path" dokku:"Ps global procfile path"`
	ProcfilePath         string `json:"procfile_path" dokku:"Ps procfile path"`
	RestartPolicy        string `json:"restart_policy" dokku:"Ps restart policy"`
	Restore              bool   `json:"restore" dokku:"Restore"`
	Running              bool   `json:"running" dokku:"Running"`
}

func (c *DefaultClient) GetAppProcessReport(appName string) (*ProcessReport, error) {
	cmd := fmt.Sprintf(psReportAppCommand, appName)
	output, err := c.Exec(cmd)

	if err != nil {
		return nil, err
	}

	report := ProcessReport{}
	if err := reports.ParseInto(output, &report); err != nil {
		return nil, err
	}

	return &report, nil
}

type ProcessesReport map[string]*ProcessReport

func (c *DefaultClient) GetAllProcessReport() (ProcessesReport, error) {
	output, err := c.Exec(psReportCommand)
	report := ProcessesReport{}

	if err == NoDeployedAppsError {
		return report, nil
	} else if err != nil {
		return nil, err
	}

	if err := reports.ParseIntoMap(output, &report); err != nil {
		return nil, err
	}

	return report, nil
}
