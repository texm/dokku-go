package dokku

import "fmt"

func (s *DokkuTestSuite) TestManageAppConfig() {
	r := s.Require()

	testApp := "test-config-app"
	r.NoError(s.Client.CreateApp(testApp))

	r.Error(s.Client.SetAppConfigValue(testApp, "invalid key", "", false))

	key := "key"
	value := "value with spaces"
	r.NoError(s.Client.SetAppConfigValue(testApp, key, value, false))

	config, err := s.Client.GetAppConfig(testApp)
	r.NoError(err)
	r.Contains(config, key)
	r.Equal(config[key], value)

	evalExport, err := s.Client.ExportAppConfig(testApp, ConfigExportFormatEval)
	r.NoError(err)
	r.Equal(evalExport, fmt.Sprintf("export %s='%s'", key, value))

	shellExport, err := s.Client.ExportAppConfig(testApp, ConfigExportFormatShell)
	r.NoError(err)
	r.Equal(shellExport, fmt.Sprintf("%s='%s'", key, value))

	key2 := "key2"
	value2 := "value2"
	r.NoError(s.Client.SetAppConfigValues(testApp, map[string]string{
		key:  value,
		key2: value2,
	}, false))
	keys, err := s.Client.GetAppConfigKeys(testApp)
	r.NoError(err)
	r.ElementsMatch(keys, []string{key, key2})
}

func (s *DokkuTestSuite) TestManageGlobalConfig() {
	r := s.Require()

	key := "key"
	value := "value"
	r.NoError(s.Client.SetGlobalConfigValue(key, value, false))

	config, err := s.Client.GetGlobalConfig()
	r.NoError(err)
	r.Contains(config, key)
	r.Equal(config[key], value)
}
