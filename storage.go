package dokku

type storageManager interface {
	EnsureStorageDirectory(directory string, chown *StorageChownOptions) error

	ListAppStorage(appName string) ([]StorageBindMount, error)
	MountAppStorage(appName string, mount StorageBindMount)
	UnmountAppStorage(appName string, mount StorageBindMount)

	GetAppStorageReport(appName string) (*AppStorageReport, error)
	GetStorageReport() (StorageReport, error)
}

type StorageChownOptions struct{}
type StorageBindMount struct{}

type AppStorageReport struct{}
type StorageReport map[string]*AppStorageReport

const (
	storageEnsureDirectoryCmd      = "storage:ensure-directory %s"
	storageEnsureDirectoryChownCmd = "storage:ensure-directory --chown %s %s"
	storageListAppCmd              = "storage:list %s"
	storageMountAppCmd             = "storage:mount %s %s"
	storageAppReportCmd            = "storage:report %s"
	storageReportCmd               = "storage:report"
	storageUnmountCmd              = "storage:unmount %s %s"
)

func (c *DefaultClient) EnsureStorageDirectory(directory string, chown *StorageChownOptions) error {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) ListAppStorage(appName string) ([]StorageBindMount, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) MountAppStorage(appName string, mount StorageBindMount) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) UnmountAppStorage(appName string, mount StorageBindMount) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetAppStorageReport(appName string) (*AppStorageReport, error) {
	//TODO implement me
	panic("implement me")
}

func (c *DefaultClient) GetStorageReport() (StorageReport, error) {
	//TODO implement me
	panic("implement me")
}
