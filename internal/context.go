package internal

import (
	"log"

	"github.com/gocraft/dbr/v2"
	logg "github.com/sirupsen/logrus"
	"github.com/yzastyle/encode-go-rest/internal/config"
	"github.com/yzastyle/encode-go-rest/internal/http"
	"github.com/yzastyle/encode-go-rest/internal/logger"
	"github.com/yzastyle/encode-go-rest/internal/logic"
	"github.com/yzastyle/encode-go-rest/internal/postgre"
)

type context struct {
	contextLogger    *logg.Entry
	connection       *dbr.Connection
	personRepository postgre.PersonRepository
	personLogic      logic.PersonLogic
}

func NewContext() *context {
	return &context{}
}

func InitContext() {
	loadConfig()
	contextLogger := initLogger()
	ctx := context{contextLogger: contextLogger}
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
	contextLogger := ctx.contextLogger
	serverConfig, err := http.LoadServerConfig()
	if err != nil {
		contextLogger.WithError(err).Fatal("Failed to load server config")
	}
	serverAdress := http.BuildServerAddress(serverConfig)
	http.StartServer(serverAdress, ctx.personLogic, contextLogger)
}

func initDataSource(ctx *context) {
	contextLogger := ctx.contextLogger
	dsConfig, err := postgre.LoadDataSourceConfig()
	if err != nil {
		contextLogger.WithError(err).Fatal("Failed to load data source config")
	}

	connectionUrl := postgre.BuildConnectionURL(dsConfig)
	dataSource := postgre.DataSource{}
	dataSource.SetConnectionURL(connectionUrl)
	dataSource.SetDataSourceType(dsConfig.Type)
	connection, err := dataSource.GetConnection()
	if err != nil {
		contextLogger.WithError(err).Fatal("Failed to get connection")
	}
	ctx.connection = connection
	contextLogger.Info("DataSource initialized successfully")
}

func initLogger() *logg.Entry {
	loggerConfig, err := logger.LoadLoggerConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	if err := logger.InitLogger(loggerConfig); err != nil {
		log.Fatal("Failed to init logger:", err)
	}
	contextLogger := logger.NewContextLogger("context")
	return contextLogger
}

func loadConfig() {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
}
