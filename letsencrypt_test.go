package dokku

import (
	"context"
	"fmt"
	"github.com/texm/dokku-go/internal/testutils"
)

func setupLetsEncryptPlugin(dc *testutils.DokkuContainer) error {
	ctx := context.Background()

	lePluginURL := "https://github.com/dokku/dokku-letsencrypt.git"
	cmd := []string{"dokku", "plugin:install", lePluginURL}
	code, err := dc.Exec(ctx, cmd)
	if err != nil {
		return fmt.Errorf("failed to install letsencrypt plugin: %s", err)
	}
	if code != 0 {
		return fmt.Errorf("failed to install letsencrypt plugin: code %d", code)
	}

	return nil
}

func (s *DokkuTestSuite) TestLetsEncrypt() {
	r := s.Require()

	r.NoError(setupLetsEncryptPlugin(s.Dokku))

	appName := "test-letsencrypt-app"
	r.NoError(s.Client.CreateApp(appName))

	active, err := s.Client.GetAppLetsEncryptEnabled(appName)
	r.NoError(err)
	r.False(active)

	_, err = s.Client.GetLetsEncryptAppList()
	r.NoError(err)

	_, err = s.Client.GetLetsEncryptCronJobEnabled()
	r.NoError(err)
}
