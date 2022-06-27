package main

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/horcrux12/clean-rest-api-template/app"
	"github.com/horcrux12/clean-rest-api-template/config"
	"github.com/horcrux12/clean-rest-api-template/constanta"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/errorModel"
	"github.com/horcrux12/clean-rest-api-template/router"
	migrate "github.com/rubenv/sql-migrate"
	"os"
	"strconv"
)

func main() {
	var arguments = constanta.DevelopmentArgument
	args := os.Args
	if len(args) > 1 {
		arguments = args[1]
	}

	config.GenerateConfiguration(arguments)
	helper.SetLogger(config.ApplicationConfiguration.GetLogFileLocation(), arguments)
	app.GenerateApplicationAttribute()

	dbMigrate()

	validatorsModel := errorModel.InitiateNewErrorsValidator()
	errorModel.AddTranslation(app.ApplicationAttribute.Validate, validatorsModel)

	logModel := applicationModel.GenerateLogModel(config.ApplicationConfiguration.GetServerVersion(), "TestApplication")
	logModel.Status = 200
	logModel.Message = "Server Start in port : " + strconv.Itoa(config.ApplicationConfiguration.GetServerPort())
	//fmt.Println(logModel)
	helper.LogInfo(logModel.ToLoggerObject())

	router.APIRouter()
}

func dbMigrate() {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_migrations"),
	}
	if app.ApplicationAttribute.DBConnection != nil {
		n, err := migrate.Exec(app.ApplicationAttribute.DBConnection, "postgres", migrations, migrate.Up)
		if err != nil {
			logModel := applicationModel.GenerateLogModel("-", config.ApplicationConfiguration.GetServerResourceApps())
			logModel.Message = err.Error()
			logModel.Status = 500
			helper.LogError(logModel.ToLoggerObject())
			os.Exit(3)
		} else {
			logModel := applicationModel.GenerateLogModel(config.ApplicationConfiguration.GetServerVersion(), config.ApplicationConfiguration.GetServerResourceApps())
			logModel.Status = 200
			logModel.Message = "Applied " + strconv.Itoa(n) + " migrations!"
			helper.LogInfo(logModel.ToLoggerObject())
		}
	} else {
		logModel := applicationModel.GenerateLogModel("-", config.ApplicationConfiguration.GetServerResourceApps())
		logModel.Message = "null database"
		logModel.Status = 500
		helper.LogError(logModel.ToLoggerObject())
		os.Exit(3)
	}
}
