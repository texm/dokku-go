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

/*
plugin:disable <name>                                                                        Disable an installed plugin (third-party only)
plugin:enable <name>                                                                         Enable a previously disabled plugin
plugin:install [--core|git-url [--committish tag|branch|commit|--name custom-plugin-name]]   Optionally download git-url (with custom tag/committish) & run install trigger for active plugins (or only core ones)
plugin:install-dependencies [--core]                                                         Run install-dependencies trigger for active plugins (or only core ones)
plugin:list                                                                                  Print active plugins
plugin:trigger <args...>                                                                     Trigger an arbitrary plugin hook
plugin:uninstall <name>                                                                      Uninstall a plugin (third-party only)
plugin:update [name [committish]]                                                            Optionally update named plugin from git (with custom tag/committish) & run update trigger for active plugins
*/
