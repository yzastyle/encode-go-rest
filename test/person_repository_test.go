package postgre

import (
	"log"
	"testing"

	"github.com/yzastyle/encode-go-rest/internal"
	"github.com/yzastyle/encode-go-rest/internal/app"
	"github.com/yzastyle/encode-go-rest/internal/postgre"
)

func setUp() postgre.PersonRepository {
	_, err := internal.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
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
	personRepository := postgre.NewPersonRepository(connection)
	return personRepository
}

func TestPersonRepository_GetAllPersons(t *testing.T) {
	personRepository := setUp()
	var persons []app.Person
	var err error
	persons, err = personRepository.GetAllPersons()
	if err != nil {
		t.Errorf("Failed to get all persons: %s", err)
	}
	if len(persons) == 0 {
		t.Error("Expected to get at least one person, but got zero")
	}
}
