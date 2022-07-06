package config

import (
	"fmt"
	"os"
	"strconv"
)

type ConfigWithEnv struct {
	Server struct {
		Version      string `json:"version"`
		Port         string `envconfig:"APP_REST_PORT"`
		Host         string `envconfig:"APP_REST_HOST"`
		PrefixPath   string `json:"prefix_path"`
		ResourceApps string `json:"resource_apps"`
	}
	Postgresql struct {
		Address           string `envconfig:"APP_REST_DB_DEFAULT_CONNECTION"`
		DefaultSchema     string `envconfig:"APP_REST_DB_DEFAULT_PARAM"`
		MaxOpenConnection int    `json:"max_open_connection"`
		MaxIdleConnection int    `json:"max_idle_connection"`
	}
	LogFile []string `json:"log_file"`
}

func (config ConfigWithEnv) GetServerVersion() string {
	return config.Server.Version
}

func convertStringParamToInt(key string, value string) int {
	intPort, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Invalid " + key + " : " + err.Error())
		os.Exit(3)
	}
	return intPort
}

func (config ConfigWithEnv) GetServerPort() int {
	return convertStringParamToInt("Server Port", config.Server.Port)
}

func (config ConfigWithEnv) GetServerHost() string {
	return config.Server.Host
}

func (config ConfigWithEnv) GetServerPrefixPath() string {
	return config.Server.PrefixPath
}

func (config ConfigWithEnv) GetServerResourceApps() string {
	return config.Server.ResourceApps
}

func (config ConfigWithEnv) GetPostgreSQLAddress() string {
	return config.Postgresql.Address
}

func (config ConfigWithEnv) GetPostgreSQLDefaultSchema() string {
	return config.Postgresql.DefaultSchema
}

func (config ConfigWithEnv) GetPostgreSQLMaxOpenConnection() int {
	return config.Postgresql.MaxOpenConnection
}

func (config ConfigWithEnv) GetPostgreSQLMaxIdleConnection() int {
	return config.Postgresql.MaxIdleConnection
}

func (config ConfigWithEnv) GetLogFileLocation() []string {
	return config.LogFile
}
