package dokku

import (
	"errors"
	"io"
	"time"
)

const (
	noAppsDokkuMessage          = "You haven't deployed any applications yet"
	nameTakenMessage            = "!     Name is already taken"
	lockCreatedMessage          = "-----> Deploy lock created"
	deployLockExistsMessage     = "Deploy lock exists"
	deployLockNotExistsMessage  = "!     Deploy lock does not exist"
	appNotExistsMessageTemplate = "!     App %s does not exist"
)

var (
	InvalidAppError        = errors.New("app does not exist")
	AppNotDeployedError    = errors.New("app is not deployed")
	NoDeployedAppsError    = errors.New("no apps have been deployed")
	NotImplementedError    = errors.New("method not implemented")
	UnexpectedMessageError = errors.New("unexpected confirmation message")
	NameTakenError         = errors.New("app name already in use")
)

type Client interface {
	Dial() error
	DialWithTimeout(timeout time.Duration) error
	Close() error

	Exec(command string) (string, error)

	appManager
	processManager
	resourceManager
	gitManager
	logsManager
	checksManager
	networksManager
}

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

type processManager interface {
	GetProcessInfo(appName string) error
	GetAppProcessReport(appName string) (*ProcessReport, error)
	GetAllProcessReport() (ProcessesReport, error)
	GetAppProcessScale(appName string) (map[string]int, error)
	SetAppProcessScale(appName string, processName string, scale int, skipDeploy bool) error
	StartApp(appName string, p *ParallelismOptions) error
	StartAllApps(p *ParallelismOptions) error
	StopApp(appName string, p *ParallelismOptions) error
	StopAllApps(p *ParallelismOptions) error
	RebuildApp(appName string, p *ParallelismOptions) error
	RebuildAllApps(p *ParallelismOptions) error
	RestartApp(appName string, p *ParallelismOptions) error
	RestartAppProcess(appName string, process string, p *ParallelismOptions) error
	RestartAllApps(p *ParallelismOptions) error
	SetAppProcessProperty(appName string, key string, value string) error
	SetGlobalProcessProperty(key string, value string) error
	SetAppProcfilePath(appName string, procPath string) error
	SetGlobalProcfilePath(procPath string) error
	SetAppRestartPolicy(appName string, policy RestartPolicy) error
	SetGlobalRestartPolicy(policy RestartPolicy) error
}

type resourceManager interface {
	GetAppResourceReport(appName string) (*ResourceReport, error)
	GetAllAppResourceReport() (ResourcesReport, error)
	SetAppDefaultResourceLimit(appName string, resource ResourceSpec, limit int) error
	ClearAppDefaultResourceLimit(appName string, resource ResourceSpec) error
	ClearAppDefaultResourceLimits(appName string) error
	SetAppProcessResourceLimit(appName string, process string, resource ResourceSpec, limit int) error
	ClearAppProcessResourceLimit(appName string, process string, resource ResourceSpec) error
	ClearAppProcessResourceLimits(appName string, process string) error
	SetAppResourceReservation(appName string, resource ResourceSpec, reserve int) error
	ClearAppResourceReservation(appName string, resource ResourceSpec) error
	ClearAppResourceReservations(appName string) error
	SetAppProcessResourceReservation(appName string, process string, resource ResourceSpec, reserve int) error
	ClearAppProcessResourceReservation(appName string, process string, resource ResourceSpec) error
	ClearAppProcessResourceReservations(appName string, process string) error
}

type gitManager interface {
	GitInitializeApp(appName string) error
	GitGetPublicKey() (string, error)
	GitSyncAppRepo(appName string, repo string, opt *GitSyncOptions) error
	GitCreateFromArchive(appName string, url string, opt *GitArchiveOptions) error
	GitCreateFromImage(appName string, image string, opt *GitImageOptions) error
	GitSetAuth(host string, username string, password string) error
	GitRemoveAuth(host string) error
	GitSetAppProperty(appName string, property string, val string) error
	GitRemoveAppProperty(appName string, property string) error
	GitAllowHost(host string) error
	GitUnlockApp(appName string, force bool) error
	GitGetAppReport(appName string) (*GitAppReport, error)
	GitGetReport() (GitReport, error)
}

type logsManager interface {
	SetEventLoggingEnabled(enabled bool) error
	GetEventLogs() (string, error)
	ListLoggedEvents() ([]string, error)

	TailAppLogs(appName string) (io.Reader, error)
	GetAppLogs(string) (string, error)
	GetNAppLogs(appName string, numLines int) (string, error)
	GetAppProcessLogs(appName, process string) (string, error)
	GetAppFailedDeployLogs(appName string) (string, error)
	GetAllFailedDeployLogs() (string, error)
}

type checksManager interface {
	GetDeployChecksReport() (ChecksReport, error)
	GetAppDeployChecksReport(appName string) (*AppChecksReport, error)
	EnableAppDeployChecks(appName string) error
	EnableAppProcessesDeployChecks(appName string, processes []string) error
	DisableAppDeployChecks(appName string) error
	DisableAppProcessesDeployChecks(appName string, processes []string) error
	SetAppDeployChecksSkipped(appName string) error
	SetAppProcessesDeployChecksSkipped(appName string, processes []string) error
}

type networksManager interface {
	CreateNetwork(name string) error
	DestroyNetwork(name string) error
	CheckNetworkExists(name string) (bool, error)
	GetNetworkInfo(name string) (interface{}, error)
	ListNetworks() ([]string, error)
	RebuildNetwork(name string) error
	RebuildAllNetworks() error
	GetNetworkReport(name string) (interface{}, error)
	SetNetworkProperty(name string, property string, value string) error
}
