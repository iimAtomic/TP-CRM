package memory

import (
	"errors"
	"fmt"
	"crm/models"
	"crm/repository"
)

type MemoryContactRepository struct {
	contacts []models.Contact
	nextID   int
}

func NewMemoryContactRepository() repository.ContactRepository {
	return &MemoryContactRepository{
		contacts: []models.Contact{},
		nextID:   1,
	}
}

func (m *MemoryContactRepository) Add(c *models.Contact) error {
	if c.Nom == "" || c.Email == "" {
		return errors.New("informations manquantes")
	}

	c.ID = m.nextID
	m.nextID++
	m.contacts = append(m.contacts, *c)

	return nil
}

func (m *MemoryContactRepository) List() ([]models.Contact, error) {
	return m.contacts, nil
}

func (m *MemoryContactRepository) Delete(id int) error {
	for i, c := range m.contacts {
		if c.ID == id {
			m.contacts = append(m.contacts[:i], m.contacts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("pas de contact avec l'ID %d", id)
}

func (m *MemoryContactRepository) Update(c *models.Contact) error {
	for i := range m.contacts {
		if m.contacts[i].ID == c.ID {
			if c.Nom != "" {
				m.contacts[i].Nom = c.Nom
			}
			if c.Email != "" {
				m.contacts[i].Email = c.Email
			}
			return nil
		}
	}
	return fmt.Errorf("aucun contact trouvé avec l'ID %d", c.ID)
}
