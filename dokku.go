package dokku

type Client interface {
	commandExecutor

	appManager
	builderManager
	certsManager
	checksManager
	configManager
	cronManager
	dockerManager
	domainsManager
	gitManager
	letsEncryptManager
	logsManager
	networkManager
	nginxManager
	pluginManager
	processManager
	proxyManager
	resourceManager
	schedulerManager
	sshKeysManager
	storageManager
}
