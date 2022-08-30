package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
	"strings"
)

type networkManager interface {
	CreateNetwork(name string) error
	DestroyNetwork(name string) error
	CheckNetworkExists(name string) (bool, error)
	GetNetworkInfo(name string) (interface{}, error)
	ListNetworks() ([]string, error)
	RebuildNetwork(name string) error
	RebuildAllNetworks() error
	GetAppNetworkReport(appName string) (*AppNetworkReport, error)
	GetNetworkReport() (NetworkReport, error)
	SetAppNetworkProperty(appName string, property NetworkProperty, value string) error
	RemoveAppNetworkProperty(appName string, property NetworkProperty) error
	SetGlobalNetworkProperty(property NetworkProperty, value string) error
	RemoveGlobalNetworkProperty(property NetworkProperty) error

	// SetProperty Aliases
	// SetAppNetworkAttachPostCreate(appName string, network string)
	// SetAppNetworkAttachPostDeploy(appName string, network string)
	// SetAppNetworkInitial(appName string, network string)
	// SetGlobalNetworkAttachPostCreate(network string)
	// SetGlobalNetworkAttachPostDeploy(network string)
	// SetGlobalNetworkInitialNetwork(network string)
	// SetAppNetworkTLD(appName string, tld string)
	// SetGlobalNetworkTLD(tld string)
	// SetAppNetworkBindAllInterfaces(appName string, shouldBind bool)
	// SetGlobalNetworkBindAllInterfaces(shouldBind bool)
}

type AppNetworkReport struct {
	AttachPostCreate         string `dokku:"Network attach post create"`
	GlobalAttachPostCreate   string `dokku:"Network global attach post create"`
	ComputedAttachPostCreate string `dokku:"Network computed attach post create"`

	AttachPostDeploy         string `dokku:"Network attach post deploy"`
	GlobalAttachPostDeploy   string `dokku:"Network global attach post deploy"`
	ComputedAttachPostDeploy string `dokku:"Network computed attach post deploy"`

	BindAllInterfaces         bool `dokku:"Network bind all interfaces"`
	GlobalBindAllInterfaces   bool `dokku:"Network global bind all interfaces"`
	ComputedBindAllInterfaces bool `dokku:"Network computed bind all interfaces"`

	InitialNetwork         string `dokku:"Network initial network"`
	GlobalInitialNetwork   bool   `dokku:"Network global initial network"`
	ComputedInitialNetwork string `dokku:"Network computed initial network"`

	TLD         string `dokku:"Network tld"`
	GlobalTLD   string `dokku:"Network global tld"`
	ComputedTLD string `dokku:"Network computed tld"`

	WebListeners string `dokku:"Network web listeners"`
}

type NetworkReport map[string]*AppNetworkReport

type NetworkProperty string

const (
	NetworkPropertyStaticWebListener = NetworkProperty("static-web-listener")
	NetworkPropertyTLD               = NetworkProperty("tld")
	NetworkPropertyBindAllInterfaces = NetworkProperty("bind-all-interfaces")

	NetworkPropertyInitialNetwork   = NetworkProperty("initial-network")
	NetworkPropertyAttachPostCreate = NetworkProperty("attach-post-create")
	NetworkPropertyAttachPostDeploy = NetworkProperty("attach-post-deploy")
)

const (
	networkCreateCmd      = "network:create %s"
	networkDestroyCmd     = "network:destroy --force %s"
	networkExistsCmd      = "network:exists %s"
	networkInfoCmd        = "network:info %s"
	networkListCmd        = "network:list --quiet"
	networkRebuildCmd     = "network:rebuild %s"
	networkRebuildAllCmd  = "network:rebuildall"
	networkReportCmd      = "network:report %s"
	networkSetPropertyCmd = "network:set %s %s %s"
)

func (c *BaseClient) CreateNetwork(name string) error {
	cmd := fmt.Sprintf(networkCreateCmd, name)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) DestroyNetwork(name string) error {
	cmd := fmt.Sprintf(networkDestroyCmd, name)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) CheckNetworkExists(name string) (bool, error) {
	cmd := fmt.Sprintf(networkExistsCmd, name)
	out, err := c.Exec(cmd)
	if out == "Network does not exist" {
		return false, nil
	} else if out == "Network exists" {
		return true, nil
	}
	return false, err
}

func (c *BaseClient) GetNetworkInfo(name string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *BaseClient) ListNetworks() ([]string, error) {
	out, err := c.Exec(networkListCmd)
	if err != nil {
		return nil, err
	}
	return strings.Split(out, "\n"), nil
}

func (c *BaseClient) RebuildNetwork(name string) error {
	cmd := fmt.Sprintf(networkRebuildCmd, name)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) RebuildAllNetworks() error {
	_, err := c.Exec(networkRebuildAllCmd)
	return err
}

func (c *BaseClient) GetAppNetworkReport(appName string) (*AppNetworkReport, error) {
	cmd := fmt.Sprintf(networkReportCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report AppNetworkReport
	if err := reports.ParseInto(out, &report); err != nil {
		return nil, err
	}

	return &report, nil
}

func (c *BaseClient) GetNetworkReport() (NetworkReport, error) {
	cmd := fmt.Sprintf(networkReportCmd, "")
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	reportMap := NetworkReport{}
	if err := reports.ParseIntoMap(out, &reportMap); err != nil {
		return nil, err
	}

	return reportMap, nil
}

func (c *BaseClient) SetAppNetworkProperty(appName string, property NetworkProperty, value string) error {
	cmd := fmt.Sprintf(networkSetPropertyCmd, appName, property, value)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) RemoveAppNetworkProperty(appName string, property NetworkProperty) error {
	cmd := fmt.Sprintf(networkSetPropertyCmd, appName, property, "")
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) SetGlobalNetworkProperty(property NetworkProperty, value string) error {
	cmd := fmt.Sprintf(networkSetPropertyCmd, "--global", property, value)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) RemoveGlobalNetworkProperty(property NetworkProperty) error {
	cmd := fmt.Sprintf(networkSetPropertyCmd, "--global", property, "")
	_, err := c.Exec(cmd)
	return err
}
