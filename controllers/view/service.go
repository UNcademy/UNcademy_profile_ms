package view

import (
	"UNcademy_profile_ms/models"
)

type Service interface {
	ViewService(username string) (*models.User, string)
}

type service struct {
	repository Repository
}

func NewViewService(repository Repository) *service {
	return &service{repository: repository}
}

// Transforma el request en el modelo que tengo definido en mi base de datos
func (s *service) ViewService(username string) (*models.User, string) {
	user, errView := s.repository.ViewRepository(username)
	return user, errView
}
