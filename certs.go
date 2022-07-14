package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
)

type certsManager interface {
	AddAppCert(appName string, crt string, key string) error
	UpdateAppCert(appName string, crt string, key string) error
	RemoveAppCerts(appName string) error
	ShowAppCertCRT(appName string) (string, error)
	ShowAppCertKey(appName string) (string, error)
	GenerateAppCert(appName string, domain string) error

	GetAppCertsReport(appName string) (*AppCertsReport, error)
	GetCertsReport() (CertsReport, error)
}

type AppCertsReport struct {
	Dir       string `dokku:"Ssl dir"`
	Enabled   bool   `dokku:"Ssl enabled"`
	Verified  string `dokku:"Ssl verified"`
	StartsAt  string `dokku:"Ssl starts at"`
	ExpiresAt string `dokku:"Ssl expires at"`
	Hostnames string `dokku:"Ssl hostnames"`
	Issuer    string `dokku:"Ssl issuer"`
	Subject   string `dokku:"Ssl subject"`
}
type CertsReport map[string]*AppCertsReport

const (
	certsAddCmd      = "certs:add %s %s %s"
	certsGenerateCmd = "certs:generate %s %s"
	certsRemoveCmd   = "certs:remove %s"
	certsReportCmd   = "certs:report %s"
	certsShowCrtCmd  = "certs:show %s crt"
	certsShowKeyCmd  = "certs:show %s key"
	certsUpdateCmd   = "certs:update %s %s %s"
)

func (c *DefaultClient) AddAppCert(appName string, crt string, key string) error {
	cmd := fmt.Sprintf(certsAddCmd, appName, crt, key)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) UpdateAppCert(appName string, crt string, key string) error {
	cmd := fmt.Sprintf(certsUpdateCmd, appName, crt, key)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) RemoveAppCerts(appName string) error {
	cmd := fmt.Sprintf(certsRemoveCmd, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) ShowAppCertCRT(appName string) (string, error) {
	cmd := fmt.Sprintf(certsShowCrtCmd, appName)
	return c.Exec(cmd)
}

func (c *DefaultClient) ShowAppCertKey(appName string) (string, error) {
	cmd := fmt.Sprintf(certsShowKeyCmd, appName)
	return c.Exec(cmd)
}

func (c *DefaultClient) GenerateAppCert(appName string, domain string) error {
	cmd := fmt.Sprintf(certsGenerateCmd, appName, domain)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) GetAppCertsReport(appName string) (*AppCertsReport, error) {
	cmd := fmt.Sprintf(certsReportCmd, appName)
	output, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report AppCertsReport
	if err := reports.ParseInto(output, &report); err != nil {
		return nil, err
	}

	return &report, nil
}

func (c *DefaultClient) GetCertsReport() (CertsReport, error) {
	cmd := fmt.Sprintf(certsReportCmd, "")
	output, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report CertsReport
	if err := reports.ParseIntoMap(output, &report); err != nil {
		return nil, err
	}

	return report, nil
}
