package dokku

type pluginManager interface {
	SetPluginEnabled(plugin string, enabled bool) error

	InstallPlugin(options PluginInstallOptions) error
	InstallPluginDependencies(coreOnly bool) error
	UninstallPlugin(plugin string) error
	UpdatePlugin(plugin string) error
	UpdatePlugins() error

	ListPlugins() ([]string, error)

	TriggerPluginHook(hookArgs []string) error
}

type PluginInstallOptions struct{}

const (
	pluginDisableCmd             = "plugin:disable %s"
	pluginEnableCmd              = "plugin:enable %s"
	pluginInstallGitCmd          = "plugin:install <git-url> --committish <tag|branch|commit> --name <custom-plugin-name>"
	pluginInstallDependenciesCmd = "plugin:install-dependencies"
	pluginListCmd                = "plugin:list"
	pluginTriggerCmd             = "plugin:trigger %s"
	pluginUninstallCmd           = "plugin:uninstall %s"
	pluginUpdateCmd              = "plugin:update %s %s"
)

func (c *DefaultClient) SetPluginEnabled(plugin string, enabled bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) InstallPlugin(options PluginInstallOptions) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) InstallPluginDependencies(coreOnly bool) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) UninstallPlugin(plugin string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) UpdatePlugin(plugin string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) UpdatePlugins() error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ListPlugins() ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) TriggerPluginHook(hookArgs []string) error {
	//TODO implement me
	panic("implement me")
}
