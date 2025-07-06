package http

import (
	"github.com/labstack/echo/v4"
	logg "github.com/sirupsen/logrus"
	"github.com/yzastyle/encode-go-rest/internal/constants"
	"github.com/yzastyle/encode-go-rest/internal/http/handlers"
	"github.com/yzastyle/encode-go-rest/internal/logic"
)

func StartServer(serverAdress string, personLogic logic.PersonLogic, contextLogger *logg.Entry) {
	e := echo.New()
	// Register routes, middleware, etc.
	personHandler := handlers.NewPersonHandler(personLogic)

	registerRoutes(e, personHandler)

	if err := e.Start(serverAdress); err != nil {
		contextLogger.WithError(err).Fatal("Failed to start server")
	}
}

func registerRoutes(e *echo.Echo, personHandler handlers.PersonHandler) {
	e.GET(constants.Persons, personHandler.GetAllPersons())
	e.GET(constants.Person, personHandler.GetPersonById())
	e.POST(constants.Persons, personHandler.CreatePerson())
	e.PUT(constants.Person, personHandler.UpdatePerson())
	e.DELETE(constants.Person, personHandler.DeletePerson())
}
