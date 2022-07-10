package dokku

type networkManager interface {
	CreateNetwork(name string) error
	DestroyNetwork(name string) error
	CheckNetworkExists(name string) (bool, error)
	GetNetworkInfo(name string) (interface{}, error)
	ListNetworks() ([]string, error)
	RebuildNetwork(name string) error
	RebuildAllNetworks() error
	GetNetworkReport(name string) (interface{}, error)
	SetNetworkProperty(name string, property string, value string) error
}

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
