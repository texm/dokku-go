package dokku

import (
	"fmt"
	"regexp"
	"strings"
)

type pluginManager interface {
	/*
		EnablePlugin(plugin string) error
		DisablePlugin(plugin string) error

		CheckPluginInstalled(plugin string) (bool, error)
		InstallPlugin(options PluginInstallOptions) error
		InstallPluginDependencies() error
		UninstallPlugin(plugin string) error
		UpdatePlugin(plugin string) error
		UpdatePlugins() error

		TriggerPluginHook(hookArgs []string) error
	*/

	ListPlugins() ([]PluginInfo, error)
}

type PluginInfo struct {
	Name        string
	Version     string
	Enabled     bool
	Description string
}

const (
	pluginInstalledCmd           = "plugin:installed %s"
	pluginDisableCmd             = "plugin:disable %s"
	pluginEnableCmd              = "plugin:enable %s"
	pluginInstallGitCmd          = "plugin:install <git-url> --committish <tag|branch|commit> --name <custom-plugin-name>"
	pluginInstallDependenciesCmd = "plugin:install-dependencies"
	pluginListCmd                = "plugin:list"
	pluginTriggerCmd             = "plugin:trigger %s"
	pluginUninstallCmd           = "plugin:uninstall %s"
	pluginUpdateCmd              = "plugin:update %s %s"
)

var multipleWhitespaceRe = regexp.MustCompile("\\s+")

func (c *DefaultClient) ListPlugins() ([]PluginInfo, error) {
	out, err := c.Exec(pluginListCmd)
	lines := strings.Split(out, "\n")
	plugins := make([]PluginInfo, len(lines))
	for i, line := range lines {
		line = strings.TrimSpace(line)
		cols := multipleWhitespaceRe.Split(line, 4)
		if len(cols) < 4 {
			return nil, fmt.Errorf("error parsing plugin list line: %s", line)
		}

		plugins[i] = PluginInfo{
			Name:        cols[0],
			Version:     cols[1],
			Enabled:     cols[2] == "enabled",
			Description: cols[3],
		}
	}

	return plugins, err
}

/*
func (c *DefaultClient) CheckPluginInstalled(plugin string) (bool, error) {

}

func (c *DefaultClient) EnablePlugin(plugin string) error {
	cmd := fmt.Sprintf(pluginEnableCmd, plugin)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) DisablePlugin(plugin string) error {
	cmd := fmt.Sprintf(pluginDisableCmd, plugin)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) InstallPlugin(options PluginInstallOptions) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) InstallPluginDependencies() error {
	_, err := c.Exec(pluginInstallDependenciesCmd)
	return err
}

func (c *DefaultClient) UninstallPlugin(plugin string) error {
	cmd := fmt.Sprintf(pluginUninstallCmd, plugin)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) UpdatePlugin(plugin string) error {
	cmd := fmt.Sprintf(pluginUpdateCmd, plugin)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) UpdatePlugins() error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) TriggerPluginHook(hookArgs []string) error {
	//TODO implement me
	panic("implement me")
}
*/
