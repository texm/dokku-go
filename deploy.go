package dokku

import "fmt"

const (
	deployImageCmd = "git:from-image %s %s"
)

func (c *DefaultClient) DeployAppFromDockerImage(appName, image string) (string, error) {
	cmd := fmt.Sprintf(deployImageCmd, appName, image)
	return c.exec(cmd)
}
