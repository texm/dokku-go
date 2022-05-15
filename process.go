package dokku

import "fmt"

const (
	psInspectCommand = "ps:inspect %s"
	psRebuildCommand = "ps:rebuild [--parallel count] [--all|<app>]"
	psReportCommand  = "ps:report [<app>] [<flag>]"
	psRestartCommand = "ps:restart [--parallel count] [--all|<app>] [<process-name>]"
	psRestoreCommand = "ps:restore [<app>]"
	psScaleCommand   = "ps:scale [--skip-deploy] <app> <proc>=<count> [<proc>=<count>...]"
	psSetCommand     = "ps:set <app> <key> <value>"
	psStartCommand   = "ps:start [--parallel count] [--all|<app>]"
	psStopCommand    = "ps:stop [--parallel count] [--all|<app>]"
)

func (c *Client) GetProcessInfo(appName string) error {
	cmd := fmt.Sprintf(psInspectCommand, appName)
	output, err := c.exec(cmd)
	if err != nil {
		return err
	}

	fmt.Println(output)

	return nil
}
