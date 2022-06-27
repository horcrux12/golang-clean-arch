package config

type ConfigWithoutEnv struct {
	Server struct {
		Version      string `json:"version"`
		Port         int    `json:"port"`
		Host         string `json:"host"`
		PrefixPath   string `json:"prefix_path"`
		ResourceApps string `json:"resource_apps"`
	}
	Postgresql struct {
		Address           string `json:"address"`
		DefaultSchema     string `json:"default_schema"`
		MaxOpenConnection int    `json:"max_open_connection"`
		MaxIdleConnection int    `json:"max_idle_connection"`
	}
	LogFile []string `json:"log_file"`
}

func (config ConfigWithoutEnv) GetServerVersion() string {
	return config.Server.Version
}

func (config ConfigWithoutEnv) GetServerPort() int {
	return config.Server.Port
}

func (config ConfigWithoutEnv) GetServerHost() string {
	return config.Server.Host
}

func (config ConfigWithoutEnv) GetServerPrefixPath() string {
	return config.Server.PrefixPath
}

func (config ConfigWithoutEnv) GetServerResourceApps() string {
	return config.Server.ResourceApps
}

func (config ConfigWithoutEnv) GetPostgreSQLAddress() string {
	return config.Postgresql.Address
}

func (config ConfigWithoutEnv) GetPostgreSQLDefaultSchema() string {
	return config.Postgresql.DefaultSchema
}

func (config ConfigWithoutEnv) GetPostgreSQLMaxOpenConnection() int {
	return config.Postgresql.MaxOpenConnection
}

func (config ConfigWithoutEnv) GetPostgreSQLMaxIdleConnection() int {
	return config.Postgresql.MaxIdleConnection
}

func (config ConfigWithoutEnv) GetLogFileLocation() []string {
	return config.LogFile
}
