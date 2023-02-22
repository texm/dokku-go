package dokku

import (
	b64 "encoding/base64"
	"fmt"
	"github.com/texm/dokku-go/internal/reports"
	"strings"
	"unicode"
)

type configManager interface {
	GetDokkuVersion() (string, error)

	SetAppJsonProperty(appName string, property AppJsonProperty, value string) error
	GetAppJsonReport(appName string) (*AppAppJsonReport, error)
	GetAllAppJsonReport() (AppJsonReport, error)

	GetGlobalConfig() (map[string]string, error)
	GetAppConfig(appName string) (map[string]string, error)

	ClearAppConfig(appName string, restart bool) error
	ClearGlobalConfig(restart bool) error

	ExportAppConfig(appName string, format ConfigExportFormat) (string, error)
	ExportGlobalConfig(format ConfigExportFormat) (string, error)

	GetAppConfigValue(appName string, key string, quoted bool) (string, error)
	GetGlobalConfigValue(key string, quoted bool) (string, error)

	SetAppConfigValue(appName string, key string, value string, restart bool) error
	UnsetAppConfigValue(appName string, key string, restart bool) error
	SetGlobalConfigValue(key string, value string, restart bool) error
	UnsetGlobalConfigValue(key string, restart bool) error

	SetAppConfigValues(appName string, config map[string]string, restart bool) error
	UnsetAppConfigValues(appName string, keys []string, restart bool) error
	SetGlobalConfigValues(config map[string]string, restart bool) error
	UnsetGlobalConfigValues(keys []string, restart bool) error

	GetAppConfigKeys(appName string) ([]string, error)
	GetGlobalConfigKeys() ([]string, error)
}

type AppAppJsonReport struct {
	Selected         string `dokku:"App json selected"`
	GlobalSelected   string `dokku:"App json global selected"`
	ComputedSelected string `dokku:"App json computed selected"`
}
type AppJsonReport map[string]*AppAppJsonReport

type ConfigExportFormat string

const (
	ConfigExportFormatShell     = ConfigExportFormat("shell")
	ConfigExportFormatEval      = ConfigExportFormat("eval")
	ConfigExportFormatTarBundle = ConfigExportFormat("tar")
)

type AppJsonProperty string

const (
	AppJsonPropertyPath = AppJsonProperty("appjson-path")
)

const (
	versionCmd = "version"

	appJsonReportCmd      = "app-json:report %s"
	appJsonSetPropertyCmd = "app-json:set %s %s %s"

	configBundleCmd = "config:bundle %s"
	configClearCmd  = "config:clear %s %s"
	configExportCmd = "config:export %s %s"
	configGetCmd    = "config:get %s %s %s"
	configKeysCmd   = "config:keys %s"
	configSetCmd    = "config:set %s --encoded %s %s"
	configShowCmd   = "config:show %s"
	configUnsetCmd  = "config:unset %s %s %s"
)

func encodeKeyValPair(key, val string) (string, error) {
	for _, r := range key {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' {
			return "", fmt.Errorf("invalid key '%s', contains %c", key, r)
		}
	}
	encodedVal := b64.StdEncoding.EncodeToString([]byte(val))
	pair := fmt.Sprintf("%s='%s'", key, encodedVal)
	return pair, nil
}

func getOptionalFlag(flag string, enabled bool) string {
	if !enabled {
		return ""
	}
	return flag
}

func (c *BaseClient) GetDokkuVersion() (string, error) {
	out, err := c.Exec(versionCmd)
	if err != nil {
		return "", err
	}
	return strings.TrimPrefix(out, "dokku version "), nil
}

func (c *BaseClient) SetAppJsonProperty(appName string, property AppJsonProperty, value string) error {
	cmd := fmt.Sprintf(appJsonSetPropertyCmd, appName, property, value)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) GetAppJsonReport(appName string) (*AppAppJsonReport, error) {
	cmd := fmt.Sprintf(appJsonReportCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report *AppAppJsonReport
	if err := reports.ParseInto(out, &report); err != nil {
		return nil, err
	}

	return report, nil
}

func (c *BaseClient) GetAllAppJsonReport() (AppJsonReport, error) {
	cmd := fmt.Sprintf(appJsonReportCmd, "")
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	var report AppJsonReport
	if err := reports.ParseIntoMap(out, &report); err != nil {
		return nil, err
	}

	return report, nil
}

func (c *BaseClient) GetGlobalConfig() (map[string]string, error) {
	return c.GetAppConfig("--global")
}

func (c *BaseClient) GetAppConfig(appName string) (map[string]string, error) {
	cmd := fmt.Sprintf(configShowCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(out, "\n")
	config := map[string]string{}
	for i := 1; i < len(lines); i++ {
		split := strings.SplitN(lines[i], ":", 2)
		if len(split) < 2 {
			continue
		}
		key := strings.TrimSpace(split[0])
		config[key] = strings.TrimSpace(split[1])
	}
	return config, nil
}

func (c *BaseClient) ClearAppConfig(appName string, restart bool) error {
	restartFlag := getOptionalFlag("--no-restart", !restart)
	cmd := fmt.Sprintf(configClearCmd, restartFlag, appName)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) ClearGlobalConfig(restart bool) error {
	return c.ClearAppConfig("--global", restart)
}

func (c *BaseClient) ExportAppConfig(appName string, format ConfigExportFormat) (string, error) {
	var cmd string
	switch format {
	case ConfigExportFormatEval:
		cmd = fmt.Sprintf(configExportCmd, "", appName)
	case ConfigExportFormatShell:
		cmd = fmt.Sprintf(configExportCmd, "--format shell", appName)
	case ConfigExportFormatTarBundle:
		cmd = fmt.Sprintf(configBundleCmd, appName)
	default:
		return "", fmt.Errorf("unknown export format '%s'", format)
	}
	return c.Exec(cmd)
}

func (c *BaseClient) ExportGlobalConfig(format ConfigExportFormat) (string, error) {
	return c.ExportAppConfig("--global", format)
}

func (c *BaseClient) GetAppConfigValue(appName string, key string, quoted bool) (string, error) {
	quoteFlag := getOptionalFlag("--quoted", quoted)
	cmd := fmt.Sprintf(configGetCmd, quoteFlag, appName, key)
	return c.Exec(cmd)
}

func (c *BaseClient) GetGlobalConfigValue(key string, quoted bool) (string, error) {
	return c.GetAppConfigValue("--global", key, quoted)
}

func (c *BaseClient) SetAppConfigValue(appName string, key string, value string, restart bool) error {
	restartFlag := getOptionalFlag("--no-restart", !restart)
	pair, err := encodeKeyValPair(key, value)
	if err != nil {
		return err
	}
	cmd := fmt.Sprintf(configSetCmd, restartFlag, appName, pair)
	_, err = c.Exec(cmd)
	return err
}

func (c *BaseClient) UnsetAppConfigValue(appName string, key string, restart bool) error {
	restartFlag := getOptionalFlag("--no-restart", !restart)
	cmd := fmt.Sprintf(configUnsetCmd, restartFlag, appName, key)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) SetGlobalConfigValue(key string, value string, restart bool) error {
	return c.SetAppConfigValue("--global", key, value, restart)
}

func (c *BaseClient) UnsetGlobalConfigValue(key string, restart bool) error {
	return c.UnsetAppConfigValue("--global", key, restart)
}

func (c *BaseClient) SetAppConfigValues(appName string, config map[string]string, restart bool) error {
	var pairs []string
	for k, v := range config {
		pair, err := encodeKeyValPair(k, v)
		if err != nil {
			return err
		}
		pairs = append(pairs, pair)
	}
	restartFlag := getOptionalFlag("--no-restart", !restart)
	strPairs := strings.Join(pairs, " ")
	cmd := fmt.Sprintf(configSetCmd, restartFlag, appName, strPairs)
	_, err := c.Exec(cmd)
	return err
}

func (c *BaseClient) UnsetAppConfigValues(appName string, keys []string, restart bool) error {
	return c.UnsetAppConfigValue(appName, strings.Join(keys, " "), restart)
}

func (c *BaseClient) SetGlobalConfigValues(config map[string]string, restart bool) error {
	return c.SetAppConfigValues("--global", config, restart)
}

func (c *BaseClient) UnsetGlobalConfigValues(keys []string, restart bool) error {
	return c.UnsetAppConfigValues("--global", keys, restart)
}

func (c *BaseClient) GetAppConfigKeys(appName string) ([]string, error) {
	cmd := fmt.Sprintf(configKeysCmd, appName)
	out, err := c.Exec(cmd)
	if err != nil {
		return nil, err
	}
	return strings.Split(out, "\n"), nil
}

func (c *BaseClient) GetGlobalConfigKeys() ([]string, error) {
	return c.GetAppConfigKeys("--global")
}
