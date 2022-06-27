package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/tkanos/gonfig"
	"os"
)

var ApplicationConfiguration Configuration

type Configuration interface {
	GetServerVersion() string
	GetServerPort() int
	GetServerHost() string
	GetServerPrefixPath() string
	GetServerResourceApps() string
	GetPostgreSQLAddress() string
	GetPostgreSQLDefaultSchema() string
	GetPostgreSQLMaxOpenConnection() int
	GetPostgreSQLMaxIdleConnection() int
	GetLogFileLocation() []string
}

func GenerateConfiguration(argument string)  {
	var errs error
	var fileName string

	enviName := os.Getenv("restProjectConfig")
	switch argument {
	case "development":
		temp := ConfigWithoutEnv{}
		fileName = "config_development.json"

		errs = gonfig.GetConf(enviName+"/"+fileName, &temp)
		if errs != nil {
			fmt.Print("Error get config development -> ", errs)
			os.Exit(2)
		}
		ApplicationConfiguration = &temp
		break
	default:
		temp := ConfigWithEnv{}
		fileName = "config_deployment.json"

		errs = gonfig.GetConf(enviName+"/"+fileName, &temp)
		if errs != nil {
			fmt.Print("Error get config deployment -> ", errs)
			os.Exit(2)
		}

		errs = envconfig.Process(enviName+"/"+fileName, &temp)
		if errs != nil {
			fmt.Print("Error get env config deployment -> ", errs)
			os.Exit(2)
		}
		ApplicationConfiguration = &temp
		break
	}
}
