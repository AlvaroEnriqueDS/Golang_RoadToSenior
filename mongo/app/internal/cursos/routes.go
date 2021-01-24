package cursos

import (
	"github.com/labstack/echo"
)


func Route(e *echo.Echo) {
	repoCurso := NewCursoRepository()
	servPerson := NewService(repoCurso)
	h := newCurso(servPerson)

	curso := e.Group("/v1/cursos")
	//curso.Use(middleware.Authentication)



	curso.POST("", h.create)
	curso.POST("/porNombre", h.byName)
	//curso.GET("/:id", h.getByID)

}