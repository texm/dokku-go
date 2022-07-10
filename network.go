package dokku

const (
	networkCreateCmd      = "network:create %s"
	networkDestroyCmd     = "network:destroy %s"
	networkExistsCmd      = "network:exists %s"
	networkInfoCmd        = "network:info %s"
	networkListCmd        = "network:list"
	networkRebuildCmd     = "network:rebuild %s"
	networkRebuildAllCmd  = "network:rebuildall"
	networkReportCmd      = "network:report %s"
	networkSetPropertyCmd = "network:set %s %s %s"
)

func (c *DefaultClient) CreateNetwork(name string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) DestroyNetwork(name string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) CheckNetworkExists(name string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetNetworkInfo(name string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ListNetworks() ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) RebuildNetwork(name string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) RebuildAllNetworks() error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetNetworkReport(name string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetNetworkProperty(name string, property string, value string) error {
	//TODO implement me
	panic("implement me")
}
