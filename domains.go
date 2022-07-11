package dokku

type domainsManager interface {
	GetAppDomainsReport(appName string) (AppDomainsReport, error)
	GetGlobalDomainsReport(appName string) (AppDomainsReport, error)

	EnableAppDomains(appName string) error
	DisableAppDomains(appName string) error

	AddAppDomain(appName string, domain string) error
	RemoveAppDomain(appName string, domain string) error
	SetAppDomains(appName string, domains []string) error
	ClearAppDomains(appName string) error

	AddGlobalDomain(domain string) error
	RemoveGlobalDomain(domain string) error
	SetGlobalDomains(domains []string) error
	ClearGlobalDomains() error
}

type AppDomainsReport struct {
}

const (
	domainsAddAppCmd       = "domains:add <app> <domain> [<domain> ...]"
	domainsAddGlobalCmd    = "domains:add-global <domain> [<domain> ...]"
	domainsClearAppCmd     = "domains:clear <app>"
	domainsClearGlobalCmd  = "domains:clear-global"
	domainsDisableAppCmd   = "domains:disable <app>"
	domainsEnableAppCmd    = "domains:enable <app>"
	domainsRemoveAppCmd    = "domains:remove <app> <domain> [<domain> ...]"
	domainsRemoveGlobalCmd = "domains:remove-global <domain> [<domain> ...]"
	domainsReportCmd       = "domains:report [<app>|--global] [<flag>]"
	domainsSetAppCmd       = "domains:set <app> <domain> [<domain> ...]"
	domainsSetGlobalCmd    = "domains:set-global <domain> [<domain> ...]"
)

func (c *DefaultClient) GetAppDomainsReport(appName string) (AppDomainsReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetGlobalDomainsReport(appName string) (AppDomainsReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) EnableAppDomains(appName string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) DisableAppDomains(appName string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) AddAppDomain(appName string, domain string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) RemoveAppDomain(appName string, domain string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetAppDomains(appName string, domains []string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ClearAppDomains(appName string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) AddGlobalDomain(domain string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) RemoveGlobalDomain(domain string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) SetGlobalDomains(domains []string) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ClearGlobalDomains() error {
	//TODO implement me
	panic("implement me")
}
