package memory

import (
	"errors"
	"fmt"
)

type Contact struct {
	ID    int
	Nom   string
	Email string
}

type Storer interface {
	Add(c *Contact) error
	List() ([]Contact, error)
	Delete(id int) error
	Update(c *Contact) error
}

type MemoryStore struct {
	contacts []Contact
	nextID   int
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		contacts: []Contact{},
		nextID:   1,
	}
}

func (m *MemoryStore) Add(c *Contact) error {
	if c.Nom == "" || c.Email == "" {
		return errors.New("informations manquantes")
	}
	c.ID = m.nextID
	m.nextID++
	m.contacts = append(m.contacts, *c)
	return nil
}

func (m *MemoryStore) List() ([]Contact, error) {
	return m.contacts, nil
}

func (m *MemoryStore) Delete(id int) error {
	for i, c := range m.contacts {
		if c.ID == id {
			m.contacts = append(m.contacts[:i], m.contacts[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("pas de contact avec l'ID %d", id)
}

func (m *MemoryStore) Update(c *Contact) error {
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
