package postgre

import (
	"github.com/gocraft/dbr/v2"
	"github.com/yzastyle/encode-go-rest/internal/app"
)

type PersonRepository interface {
	GetAllPersons() ([]app.Person, error)
	//GetPersonById(id string) (app.Person, error)
	//CreatePerson(person app.Person) (app.Person, error)
	//UpdatePerson(person app.Person) (app.Person, error)
	//DeletePerson(id string) error
}

type personRepositoryImpl struct {
	connection *dbr.Connection
}

func NewPersonRepository(connection *dbr.Connection) PersonRepository {
	return &personRepositoryImpl{connection: connection}
}

func (r *personRepositoryImpl) GetAllPersons() ([]app.Person, error) {
	var persons []app.Person

	session := r.connection.NewSession(nil)
	session.Select("*").From("person").Load(&persons)

	return persons, nil
}
