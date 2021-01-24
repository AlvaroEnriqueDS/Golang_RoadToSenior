package cursos

import "mongo/app/internal/core/domain"

type Cursos struct {
	ID          string `json:"_id" bson:"_id,omitempty"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Clases      []struct {
		ID     string   `json:"_id" bson:"_id,omitempty"`
		Orden  int      `json:"orden"`
		Nombre string   `json:"nombre"`
		Video  string `json:"video"`
	} `json:"clases"`
}

type CursoRepository interface {
	GetRepository(id string) (domain.Cursos, error)

	SaveRepository(curso *domain.Cursos) error
	GetByNameRepository(name string) ([]domain.Cursos, error)
}

type CursoService interface {
	GetService(id string) (domain.Cursos,error)

	CreateService(curso *domain.Cursos) error
	GetByNameService(name string) ([]domain.Cursos, error)
}