package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/yzastyle/encode-go-rest/internal/app"
	"github.com/yzastyle/encode-go-rest/internal/logic"
)

const (
	pathParamId = "id"
)

type PersonHandler interface {
	GetAllPersons() func(echo.Context) error
	GetPersonById() func(echo.Context) error
	CreatePerson() func(echo.Context) error
	UpdatePerson() func(echo.Context) error
	DeletePerson() func(echo.Context) error
}

type personHandlerImpl struct {
	personLogic logic.PersonLogic
}

func NewPersonHandler(l logic.PersonLogic) PersonHandler {
	return &personHandlerImpl{personLogic: l}
}

func (h *personHandlerImpl) GetAllPersons() func(echo.Context) error {
	persons := h.personLogic.GetAllPersons()

	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, persons)
	}
}

func (h *personHandlerImpl) GetPersonById() func(echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param(pathParamId)
		person := h.personLogic.GetPersonById(id)
		if person != nil {
			return c.JSON(http.StatusOK, person)
		}
		return c.String(http.StatusNotFound, "Person with id="+id+" not found")
	}
}

func (h *personHandlerImpl) CreatePerson() func(echo.Context) error {
	return func(c echo.Context) error {
		u := new(app.PersonDTO)
		if err := c.Bind(u); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		person := app.FromDTO(u)
		person.Id = app.BuildId().String()

		if err := h.personLogic.CreatePerson(&person); err != nil {
			return c.String(http.StatusInternalServerError, "Failed to create person")
		}
		return c.JSON(http.StatusCreated, person)
	}
}

func (h *personHandlerImpl) UpdatePerson() func(echo.Context) error {
	return func(c echo.Context) error {
		u := new(app.PersonDTO)
		if err := c.Bind(u); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}
		person := app.FromDTO(u)
		person.Id = c.Param(pathParamId)

		if err := h.personLogic.UpdatePerson(&person); err != nil {
			return c.String(http.StatusInternalServerError, "Failed to update person")
		}
		return c.JSON(http.StatusOK, person)
	}
}

func (h *personHandlerImpl) DeletePerson() func(echo.Context) error {
	return func(c echo.Context) error {
		if err := h.personLogic.DeletePerson(c.Param(pathParamId)); err != nil {
			return c.String(http.StatusInternalServerError, "Failed to update person")
		}
		return c.NoContent(http.StatusNoContent)
	}
}
