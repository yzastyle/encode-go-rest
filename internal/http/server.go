package http

import (
	"github.com/labstack/echo/v4"
	"github.com/yzastyle/encode-go-rest/internal/http/handlers"
	"github.com/yzastyle/encode-go-rest/internal/logic"
)

const (
	PersonUrl = "/api/v1/persons"
)

func StartServer(serverAdress string, personLogic logic.PersonLogic) {
	e := echo.New()
	// Register routes, middleware, etc.
	// e.GET("/example", exampleHandler)
	personHandler := handlers.NewPersonHandler(personLogic)

	registerRoutes(e, personHandler)

	if err := e.Start(serverAdress); err != nil {
		e.Logger.Fatal("Failed to start server:", err)
	}
}

func registerRoutes(e *echo.Echo, personHandler handlers.PersonHandler) {
	e.GET(PersonUrl, personHandler.GetAllPersons())
}
