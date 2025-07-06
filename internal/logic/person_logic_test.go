package logic

import (
	//"fmt"
	"log"
	"math/rand"
	"strconv"
	"testing"

	"github.com/yzastyle/encode-go-rest/internal/app"
	"github.com/yzastyle/encode-go-rest/internal/config"
	"github.com/yzastyle/encode-go-rest/internal/logger"
	"github.com/yzastyle/encode-go-rest/internal/postgre"
)

func setUp() PersonLogic {
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	loggerConfig, err := logger.LoadLoggerConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}
	if err := logger.InitLogger(loggerConfig); err != nil {
		log.Fatal("Failed to init logger:", err)
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
	personLogic := NewPersonLogic(personRepository)
	return personLogic
}

func TestGetAllPersons(t *testing.T) {
	personLogic := setUp()

	persons := personLogic.GetAllPersons()

	if len(persons) == 0 {
		t.Error("Expected to get at least one person, but got zero")
	}
}

func TestGetPersonById(t *testing.T) {
	personLogic := setUp()

	person := personLogic.GetPersonById("8c2ee53f-ae6a-4db3-9597-316a2f30c619")

	if person.GetId() != "8c2ee53f-ae6a-4db3-9597-316a2f30c619" {
		t.Errorf("Expected person ID to be '8c2ee53f-ae6a-4db3-9597-316a2f30c619', but got '%s'", person.GetId())
	}
	if person.GetEmail() != "test@gmail.com" {
		t.Errorf("Expected person email to be 'test@gmail.com', but got '%s'", person.GetEmail())
	}
	if person.GetFirstName() != "Alex" {
		t.Errorf("Expected person first name to be 'Alex', but got '%s'", person.GetFirstName())
	}
	if person.GetLastName() != "Smith" {
		t.Errorf("Expected person last name to be 'Smith', but got '%s'", person.GetLastName())
	}
	if person.GetPhone() != "79005002030" {
		t.Errorf("Expected person phone to be '79005002030', but got '%s'", person.GetPhone())
	}
}

func TestCreatePerson(t *testing.T) {
	personLogic := setUp()

	person := &app.Person{Id: app.BuildId().String(),
		Email:     "Logicnewemail@gmail.com",
		FirstName: "LJane",
		LastName:  "LDoe",
		Phone:     "+1234567111"}

	err := personLogic.CreatePerson(person)
	if err != nil {
		t.Errorf("Failed to create person: %v", err)
	}
	retrievedPerson := personLogic.GetPersonById(person.Id)
	if retrievedPerson.GetId() != person.Id {
		t.Errorf("Failed to create person: %v", err)
	}
}

func TestUpdatePerson(t *testing.T) {
	personLogic := setUp()

	randPrefix := strconv.Itoa(rand.Intn(10000))

	person := personLogic.GetPersonById("9843b4b8-6c55-44a1-8d89-d581105988b9")
	person.Email = randPrefix + "LUPDATED@gmail.com"
	person.Phone = "8800811" + randPrefix
	person.FirstName = randPrefix + "LUpdatedFirstName"
	person.LastName = randPrefix + "LUpdatedLastName"

	err := personLogic.UpdatePerson(person)
	if err != nil {
		t.Errorf("Failed to update person: %v", err)
	}
	retrievedPerson := personLogic.GetPersonById(person.Id)
	if retrievedPerson.Email != person.Email {
		t.Errorf("Failed to update person email: expected %s, got %s", person.Email, retrievedPerson.Email)
	}
}

func TestDeletePerson(t *testing.T) {
	personLogic := setUp()

	person := &app.Person{Id: app.BuildId().String(),
		Email:     "ForDeleteLNewemail@gmail.com",
		FirstName: "ForDeleteL",
		LastName:  "ForDeleteL",
		Phone:     "+1234567111"}

	err := personLogic.CreatePerson(person)
	if err != nil {
		t.Errorf("Failed to create person: %v", err)
	}
	personLogic.DeletePerson(person.Id)
	retrievedPerson := personLogic.GetPersonById(person.Id)
	if retrievedPerson != nil {
		t.Errorf("Expected person to be deleted, but it still exists: %v", retrievedPerson)
	}
}
