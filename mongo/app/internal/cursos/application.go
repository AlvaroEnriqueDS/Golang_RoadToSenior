package cursos

import (
	"errors"
	"mongo/app/internal/core/domain"
	"mongo/app/internal/core/ports"
)

type service struct {
	repo ports.CursoRepository
}

func NewService(repo ports.CursoRepository) *service {
	return &service{
		repo:repo,
	}
}

func (s *service) GetService(id string) (domain.Cursos,error){
	return s.repo.GetRepository(id)
}

func (s *service) CreateService(curso *domain.Cursos) error {
	if curso == nil {
		return errors.New("el curso no puede ser nula")
	}
	curso.ID = ""
	//logica de negocio
	return s.repo.SaveRepository(curso)
}

func (s *service) GetByNameService(name string) ([]domain.Cursos, error) {
	return s.repo.GetByNameRepository(name)
}

