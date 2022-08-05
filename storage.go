package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
	"strings"
)

type storageManager interface {
	EnsureStorageDirectory(directory string, chown StorageChownOption) error

	ListAppStorage(appName string) ([]StorageBindMount, error)
	MountAppStorage(appName string, mount StorageBindMount) error
	UnmountAppStorage(appName string, mount StorageBindMount) error

	GetAppStorageReport(appName string) (*AppStorageReport, error)
	GetStorageReport() (StorageReport, error)
}

type StorageChownOption string

const (
	StorageChownOptionHerokuish = StorageChownOption("herokuish")
	StorageChownOptionHeroku    = StorageChownOption("heroku")
	StorageChownOptionPacketo   = StorageChownOption("packeto")
	StorageChownOptionNone      = StorageChownOption("false")
)

type StorageBindMount struct {
	HostDir      string
	ContainerDir string
}

func (m *StorageBindMount) String() string {
	return fmt.Sprintf("%s:%s", m.HostDir, m.ContainerDir)
}

type rawAppStorageReport struct {
	BuildMounts  string `dokku:"Storage build mounts"`
	DeployMounts string `dokku:"Storage deploy mounts"`
	RunMounts    string `dokku:"Storage run mounts"`
}
type rawStorageReport map[string]*rawAppStorageReport

func parseMount(mounts string) []StorageBindMount {
	var mountsList []StorageBindMount
	for _, mount := range strings.Split(mounts, "-v") {
		parts := strings.SplitN(strings.TrimSpace(mount), ":", 2)
		if len(parts) != 2 {
			continue
		}
		mountsList = append(mountsList, StorageBindMount{
			HostDir:      parts[0],
			ContainerDir: parts[1],
		})
	}
	return mountsList
}

func (rr *rawAppStorageReport) Parse() *AppStorageReport {
	return &AppStorageReport{
		BuildMounts:  parseMount(rr.BuildMounts),
		DeployMounts: parseMount(rr.DeployMounts),
		RunMounts:    parseMount(rr.RunMounts),
	}
}

type AppStorageReport struct {
	BuildMounts  []StorageBindMount
	DeployMounts []StorageBindMount
	RunMounts    []StorageBindMount
}
type StorageReport map[string]*AppStorageReport

const (
	storageEnsureDirectoryCmd = "storage:ensure-directory --chown %s %s"
	storageListAppCmd         = "storage:list %s"
	storageMountAppCmd        = "storage:mount %s %s"
	storageReportCmd          = "storage:report %s"
	storageUnmountCmd         = "storage:unmount %s %s"
)

func (c *DefaultClient) EnsureStorageDirectory(directory string, chown StorageChownOption) error {
	cmd := fmt.Sprintf(storageEnsureDirectoryCmd, chown, directory)
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) ListAppStorage(appName string) ([]StorageBindMount, error) {
	cmd := fmt.Sprintf(storageListAppCmd, appName)
	out, err := c.Exec(cmd)

	var mounts []StorageBindMount
	for i, line := range strings.Split(out, "\n") {
		if i == 0 {
			continue
		}
		cols := strings.Split(line, ":")
		if len(cols) != 2 {
			return nil, fmt.Errorf("error parsing storage list line '%s'", line)
		}
		mounts = append(mounts, StorageBindMount{
			HostDir:      strings.TrimSpace(cols[0]),
			ContainerDir: strings.TrimSpace(cols[1]),
		})
	}

	return mounts, err
}

func (c *DefaultClient) MountAppStorage(appName string, mount StorageBindMount) error {
	cmd := fmt.Sprintf(storageMountAppCmd, appName, mount.String())
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) UnmountAppStorage(appName string, mount StorageBindMount) error {
	cmd := fmt.Sprintf(storageUnmountCmd, appName, mount.String())
	_, err := c.Exec(cmd)
	return err
}

func (c *DefaultClient) GetAppStorageReport(appName string) (*AppStorageReport, error) {
	cmd := fmt.Sprintf(storageReportCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var rawReport *rawAppStorageReport
	if err := reports.ParseInto(out, &rawReport); err != nil {
		return nil, err
	}

	return rawReport.Parse(), nil
}

func (c *DefaultClient) GetStorageReport() (StorageReport, error) {
	cmd := fmt.Sprintf(storageReportCmd, "")
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var rawReport rawStorageReport
	if err := reports.ParseIntoMap(out, &rawReport); err != nil {
		return nil, err
	}

	report := StorageReport{}
	for i, raw := range rawReport {
		report[i] = raw.Parse()
	}

	return report, nil
}
