package repository

import "crm/models"

type ContactRepository interface {
	Add(c *models.Contact) error
	List() ([]models.Contact, error)
	Delete(id int) error
	Update(c *models.Contact) error
}
