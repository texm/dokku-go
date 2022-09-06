package dokku

import (
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
	"strconv"
	"strings"
)

type processManager interface {
	GetProcessInfo(appName string) error
	GetAppProcessReport(appName string) (*AppProcessReport, error)
	GetAllProcessReport() (ProcessReport, error)
	GetAppProcessScale(appName string) (map[string]int, error)
	SetAppProcessScale(appName string, processName string, scale int, skipDeploy bool) (*CommandOutputStream, error)
	StartApp(appName string, p *ParallelismOptions) (*CommandOutputStream, error)
	StartAllApps(p *ParallelismOptions) (*CommandOutputStream, error)
	StopApp(appName string, p *ParallelismOptions) (*CommandOutputStream, error)
	StopAllApps(p *ParallelismOptions) (*CommandOutputStream, error)
	RebuildApp(appName string, p *ParallelismOptions) (*CommandOutputStream, error)
	RebuildAllApps(p *ParallelismOptions) (*CommandOutputStream, error)
	RestartApp(appName string, p *ParallelismOptions) (*CommandOutputStream, error)
	RestartAppProcess(appName string, process string, p *ParallelismOptions) (*CommandOutputStream, error)
	RestartAllApps(p *ParallelismOptions) (*CommandOutputStream, error)
	SetAppProcfilePath(appName string, procPath string) error
	SetGlobalProcfilePath(procPath string) error
	SetAppRestartPolicy(appName string, policy RestartPolicy) error
	SetGlobalRestartPolicy(policy RestartPolicy) error

	/*
		SetAppProcessProperty(appName string, key string, value string) error
		SetGlobalProcessProperty(key string, value string) error
	*/
}

type AppProcessReport struct {
	Deployed             bool   `json:"deployed" dokku:"Deployed"`
	Processes            int    `json:"processes" dokku:"Processes"`
	CanScale             bool   `json:"can_scale" dokku:"Ps can scale"`
	ComputedProcfilePath string `json:"computed_procfile_path" dokku:"Ps computed procfile path"`
	GlobalProcfilePath   string `json:"global_procfile_path" dokku:"Ps global procfile path"`
	ProcfilePath         string `json:"procfile_path" dokku:"Ps procfile path"`
	RestartPolicy        string `json:"restart_policy" dokku:"Ps restart policy"`
	Restore              bool   `json:"restore" dokku:"Restore"`
	Running              bool   `json:"running" dokku:"Running"`
}

type ProcessReport map[string]*AppProcessReport

type RestartPolicy interface {
	GetPolicy() string
}

type restartPolicy struct {
	policy string
	option string
}

func (p restartPolicy) GetPolicy() string {
	if p.option != "" {
		return fmt.Sprintf("%s:%s", p.policy, p.option)
	}
	return p.policy
}

var (
	RestartPolicyAlways        = restartPolicy{policy: "always"}
	RestartPolicyNever         = restartPolicy{policy: "no"}
	RestartPolicyUnlessStopped = restartPolicy{policy: "unless-stopped"}
	RestartPolicyOnFailure     = restartPolicy{policy: "on-failure"}
)

func RetryableRestartPolicy(maxRetries int) RestartPolicy {
	return restartPolicy{
		policy: "on-failure",
		option: fmt.Sprintf("%d", maxRetries),
	}
}

type ParallelismOptions struct {
	Count            int
	UseAvailableCPUs bool
}

func getParallelism(p *ParallelismOptions) int {
	if p == nil {
		return 1
	} else if p.UseAvailableCPUs {
		return -1
	}
	return p.Count
}

const (
	psInspectCommand           = "ps:inspect %s"
	psRebuildCommand           = "ps:rebuild --parallel %d %s"
	psReportCommand            = "ps:report"
	psReportAppCommand         = "ps:report %s"
	psRestartCommand           = "ps:restart --parallel %d %s"
	psRestartAppProcessCommand = "ps:restart --parallel %d %s %s"
	psRestoreCommand           = "ps:restore %s"
	psScaleCommand             = "ps:scale %s %s"
	psSetCommand               = "ps:set %s %s %s"
	psStartCommand             = "ps:start --parallel %d %s"
	psStopCommand              = "ps:stop --parallel %d %s"
)

func (c *BaseClient) GetProcessInfo(appName string) error {
	cmd := fmt.Sprintf(psInspectCommand, appName)
	output, err := c.Exec(cmd)
	if err != nil {
		if strings.HasPrefix(output, "\"docker container inspect\" requires at least 1 argument.") {
			return AppNotDeployedError
		}
		return err
	}

	return NotImplementedError
}

func (c *BaseClient) GetAppProcessReport(appName string) (*AppProcessReport, error) {
	cmd := fmt.Sprintf(psReportAppCommand, appName)
	output, err := c.Exec(cmd)

	if err != nil {
		return nil, err
	}

	report := AppProcessReport{}
	if err := reports.ParseInto(output, &report); err != nil {
		return nil, err
	}

	return &report, nil
}

func (c *BaseClient) GetAllProcessReport() (ProcessReport, error) {
	output, err := c.Exec(psReportCommand)
	report := ProcessReport{}

	if err == NoDeployedAppsError {
		return report, nil
	} else if err != nil {
		return nil, err
	}

	if err := reports.ParseIntoMap(output, &report); err != nil {
		return nil, err
	}

	return report, nil
}

func (c *BaseClient) GetAppProcessScale(appName string) (map[string]int, error) {
	cmd := fmt.Sprintf(psScaleCommand, appName, "")
	output, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	scaleReport := map[string]int{}

	lines := strings.Split(output, "\n")
	if len(lines) < 3 {
		return nil, fmt.Errorf("invalid process scale returned")
	}

	for i := 3; i < len(lines); i++ {
		parts := strings.Split(lines[i], ":")
		processName := parts[0]
		scale := strings.Trim(parts[1], " :")
		scaleInt, err := strconv.Atoi(scale)
		if err != nil {
			return nil, fmt.Errorf("failed to convert scale (%s): %w", scale, err)
		}
		scaleReport[processName] = scaleInt
	}

	return scaleReport, nil
}

func (c *BaseClient) SetAppProcessScale(appName string, processName string, scale int, skipDeploy bool) (*CommandOutputStream, error) {
	scaleAssignment := fmt.Sprintf("%s=%d", processName, scale)
	cmd := fmt.Sprintf(psScaleCommand, appName, scaleAssignment)
	if skipDeploy {
		cmd += " --skip-deploy"
	}
	return c.ExecStreaming(cmd)
}

func (c *BaseClient) StartApp(appName string, p *ParallelismOptions) (*CommandOutputStream, error) {
	cmd := fmt.Sprintf(psStartCommand, getParallelism(p), appName)
	return c.ExecStreaming(cmd)
}

func (c *BaseClient) StartAllApps(p *ParallelismOptions) (*CommandOutputStream, error) {
	cmd := fmt.Sprintf(psStartCommand, getParallelism(p), "--all")
	return c.ExecStreaming(cmd)
}

func (c *BaseClient) StopApp(appName string, p *ParallelismOptions) (*CommandOutputStream, error) {
	cmd := fmt.Sprintf(psStopCommand, getParallelism(p), appName)
	return c.ExecStreaming(cmd)
}

func (c *BaseClient) StopAllApps(p *ParallelismOptions) (*CommandOutputStream, error) {
	cmd := fmt.Sprintf(psStopCommand, getParallelism(p), "--all")
	return c.ExecStreaming(cmd)
}

func (c *BaseClient) RebuildApp(appName string, p *ParallelismOptions) (*CommandOutputStream, error) {
	cmd := fmt.Sprintf(psRebuildCommand, getParallelism(p), appName)
	return c.ExecStreaming(cmd)
}

func (c *BaseClient) RebuildAllApps(p *ParallelismOptions) (*CommandOutputStream, error) {
	cmd := fmt.Sprintf(psRebuildCommand, getParallelism(p), "--all")
	return c.ExecStreaming(cmd)
}

func (c *BaseClient) RestartApp(appName string, p *ParallelismOptions) (*CommandOutputStream, error) {
	cmd := fmt.Sprintf(psRestartCommand, getParallelism(p), appName)
	return c.ExecStreaming(cmd)
}

func (c *BaseClient) RestartAppProcess(appName string, process string, p *ParallelismOptions) (*CommandOutputStream, error) {
	cmd := fmt.Sprintf(psRestartAppProcessCommand, getParallelism(p), appName, process)
	return c.ExecStreaming(cmd)
}

func (c *BaseClient) RestartAllApps(p *ParallelismOptions) (*CommandOutputStream, error) {
	cmd := fmt.Sprintf(psRestartCommand, getParallelism(p), "--all")
	return c.ExecStreaming(cmd)
}

func (c *BaseClient) setAppProcessProperty(appName string, key string, value string) error {
	cmd := fmt.Sprintf(psSetCommand, appName, key, value)
	_, err := c.Exec(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *BaseClient) setGlobalProcessProperty(key string, value string) error {
	cmd := fmt.Sprintf(psSetCommand, "--global", key, value)
	_, err := c.Exec(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *BaseClient) SetAppProcfilePath(appName string, procPath string) error {
	return c.setAppProcessProperty(appName, "procfile-path", procPath)
}

func (c *BaseClient) SetGlobalProcfilePath(procPath string) error {
	return c.setGlobalProcessProperty("procfile-path", procPath)
}

func (c *BaseClient) SetAppRestartPolicy(appName string, p RestartPolicy) error {
	return c.setAppProcessProperty(appName, "restart-policy", p.GetPolicy())
}

func (c *BaseClient) SetGlobalRestartPolicy(p RestartPolicy) error {
	return c.setGlobalProcessProperty("restart-policy", p.GetPolicy())
}
