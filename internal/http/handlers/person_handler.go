package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	logg "github.com/sirupsen/logrus"
	"github.com/yzastyle/encode-go-rest/internal/app"
	"github.com/yzastyle/encode-go-rest/internal/constants"
	"github.com/yzastyle/encode-go-rest/internal/logger"
	"github.com/yzastyle/encode-go-rest/internal/logic"
)

const (
	pathParamId     = "id"
	queryParamLimit = "limit"
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
	logger      *logg.Entry
	timeout     time.Duration
}

func NewPersonHandler(l logic.PersonLogic) PersonHandler {
	return &personHandlerImpl{personLogic: l,
		logger:  logger.NewRequestLogger(),
		timeout: time.Duration(100)}
}

func (h *personHandlerImpl) GetAllPersons() func(echo.Context) error {
	duration := h.timeout * time.Millisecond
	rootCtx := context.Background()

	return func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(rootCtx, duration)
		defer cancel()

		log := h.logger.WithFields(logg.Fields{"request_id": uuid.New().String(),
			"method": "GET",
			"path":   constants.Persons})

		log.Debug("GetAllPersons")

		var criteriaDto app.PersonSearchCriteriaDTO
		err := c.Bind(&criteriaDto)
		if err != nil {
			log.WithError(err).Error("An error occurred while executing the request.")
			return c.String(http.StatusBadRequest, "bad request")
		}
		persons := h.personLogic.GetAllPersons(ctx, criteriaDto)

		select {
		case <-ctx.Done():
			log.Info("Time occurred while executing the request.")
			return c.String(http.StatusRequestTimeout, "timeout")
		default:
			return c.JSON(http.StatusOK, persons)
		}
	}
}

func (h *personHandlerImpl) GetPersonById() func(echo.Context) error {
	return func(c echo.Context) error {
		log := h.logger.WithFields(logg.Fields{"request_id": uuid.New().String(),
			"method": "GET",
			"path":   constants.Person})

		id := c.Param(pathParamId)
		log.Debug("GetPersonById with id=" + id)
		person := h.personLogic.GetPersonById(id)
		if person != nil {
			return c.JSON(http.StatusOK, person)
		}
		log.Debug("GetPersonById: Person with id=" + id + " not found")
		return c.String(http.StatusNotFound, "Person with id="+id+" not found")
	}
}

func (h *personHandlerImpl) CreatePerson() func(echo.Context) error {
	return func(c echo.Context) error {
		log := h.logger.WithFields(logg.Fields{"request_id": uuid.New().String(),
			"method": "POST",
			"path":   constants.Persons})

		log.Debug("CreatePerson")
		u := new(app.PersonDTO)
		if err := c.Bind(u); err != nil {
			log.WithError(err).Error("An error occurred while executing the request.")
			return c.String(http.StatusBadRequest, "bad request")
		}
		person := app.FromDTO(u)
		person.Id = uuid.New().String()

		if err := h.personLogic.CreatePerson(&person); err != nil {
			log.WithError(err).Error("An error occurred while executing the request.")
			return c.String(http.StatusInternalServerError, "Failed to create person")
		}
		return c.JSON(http.StatusCreated, person)
	}
}

func (h *personHandlerImpl) UpdatePerson() func(echo.Context) error {
	return func(c echo.Context) error {
		log := h.logger.WithFields(logg.Fields{"request_id": uuid.New().String(),
			"method": "PUT",
			"path":   constants.Person})

		log.Debug("UpdatePerson")
		u := new(app.PersonDTO)
		if err := c.Bind(u); err != nil {
			log.WithError(err).Error("An error occurred while executing the request.")
			return c.String(http.StatusBadRequest, "bad request")
		}
		person := app.FromDTO(u)
		person.Id = c.Param(pathParamId)

		if err := h.personLogic.UpdatePerson(&person); err != nil {
			log.WithError(err).Error("An error occurred while executing the request.")
			return c.String(http.StatusInternalServerError, "Failed to update person")
		}
		return c.JSON(http.StatusOK, person)
	}
}

func (h *personHandlerImpl) DeletePerson() func(echo.Context) error {
	return func(c echo.Context) error {
		log := h.logger.WithFields(logg.Fields{"request_id": uuid.New().String(),
			"method": "DELETE",
			"path":   constants.Persons})

		log.Debug("DeletePerson")
		if err := h.personLogic.DeletePerson(c.Param(pathParamId)); err != nil {
			log.WithError(err).Error("An error occurred while executing the request.")
			return c.String(http.StatusInternalServerError, "Failed to update person")
		}
		return c.NoContent(http.StatusNoContent)
	}
}
