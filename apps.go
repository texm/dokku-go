package dokku

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	noAppsMessage               = " !     You haven't deployed any applications yet"
	nameTakenMessage            = " !     Name is already taken"
	lockCreatedMessage          = "-----> Deploy lock created"
	deployLockExistsMessage     = "Deploy lock exists"
	deployLockNotExistsMessage  = "!     Deploy lock does not exist"
	appNotExistsMessageTemplate = "!     App %s does not exist"
)

const (
	appCloneCommand    = "apps:clone %s %s"
	appCreateCommand   = "apps:create %s"
	appDestroyCommand  = "apps:destroy --force %s"
	appExistsCommand   = "apps:exists %s"
	appListCommand     = "apps:list"
	appLockCommand     = "apps:lock %s"
	appIsLockedCommand = "apps:locked %s"
	appRenameCommand   = "apps:rename --skip-deploy %s %s"
	appReportCommand   = "apps:report %s"
	appUnlockCommand   = "apps:unlock %s"
)

var (
	InvalidAppError     = errors.New("App does not exist")
	NotImplementedError = errors.New("Method not implemented")
	NameTakenError      = errors.New("App name already in use")
	WeirdMessageError   = errors.New("Unexpected confirmation message")
)

func (c *Client) CloneApp(oldName, newName string) error {
	// cmd := fmt.Sprintf(cloneAppCommand, oldName, newName)

	return NotImplementedError
}

func (c *Client) CreateApp(name string) error {
	cmd := fmt.Sprintf(appCreateCommand, name)
	output, err := c.exec(cmd)
	if output == nameTakenMessage {
		return NameTakenError
	}
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DestroyApp(name string) error {
	cmd := fmt.Sprintf(appDestroyCommand, name)
	out, err := c.exec(cmd)
	if out == fmt.Sprintf(appNotExistsMessageTemplate, name) {
		return InvalidAppError
	} else if err != nil {
		return err
	}

	return nil
}

func (c *Client) CheckAppExists(name string) (bool, error) {
	cmd := fmt.Sprintf(appExistsCommand, name)
	out, err := c.exec(cmd)
	if out == fmt.Sprintf(appNotExistsMessageTemplate, name) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func (c *Client) ListApps() ([]string, error) {
	output, err := c.exec(appListCommand)
	if err != nil {
		return nil, err
	}

	var apps []string
	split := strings.Split(output, "\n")
	appList := split[1:]

	if len(appList) == 1 && appList[0] == noAppsMessage {
		return apps, nil
	}

	return appList, nil
}

func (c *Client) LockApp(name string) error {
	cmd := fmt.Sprintf(appLockCommand, name)
	out, err := c.exec(cmd)
	if err != nil {
		return err
	}
	if out != lockCreatedMessage {
		return WeirdMessageError
	}
	return nil
}

func (c *Client) IsLocked(name string) (bool, error) {
	cmd := fmt.Sprintf(appIsLockedCommand, name)
	out, err := c.exec(cmd)
	if out == deployLockNotExistsMessage {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return out == deployLockExistsMessage, nil
}

func (c *Client) RenameApp(oldName, newName string) error {
	return NotImplementedError
}

var rowRe = regexp.MustCompile(`^\s+([\s\w]*):\s*(\S+)$`)

func rowValue(row string) string {
	matches := rowRe.FindStringSubmatch(row)
	if matches == nil || len(matches) < 3 {
		return ""
	}
	return matches[2]
}

type AppInfo struct {
	CreatedAt            time.Time
	DeploySource         string
	DeploySourceMetadata string
	Directory            string
	IsLocked             bool
}

func (c *Client) AppReport(name string) (*AppInfo, error) {
	cmd := fmt.Sprintf(appReportCommand, name)
	out, err := c.exec(cmd)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out, "\n")
	createdAt := rowValue(lines[1])
	deploySource := rowValue(lines[2])
	deploySourceMetadata := rowValue(lines[3])
	directory := rowValue(lines[4])
	isLocked := rowValue(lines[5])

	stamp, err := strconv.ParseInt(createdAt, 10, 64)
	if err != nil {
		stamp = 0
	}

	info := &AppInfo{
		CreatedAt:            time.Unix(stamp, 0),
		DeploySource:         deploySource,
		DeploySourceMetadata: deploySourceMetadata,
		Directory:            directory,
		IsLocked:             isLocked == "true",
	}

	return info, nil
}

func (c *Client) UnlockApp(name string) error {
	cmd := fmt.Sprintf(appUnlockCommand, name)
	_, err := c.exec(cmd)
	return err
}
