package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
	"strings"
)

type domainsManager interface {
	GetAppDomainsReport(appName string) (*AppDomainsReport, error)
	GetGlobalDomainsReport() (*GlobalDomainsReport, error)
	GetDomainsReport() (DomainsReport, error)

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

type GlobalDomainsReport struct {
	Enabled bool
	Domains []string
}

type AppDomainsReport struct {
	AppEnabled    bool
	AppDomains    []string
	GlobalEnabled bool
	GlobalDomains []string
}

type DomainsReport map[string]*AppDomainsReport

const (
	domainsAddAppCmd       = "domains:add %s %s"
	domainsAddGlobalCmd    = "domains:add-global %s"
	domainsClearAppCmd     = "domains:clear %s"
	domainsClearGlobalCmd  = "domains:clear-global"
	domainsDisableAppCmd   = "domains:disable %s"
	domainsEnableAppCmd    = "domains:enable %s"
	domainsRemoveAppCmd    = "domains:remove %s %s"
	domainsRemoveGlobalCmd = "domains:remove-global %s"
	domainsReportCmd       = "domains:report %s"
	domainsSetAppCmd       = "domains:set %s %s"
	domainsSetGlobalCmd    = "domains:set-global %s"
)

type rawAppDomainsReport struct {
	AppEnabled    bool   `dokku:"Domains app enabled"`
	AppDomains    string `dokku:"Domains app vhosts"`
	GlobalEnabled bool   `dokku:"Domains global enabled"`
	GlobalDomains string `dokku:"Domains global vhosts"`
}

type rawGlobalDomainsReport struct {
	GlobalEnabled bool   `dokku:"Domains global enabled"`
	GlobalDomains string `dokku:"Domains global vhosts"`
}

func parseRawAppDomainsReport(rawReport rawAppDomainsReport) (*AppDomainsReport, error) {
	var appDomains []string
	if len(strings.TrimSpace(rawReport.AppDomains)) > 0 {
		appDomains = strings.Split(rawReport.AppDomains, " ")
	}
	var globalDomains []string
	if len(strings.TrimSpace(rawReport.GlobalDomains)) > 0 {
		globalDomains = strings.Split(rawReport.GlobalDomains, " ")
	}

	return &AppDomainsReport{
		AppEnabled:    rawReport.AppEnabled,
		AppDomains:    appDomains,
		GlobalEnabled: rawReport.GlobalEnabled,
		GlobalDomains: globalDomains,
	}, nil
}

func (c *BaseClient) GetAppDomainsReport(appName string) (*AppDomainsReport, error) {
	cmd := fmt.Sprintf(domainsReportCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var rawReport rawAppDomainsReport
	if err := reports.ParseInto(out, &rawReport); err != nil {
		return nil, err
	}

	return parseRawAppDomainsReport(rawReport)
}

func (c *BaseClient) GetGlobalDomainsReport() (*GlobalDomainsReport, error) {
	cmd := fmt.Sprintf(domainsReportCmd, "--global")
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var rawReport rawGlobalDomainsReport
	if err := reports.ParseInto(out, &rawReport); err != nil {
		return nil, err
	}

	domains := strings.Split(rawReport.GlobalDomains, " ")

	return &GlobalDomainsReport{
		Domains: domains,
		Enabled: rawReport.GlobalEnabled,
	}, nil
}

func (c *BaseClient) GetDomainsReport() (DomainsReport, error) {
	cmd := fmt.Sprintf(domainsReportCmd, "")
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	rawReportMap := map[string]rawAppDomainsReport{}
	if err := reports.ParseIntoMap(out, &rawReportMap); err != nil {
		return nil, err
	}

	reportMap := map[string]*AppDomainsReport{}
	for appName, rawReport := range rawReportMap {
		r, err := parseRawAppDomainsReport(rawReport)
		if err != nil {
			return nil, err
		}
		reportMap[appName] = r
	}

	return reportMap, nil
}

func (c *BaseClient) EnableAppDomains(appName string) error {
	cmd := fmt.Sprintf(domainsEnableAppCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) DisableAppDomains(appName string) error {
	cmd := fmt.Sprintf(domainsDisableAppCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) AddAppDomain(appName string, domain string) error {
	cmd := fmt.Sprintf(domainsAddAppCmd, appName, domain)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) RemoveAppDomain(appName string, domain string) error {
	cmd := fmt.Sprintf(domainsRemoveAppCmd, appName, domain)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) SetAppDomains(appName string, domains []string) error {
	cmd := fmt.Sprintf(domainsSetAppCmd, appName, strings.Join(domains, " "))
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) ClearAppDomains(appName string) error {
	cmd := fmt.Sprintf(domainsClearAppCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) AddGlobalDomain(domain string) error {
	cmd := fmt.Sprintf(domainsAddGlobalCmd, domain)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) RemoveGlobalDomain(domain string) error {
	cmd := fmt.Sprintf(domainsRemoveGlobalCmd, domain)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) SetGlobalDomains(domains []string) error {
	cmd := fmt.Sprintf(domainsSetGlobalCmd, strings.Join(domains, " "))
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) ClearGlobalDomains() error {
	_, err := c.Exec(domainsClearGlobalCmd)
	return err
}
