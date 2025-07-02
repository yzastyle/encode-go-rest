package main

import (
	"log"

	"github.com/gocraft/dbr/v2"
	"github.com/yzastyle/encode-go-rest/internal"
	"github.com/yzastyle/encode-go-rest/internal/postgre"
)

type context struct {
	connection           *dbr.Connection
	personRepositoryImpl postgre.PersonRepository
}

func (c *context) setConnection(connection *dbr.Connection) {
	c.connection = connection
}

func main() {
	ctx := &context{}
	loadConfig()
	initDataSource(ctx)
	initRepositories(ctx)
	test(ctx)
}

func test(ctx *context) {

}

func loadConfig() {
	_, err := internal.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
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
	ctx.setConnection(connection)
}
func initRepositories(ctx *context) {
	ctx.personRepositoryImpl = postgre.NewPersonRepository(ctx.connection)

}
