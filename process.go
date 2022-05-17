package dokku

import (
	"fmt"

	"github.com/texm/dokku-go/internal/reports"
)

const (
	psInspectCommand   = "ps:inspect %s"
	psRebuildCommand   = "ps:rebuild [--parallel count] [--all|<app>]"
	psReportCommand    = "ps:report"
	psAppReportCommand = "ps:report %s [<flag>]"
	psRestartCommand   = "ps:restart [--parallel count] [--all|<app>] [<process-name>]"
	psRestoreCommand   = "ps:restore [<app>]"
	psScaleCommand     = "ps:scale [--skip-deploy] <app> <proc>=<count> [<proc>=<count>...]"
	psSetCommand       = "ps:set <app> <key> <value>"
	psStartCommand     = "ps:start [--parallel count] [--all|<app>]"
	psStopCommand      = "ps:stop [--parallel count] [--all|<app>]"
)

func (c *DokkuClient) GetProcessInfo(appName string) error {
	cmd := fmt.Sprintf(psInspectCommand, appName)
	output, err := c.exec(cmd)
	if err != nil {
		return err
	}

	fmt.Println(output)

	return nil
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

type ProcessesReport map[string]*ProcessReport

func (c *DokkuClient) GetAllProcessReport() (ProcessesReport, error) {
	output, err := c.exec(psReportCommand)
	report := ProcessesReport{}

	if err == NoDeployedAppsError {
		return report, nil
	} else if err != nil {
		return nil, err
	}

	if err := reports.ParseInto(output, &report); err != nil {
		return nil, err
	}

	return report, nil
}
