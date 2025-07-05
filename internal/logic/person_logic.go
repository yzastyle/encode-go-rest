package logic

import (
	"github.com/yzastyle/encode-go-rest/internal/app"
	"github.com/yzastyle/encode-go-rest/internal/postgre"
)

type PersonLogic interface {
	GetAllPersons() []app.Person
	GetPersonById(id string) *app.Person
	CreatePerson(person *app.Person) error
	UpdatePerson(person *app.Person) error
	DeletePerson(id string)
}

type personLogicImpl struct {
	personRepository postgre.PersonRepository
}

func NewPersonLogic(personRepository postgre.PersonRepository) PersonLogic {
	return &personLogicImpl{personRepository: personRepository}
}

func (l *personLogicImpl) GetAllPersons() []app.Person {
	return l.personRepository.GetAllPersons()
}

func (l *personLogicImpl) GetPersonById(id string) *app.Person {
	return l.personRepository.GetPersonById(id)
}

func (l *personLogicImpl) CreatePerson(person *app.Person) error {
	return l.personRepository.CreatePerson(person)
}

func (l *personLogicImpl) UpdatePerson(person *app.Person) error {
	return l.personRepository.UpdatePerson(person)
}

func (l *personLogicImpl) DeletePerson(id string) {
	l.personRepository.DeletePerson(id)
}
