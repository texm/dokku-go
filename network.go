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
	SetAppNetworkProperty(appName string, property string, value string) error
	RemoveAppNetworkProperty(appName string, property string) error
	SetGlobalNetworkProperty(property string, value string) error
	RemoveGlobalNetworkProperty(property string) error
}

type AppNetworkReport struct {
	AttachPostCreate          string `dokku:"Network attach post create"`
	AttachPostDeploy          string `dokku:"Network attach post deploy"`
	BindAllInterfaces         bool   `dokku:"Network bind all interfaces"`
	ComputedAttachPostCreate  string `dokku:"Network computed attach post create"`
	ComputedAttachPostDeploy  string `dokku:"Network computed attach post deploy"`
	ComputedBindAllInterfaces bool   `dokku:"Network computed bind all interfaces"`
	ComputedInitialNetwork    string `dokku:"Network computed initial network"`
	ComputedTLD               string `dokku:"Network computed tld"`
	GlobalAttachPostCreate    string `dokku:"Network global attach post create"`
	GlobalAttachPostDeploy    string `dokku:"Network global attach post deploy"`
	GlobalBindAllInterfaces   bool   `dokku:"Network global bind all interfaces"`
	GlobalInitialNetwork      bool   `dokku:"Network global initial network"`
	GlobalTLD                 string `dokku:"Network global tld"`
	InitialNetwork            string `dokku:"Network initial network"`
	TLD                       string `dokku:"Network tld"`
	WebListeners              string `dokku:"Network web listeners"`
}

type NetworkReport map[string]*AppNetworkReport

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

func (c *DefaultClient) CreateNetwork(name string) error {
	cmd := fmt.Sprintf(networkCreateCmd, name)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) DestroyNetwork(name string) error {
	cmd := fmt.Sprintf(networkDestroyCmd, name)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) CheckNetworkExists(name string) (bool, error) {
	cmd := fmt.Sprintf(networkExistsCmd, name)
	out, err := c.Exec(cmd)
	if out == "Network does not exist" {
		return false, nil
	} else if out == "Network exists" {
		return true, nil
	}
	return false, err
}

func (c *DefaultClient) GetNetworkInfo(name string) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ListNetworks() ([]string, error) {
	out, err := c.Exec(networkListCmd)
	if err != nil {
		return nil, err
	}
	return strings.Split(out, "\n"), nil
}

func (c *DefaultClient) RebuildNetwork(name string) error {
	cmd := fmt.Sprintf(networkRebuildCmd, name)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) RebuildAllNetworks() error {
	_, err := c.Exec(networkRebuildAllCmd)
	return err
}

func (c *DefaultClient) GetAppNetworkReport(appName string) (*AppNetworkReport, error) {
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

func (c *DefaultClient) GetNetworkReport() (NetworkReport, error) {
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

func (c *DefaultClient) SetAppNetworkProperty(appName string, property string, value string) error {
	cmd := fmt.Sprintf(networkSetPropertyCmd, appName, property, value)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) RemoveAppNetworkProperty(appName string, property string) error {
	cmd := fmt.Sprintf(networkSetPropertyCmd, appName, property, "")
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) SetGlobalNetworkProperty(property string, value string) error {
	cmd := fmt.Sprintf(networkSetPropertyCmd, "--global", property, value)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) RemoveGlobalNetworkProperty(property string) error {
	cmd := fmt.Sprintf(networkSetPropertyCmd, "--global", property, "")
	_, err := c.Exec(cmd)
	return err
}
