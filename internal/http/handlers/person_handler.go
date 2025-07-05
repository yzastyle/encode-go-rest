package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	//"github.com/yzastyle/encode-go-rest/internal/app"
	"github.com/yzastyle/encode-go-rest/internal/logic"
)

type PersonHandler interface {
	GetAllPersons() func(echo.Context) error
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
