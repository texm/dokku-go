package dokku

import (
	"log"
)

func (s *DokkuTestSuite) TestDockerImageDeploy() {
	r := s.Require()
	var err error

	testAppName := "test-logs-app"
	err = s.Client.CreateApp(testAppName)
	r.NoError(err, "failed to create app")

	// checks:skip <app>

	image := "crccheck/hello-world"
	out, err := s.Client.DeployAppFromDockerImage(testAppName, image)
	r.NoError(err)
	log.Printf("command output: %s\n", out)
}
