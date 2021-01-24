package server

import (
	"github.com/labstack/echo"
	"mongo/app/internal/cursos"
)

func Start(){

	e := echo.New()

	cursos.Route(e)

	e.Start(":8080")
}
