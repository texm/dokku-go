package dokku

import (
	"errors"
	"io"
	"time"
)

const (
	noAppsDokkuMessage = "You haven't deployed any applications yet"
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
	// client connection methods

	Dial() error
	DialWithTimeout(timeout time.Duration) error
	Close() error

	Exec(command string) (string, error)

	// apps

	CloneApp(currentAppName string, newAppName string) error
	CreateApp(appName string) error
	DestroyApp(appName string) error
	CheckAppExists(appName string) (bool, error)
	ListApps() ([]string, error)
	LockApp(appName string) error
	IsLocked(appName string) (bool, error)
	RenameApp(currentAppName string, newAppName string) error
	GetAppReport(appName string) (*AppReport, error)
	GetAllAppReport() (AppsReport, error)
	UnlockApp(appName string) error

	// ps

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

	// resource

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

	// git

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

	// logs & events

	SetEventLoggingEnabled(enabled bool) error
	GetEventLogs() (string, error)
	ListLoggedEvents() ([]string, error)

	TailAppLogs(appName string) (io.Reader, error)
	GetAppLogs(string) (string, error)
	GetNAppLogs(appName string, numLines int) (string, error)
	GetAppProcessLogs(appName, process string) (string, error)
	GetAppFailedDeployLogs(appName string) (string, error)
	GetAllFailedDeployLogs() (string, error)

	// deploy

	SetAppDeployChecksEnabled(appName string, enabled bool) error
	DeployAppFromDockerImage(appName, image string) (string, error)
}
