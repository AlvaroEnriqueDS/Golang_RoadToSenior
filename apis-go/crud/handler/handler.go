package handler

import "apis-go/crud/model"

// Storage .
type Storage interface {
	Create(person *model.Person) error
	Update(ID string, person *model.Person) error
	Delete(ID string) error
	GetByID(ID int) (model.Person, error)
	GetAll() (model.Persons, error)
}