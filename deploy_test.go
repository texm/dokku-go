package dokku

import (
	"fmt"
)

func (s *DokkuTestSuite) TestDockerImageDeploy() {
	r := s.Require()
	var err error

	testAppName := "test-deploy-app"
	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	err = s.Client.SetAppDeployChecksEnabled(testAppName, false)
	fmt.Printf("checks err:%s\n", err)

	_, err = s.Client.DeployAppFromDockerImage(testAppName, "dokku/node-js-getting-started:latest")
	r.NoError(err)
}
