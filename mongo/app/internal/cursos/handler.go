package cursos

import (
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"mongo/app/internal/core/domain"
	"mongo/app/internal/core/ports"
	"net/http"
)

type cursos struct {
	storage ports.CursoService
}

func newCurso(storage ports.CursoService) cursos {
	return cursos{storage}
}

func (p *cursos) create(c echo.Context) error {
	data := domain.Cursos{}
	err := c.Bind(&data)
	if err != nil {
		//response := newResponse(Error, "La persona no tiene una estructura correcta", nil)
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = p.storage.CreateService(&data)
	if err != nil {
		//response := newResponse(Error, "Hubo un problema al crear la persona", nil)
		return c.JSON(http.StatusInternalServerError, err)
	}

	//response := newResponse(Message, "Persona creada correctamente", nil)
	return c.JSON(http.StatusCreated, nil)
}

func (p *cursos) byName(c echo.Context) error {
	var data map[string]interface{}
	err := c.Bind(&data)
	if err != nil {
		fmt.Println(errors.New("NO SE PUDO LEER LA DATA QUE ENVIASTE"))
		return c.JSON(http.StatusInternalServerError, err)
	}

	name, ok := data["nombre"]
	if !ok {
		//fmt.Println(errors.New("NO SE PUDO OBETENER EL CAMPO NOMBRE"))
		return c.JSON(http.StatusInternalServerError, errors.New("NO SE PUDO OBETENER EL CAMPO NOMBRE"))
	}

	cursos, err := p.storage.GetByNameService(name.(string))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(200, cursos)
}
