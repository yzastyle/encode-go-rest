package postgre

import (
	"log"

	"github.com/gocraft/dbr/v2"
	"github.com/yzastyle/encode-go-rest/internal/app"
)

const (
	Person = "person"
	All    = "*"
)

type PersonRepository interface {
	GetAllPersons() []app.Person
	GetPersonById(id string) *app.Person
	CreatePerson(person *app.Person) error
	UpdatePerson(person *app.Person) error
	DeletePerson(id string)
}

type personRepositoryImpl struct {
	connection *dbr.Connection
}

func NewPersonRepository(connection *dbr.Connection) PersonRepository {
	return &personRepositoryImpl{connection: connection}
}

func (r *personRepositoryImpl) GetAllPersons() []app.Person {
	var persons []app.Person

	session := r.connection.NewSession(nil)
	session.Select(All).From(Person).Load(&persons)

	return persons
}

func (r *personRepositoryImpl) GetPersonById(id string) *app.Person {
	var person app.Person

	session := r.connection.NewSession(nil)
	session.Select(All).From(Person).Where("id = ?", id).Load(&person)
	return &person
}

func (r *personRepositoryImpl) CreatePerson(person *app.Person) error {
	session := r.connection.NewSession(nil)
	_, err := session.InsertInto(Person).
		Columns("id", "email", "phone", "first_name", "last_name").
		Record(person).
		Exec()

	if err != nil {
		return err
	}
	return nil
}

func (r *personRepositoryImpl) UpdatePerson(person *app.Person) error {
	session := r.connection.NewSession(nil)
	_, err := session.Update(Person).
		Set("email", person.GetEmail()).
		Set("phone", person.GetPhone()).
		Set("first_name", person.GetFirstName()).
		Set("last_name", person.GetLastName()).
		Where("id = ?", person.GetId()).
		Exec()
	if err != nil {
		return err
	}
	return nil
}

func (r *personRepositoryImpl) DeletePerson(id string) {
	session := r.connection.NewSession(nil)
	_, err := session.DeleteFrom(Person).Where("id = ?", id).Exec()
	if err != nil {
		log.Printf("Failed to delete person with id %s: %v", id, err)
	}
}
