package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"

	"crm/models"
)

const filePath = "contacts.json"

type JSONContactRepository struct {
	mu       sync.Mutex
	contacts []models.Contact
}

func NewJSONContactRepository() *JSONContactRepository {
	repo := &JSONContactRepository{}
	repo.load()
	return repo
}

func (r *JSONContactRepository) load() {
	data, err := os.ReadFile(filePath)
	if err != nil {
		r.contacts = []models.Contact{}
		return
	}
	json.Unmarshal(data, &r.contacts)
}

func (r *JSONContactRepository) save() {
	data, _ := json.MarshalIndent(r.contacts, "", "  ")
	os.WriteFile(filePath, data, 0644)
}

// Méthode List() conforme à l'interface
func (r *JSONContactRepository) List() ([]models.Contact, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.contacts, nil
}

// Ajout d'un contact
func (r *JSONContactRepository) Add(c *models.Contact) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.contacts) == 0 {
		c.ID = 1
	} else {
		c.ID = r.contacts[len(r.contacts)-1].ID + 1
	}

	r.contacts = append(r.contacts, *c)
	r.save()
	return nil
}

// Suppression d'un contact par ID
func (r *JSONContactRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, c := range r.contacts {
		if c.ID == id {
			r.contacts = append(r.contacts[:i], r.contacts[i+1:]...)
			r.save()
			return nil
		}
	}
	return errors.New("contact non trouvé")
}

// Mise à jour d'un contact
func (r *JSONContactRepository) Update(c *models.Contact) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.contacts {
		if r.contacts[i].ID == c.ID {
			if c.Nom != "" {
				r.contacts[i].Nom = c.Nom
			}
			if c.Email != "" {
				r.contacts[i].Email = c.Email
			}
			r.save()
			return nil
		}
	}
	return fmt.Errorf("aucun contact trouvé avec l'ID %d", c.ID)
}
