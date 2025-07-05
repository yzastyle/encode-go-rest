package internal

import (
	"log"

	"github.com/gocraft/dbr/v2"
	"github.com/yzastyle/encode-go-rest/internal/config"
	"github.com/yzastyle/encode-go-rest/internal/http"
	"github.com/yzastyle/encode-go-rest/internal/logger"
	"github.com/yzastyle/encode-go-rest/internal/logic"
	"github.com/yzastyle/encode-go-rest/internal/postgre"
)

type context struct {
	connection       *dbr.Connection
	personRepository postgre.PersonRepository
	personLogic      logic.PersonLogic
}

func NewContext() *context {
	return &context{}
}

func InitContext() {
	loggerConfig := loadConfig()
	initLogger(loggerConfig)
	ctx := context{}
	initDataSource(&ctx)
	initRepositories(&ctx)
	initLogic(&ctx)
	initServer(&ctx)
}

func initLogic(ctx *context) {
	ctx.personLogic = logic.NewPersonLogic(ctx.personRepository)
}

func initRepositories(ctx *context) {
	ctx.personRepository = postgre.NewPersonRepository(ctx.connection)

}

func initServer(ctx *context) {
	serverConfig, err := http.LoadServerConfig()
	if err != nil {
		log.Fatal("Failed to load server config:", err)
	}
	serverAdress := http.BuildServerAddress(serverConfig)
	http.StartServer(serverAdress, ctx.personLogic)
}

func initDataSource(ctx *context) {
	dsConfig, err := postgre.LoadDataSourceConfig()
	if err != nil {
		log.Fatal("Failed to load data source config:", err)
	}

	connectionUrl := postgre.BuildConnectionURL(dsConfig)
	dataSource := postgre.DataSource{}
	dataSource.SetConnectionURL(connectionUrl)
	dataSource.SetDataSourceType(dsConfig.Type)
	connection, err := dataSource.GetConnection()
	if err != nil {
		log.Fatal("Failed to get connection:", err)
	}
	ctx.connection = connection
}

func initLogger(config *config.LoggerConfig) {
	logger.InitLogger(config)
}

func loadConfig() *config.LoggerConfig {
	_, loggerConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	return loggerConfig
}
