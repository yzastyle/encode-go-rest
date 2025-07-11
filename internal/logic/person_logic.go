package logic

import (
	"context"

	"github.com/yzastyle/encode-go-rest/internal/app"
	"github.com/yzastyle/encode-go-rest/internal/postgre"
)

type PersonLogic interface {
	GetAllPersons(ctx context.Context, criteriaDTO app.PersonSearchCriteriaDTO) []app.Person
	GetPersonById(ctx context.Context, id string) *app.Person
	CreatePerson(ctx context.Context, person *app.Person) error
	UpdatePerson(ctx context.Context, person *app.Person) error
	DeletePerson(ctx context.Context, id string) error
}

type personLogicImpl struct {
	personRepository postgre.PersonRepository
}

func NewPersonLogic(personRepository postgre.PersonRepository) PersonLogic {
	return &personLogicImpl{personRepository: personRepository}
}

func (l *personLogicImpl) GetAllPersons(ctx context.Context, criteriaDTO app.PersonSearchCriteriaDTO) []app.Person {
	if persons := l.personRepository.GetAllPersons(ctx, criteriaDTO); persons != nil {
		return persons
	}
	return make([]app.Person, 0)
}

func (l *personLogicImpl) GetPersonById(ctx context.Context, id string) *app.Person {
	return l.personRepository.GetPersonById(ctx, id)
}

func (l *personLogicImpl) CreatePerson(ctx context.Context, person *app.Person) error {
	return l.personRepository.CreatePerson(ctx, person)
}

func (l *personLogicImpl) UpdatePerson(ctx context.Context, person *app.Person) error {
	return l.personRepository.UpdatePerson(ctx, person)
}

func (l *personLogicImpl) DeletePerson(ctx context.Context, id string) error {
	return l.personRepository.DeletePerson(ctx, id)
}
