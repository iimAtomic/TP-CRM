package service

import (
	"crm/models"
	"crm/repository"
)

type ContactService struct {
	repo repository.ContactRepository
}

func NewContactService(r repository.ContactRepository) *ContactService {
	return &ContactService{repo: r}
}

func (s *ContactService) Add(nom, email string) error {
	contact := &models.Contact{
		Nom:   nom,
		Email: email,
	}
	return s.repo.Add(contact)
}

func (s *ContactService) List() ([]models.Contact, error) {
	return s.repo.List()
}

func (s *ContactService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *ContactService) Update(id int, nom, email string) error {
	contact := &models.Contact{
		ID:    id,
		Nom:   nom,
		Email: email,
	}
	return s.repo.Update(contact)
}
