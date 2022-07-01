package dokku

import (
	"errors"
	"fmt"
	"strings"

	"github.com/texm/dokku-go/internal/reports"
)

const (
	nameTakenMessage            = " !     Name is already taken"
	lockCreatedMessage          = "-----> Deploy lock created"
	deployLockExistsMessage     = "Deploy lock exists"
	deployLockNotExistsMessage  = "!     Deploy lock does not exist"
	appNotExistsMessageTemplate = "!     App %s does not exist"
)

const (
	appCloneCommand     = "apps:clone %s %s"
	appCreateCommand    = "apps:create %s"
	appDestroyCommand   = "apps:destroy --force %s"
	appExistsCommand    = "apps:exists %s"
	appListCommand      = "apps:list"
	appLockCommand      = "apps:lock %s"
	appIsLockedCommand  = "apps:locked %s"
	appRenameCommand    = "apps:rename --skip-deploy %s %s"
	appReportCommand    = "apps:report %s"
	appReportAllCommand = "apps:report"
	appUnlockCommand    = "apps:unlock %s"
)

func (c *DefaultClient) CloneApp(oldName, newName string) error {
	// cmd := fmt.Sprintf(cloneAppCommand, oldName, newName)

	return NotImplementedError
}

func (c *DefaultClient) CreateApp(name string) error {
	cmd := fmt.Sprintf(appCreateCommand, name)
	output, err := c.Exec(cmd)
	if output == nameTakenMessage {
		return NameTakenError
	}
	if err != nil {
		return err
	}

	return nil
}

func (c *DefaultClient) DestroyApp(name string) error {
	cmd := fmt.Sprintf(appDestroyCommand, name)
	out, err := c.Exec(cmd)
	if out == fmt.Sprintf(appNotExistsMessageTemplate, name) {
		return InvalidAppError
	} else if err != nil {
		return err
	}

	return nil
}

func (c *DefaultClient) CheckAppExists(name string) (bool, error) {
	cmd := fmt.Sprintf(appExistsCommand, name)
	_, err := c.Exec(cmd)
	if err == InvalidAppError {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (c *DefaultClient) ListApps() ([]string, error) {
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

func (c *DefaultClient) LockApp(name string) error {
	cmd := fmt.Sprintf(appLockCommand, name)
	out, err := c.Exec(cmd)
	if err != nil {
		return err
	}
	if out != lockCreatedMessage {
		return UnexpectedMessageError
	}
	return nil
}

func (c *DefaultClient) IsLocked(name string) (bool, error) {
	cmd := fmt.Sprintf(appIsLockedCommand, name)
	out, err := c.Exec(cmd)
	if out == deployLockNotExistsMessage {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return out == deployLockExistsMessage, nil
}

func (c *DefaultClient) RenameApp(oldName, newName string) error {
	return NotImplementedError
}

type AppReport struct {
	CreatedAtTimestamp   int64  `dokku:"App created at"`
	DeploySource         string `dokku:"App deploy source"`
	DeploySourceMetadata string `dokku:"App deploy source metadata"`
	Directory            string `dokku:"App dir"`
	IsLocked             bool   `dokku:"App locked"`
}

func (c *DefaultClient) GetAppReport(name string) (*AppReport, error) {
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

type AppsReport map[string]*AppReport

func (c *DefaultClient) GetAllAppReport() (AppsReport, error) {
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

func (c *DefaultClient) UnlockApp(name string) error {
	cmd := fmt.Sprintf(appUnlockCommand, name)
	_, err := c.Exec(cmd)
	return err
}
