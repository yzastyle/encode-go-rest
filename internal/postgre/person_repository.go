package postgre

import (
	"github.com/gocraft/dbr/v2"
	logg "github.com/sirupsen/logrus"
	"github.com/yzastyle/encode-go-rest/internal/app"
	"github.com/yzastyle/encode-go-rest/internal/logger"
)

const (
	Person = "person"
	All    = "*"
)

type PersonRepository interface {
	GetAllPersons(criteriaDTO app.PersonSearchCriteriaDTO) []app.Person
	GetPersonById(id string) *app.Person
	CreatePerson(person *app.Person) error
	UpdatePerson(person *app.Person) error
	DeletePerson(id string) error
}

type personRepositoryImpl struct {
	connection   *dbr.Connection
	logger       *logg.Entry
	queryBuilder *QueryBuilder
}

func NewPersonRepository(connection *dbr.Connection) PersonRepository {
	return &personRepositoryImpl{connection: connection,
		logger:       logger.NewRepositoryLogger("person"),
		queryBuilder: &QueryBuilder{}}
}

func (r *personRepositoryImpl) GetAllPersons(criteriaDTO app.PersonSearchCriteriaDTO) []app.Person {
	log := r.logger.WithField("operation", "get_all")

	var persons []app.Person
	session := r.connection.NewSession(nil)
	sel := r.queryBuilder.CreateQuery(session.Select(All).From(Person)).
		HasFirstName(criteriaDTO.FirstName).
		HasLastName(criteriaDTO.LastName).
		HasEmail(criteriaDTO.Email).
		HasPhone(criteriaDTO.Phone).
		WithLimit(criteriaDTO.Limit).
		WithOffest(criteriaDTO.Offset).
		Build()

	_, err := sel.Load(&persons)
	if err != nil {
		return persons
	}
	log.WithField("count", len(persons)).Debug("Successfully fetched all persons")
	return persons
}

func (r *personRepositoryImpl) GetPersonById(id string) *app.Person {
	log := r.logger.WithFields(logg.Fields{"operation": "get_by_id", "person_id": id})

	var person app.Person

	session := r.connection.NewSession(nil)
	err := session.Select(All).From(Person).Where("id = ?", id).LoadOne(&person)
	if err != nil {
		if err == dbr.ErrNotFound {
			return nil
		}
		return nil
	}
	log.Debug("Successfully fetched person by id")
	return &person
}

func (r *personRepositoryImpl) CreatePerson(person *app.Person) error {
	log := r.logger.WithFields(logg.Fields{"operation": "create",
		"id":         person.Id,
		"email":      person.Email,
		"phone":      person.Phone,
		"first_name": person.FirstName,
		"last_name":  person.LastName})

	session := r.connection.NewSession(nil)
	_, err := session.InsertInto(Person).
		Columns("id", "email", "phone", "first_name", "last_name").
		Record(person).
		Exec()
	if err != nil {
		return err
	}
	log.Debug("Successfully created person")
	return nil
}

func (r *personRepositoryImpl) UpdatePerson(person *app.Person) error {
	log := r.logger.WithFields(logg.Fields{"operation": "update",
		"id":         person.Id,
		"email":      person.Email,
		"phone":      person.Phone,
		"first_name": person.FirstName,
		"last_name":  person.LastName})

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
	log.Debug("Successfully updated person")
	return nil
}

func (r *personRepositoryImpl) DeletePerson(id string) error {
	log := r.logger.WithFields(logg.Fields{"operation": "delete_by_id", "person_id": id})

	session := r.connection.NewSession(nil)
	_, err := session.DeleteFrom(Person).Where("id = ?", id).Exec()
	if err != nil {
		return err
	} else {
		log.Debug("Successfully deleted person by id")
	}
	return nil
}
