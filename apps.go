package dokku

import (
	"errors"
	"fmt"
	"strings"

	"github.com/texm/dokku-go/internal/reports"
)

type appManager interface {
	CloneApp(currentAppName string, newAppName string, options *AppManagementOptions) error
	CreateApp(appName string) error
	DestroyApp(appName string) error
	CheckAppExists(appName string) (bool, error)
	ListApps() ([]string, error)
	LockApp(appName string) error
	IsLocked(appName string) (bool, error)
	RenameApp(currentAppName string, newAppName string, options *AppManagementOptions) error
	GetAppReport(appName string) (*AppReport, error)
	GetAllAppReport() (AppsReport, error)
	UnlockApp(appName string) error
}

type AppReport struct {
	CreatedAtTimestamp   int64  `dokku:"App created at"`
	DeploySource         string `dokku:"App deploy source"`
	DeploySourceMetadata string `dokku:"App deploy source metadata"`
	Directory            string `dokku:"App dir"`
	IsLocked             bool   `dokku:"App locked"`
}

type AppManagementOptions struct {
	SkipDeploy     bool
	IgnoreExisting bool
}

type AppsReport map[string]*AppReport

const (
	appCloneCommand     = "apps:clone %s %s %s"
	appCreateCommand    = "apps:create %s"
	appDestroyCommand   = "apps:destroy --force %s"
	appExistsCommand    = "apps:exists %s"
	appListCommand      = "apps:list"
	appLockCommand      = "apps:lock %s"
	appIsLockedCommand  = "apps:locked %s"
	appRenameCommand    = "apps:rename %s %s %s"
	appReportCommand    = "apps:report %s"
	appReportAllCommand = "apps:report"
	appUnlockCommand    = "apps:unlock %s"
)

func (o *AppManagementOptions) asFlags() string {
	if o == nil {
		return ""
	}
	var flags []string
	if o.SkipDeploy {
		flags = append(flags, "--skip-deploy")
	}
	if o.IgnoreExisting {
		flags = append(flags, "--ignore-existing")
	}
	return strings.Join(flags, " ")
}

func (c *BaseClient) CloneApp(oldName string, newName string, opts *AppManagementOptions) error {
	cmd := fmt.Sprintf(appCloneCommand, oldName, newName, opts.asFlags())
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) CreateApp(name string) error {
	cmd := fmt.Sprintf(appCreateCommand, name)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) DestroyApp(name string) error {
	cmd := fmt.Sprintf(appDestroyCommand, name)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) CheckAppExists(name string) (bool, error) {
	cmd := fmt.Sprintf(appExistsCommand, name)
	_, err := c.Exec(cmd)
	if err == InvalidAppError {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (c *BaseClient) ListApps() ([]string, error) {
	output, err := c.Exec(appListCommand)
	if err != nil {
		if errors.Is(err, NoDeployedAppsError) {
			return []string{}, nil
		}
		return nil, err
	}

	split := strings.Split(output, "\n")
	appList := split[1:]

	return appList, nil
}

func (c *BaseClient) LockApp(name string) error {
	cmd := fmt.Sprintf(appLockCommand, name)
	out, err := c.Exec(cmd)
	if err != nil {
		return err
	}
	if out != lockCreatedMsg {
		return UnexpectedMessageError
	}
	return nil
}

func (c *BaseClient) IsLocked(name string) (bool, error) {
	cmd := fmt.Sprintf(appIsLockedCommand, name)
	out, err := c.Exec(cmd)
	if out == deployLockNotExistsMsg {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return out == "Deploy lock exists", nil
}

func (c *BaseClient) RenameApp(oldName string, newName string, opts *AppManagementOptions) error {
	cmd := fmt.Sprintf(appRenameCommand, oldName, newName, opts.asFlags())
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GetAppReport(name string) (*AppReport, error) {
	cmd := fmt.Sprintf(appReportCommand, name)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	report := AppReport{}
	if err := reports.ParseInto(out, &report); err != nil {
		return nil, err
	}

	return &report, nil
}

func (c *BaseClient) GetAllAppReport() (AppsReport, error) {
	cmd := fmt.Sprintf(appReportAllCommand)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	report := AppsReport{}
	if err := reports.ParseIntoMap(out, &report); err != nil {
		return nil, err
	}

	return report, nil
}

func (c *BaseClient) UnlockApp(name string) error {
	cmd := fmt.Sprintf(appUnlockCommand, name)
	_, err := c.Exec(cmd)
	return err
}
