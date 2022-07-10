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
