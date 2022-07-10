package dokku

import "fmt"

const (
	enableChecksCmd  = "checks:enable %s"
	disableChecksCmd = "checks:disable %s"
	skipChecksCmd    = "checks:skip %s"
	reportChecksCmd  = "checks:report %s"
	deployImageCmd   = "git:from-image %s %s"
)

func (c *DefaultClient) SetAppDeployChecksEnabled(appName string, enabled bool) error {
	cmd := enableChecksCmd
	if !enabled {
		cmd = disableChecksCmd
	}
	_, err := c.Exec(fmt.Sprintf(cmd, appName))
	return err
}

func (c *DefaultClient) DeployAppFromDockerImage(appName, image string) (string, error) {
	cmd := fmt.Sprintf(deployImageCmd, appName, image)
	return c.Exec(cmd)
}
