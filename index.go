package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
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

var store Storer
var reader = bufio.NewReader(os.Stdin)

func main() {
	store = NewMemoryStore()

	for {
		fmt.Println("\n****** Faites un choix ******** ")
		fmt.Println("1- Ajouter un contact")
		fmt.Println("2- Lister les contacts")
		fmt.Println("3- Supprimer un contact")
		fmt.Println("4- Mettre à jour un contact")
		fmt.Println("5- Quitter l'application")
		fmt.Println(" ********\n")

		fmt.Print("Votre choix : ")
		var Choix int
		_, err := fmt.Scanln(&Choix)
		if err != nil {
			fmt.Println("Erreur de saisie. Veuillez entrer un numéro valide.")
			reader.ReadString('\n') // vider le buffer
			continue
		}

		switch Choix {
		case 1:
			addContact()
		case 2:
			listContact()
		case 3:
			supContact()
		case 4:
			updateContact()
		case 5:
			fmt.Println("Au revoir !")
			os.Exit(0)
		default:
			fmt.Println("Choix invalide. Veuillez entrer un numéro entre 1 et 5.")
		}
	}
}

func addContact() {
	fmt.Println("\n******** Entrez votre contact ********")
	fmt.Print("Nom : ")
	Nom, _ := reader.ReadString('\n')
	Nom = strings.TrimSpace(Nom)

	fmt.Print("Email : ")
	Email, _ := reader.ReadString('\n')
	Email = strings.TrimSpace(Email)

	err := store.Add(&Contact{Nom: Nom, Email: Email})
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	fmt.Println("Contact ajouté avec succès !")
}

func listContact() {
	contacts, _ := store.List()
	if len(contacts) == 0 {
		fmt.Println("\nLa liste de contacts est vide.")
		return
	}

	fmt.Println("\nListe des Contacts :")
	for _, c := range contacts {
		fmt.Printf("ID: %d, Nom: %s, Email: %s\n", c.ID, c.Nom, c.Email)
	}
}

func supContact() {
	listContact()
	fmt.Print("\nVeuillez entrer l'ID du contact à supprimer : ")
	var contactID int
	fmt.Scanln(&contactID)

	err := store.Delete(contactID)
	if err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Printf("Contact avec l'ID %d supprimé avec succès.\n", contactID)
	}
}

func updateContact() {
	listContact()
	fmt.Print("\nEntrez l'ID du contact à mettre à jour : ")
	var contactID int
	fmt.Scanln(&contactID)

	fmt.Print("Nouveau Nom (laisser vide pour ne pas changer) : ")
	newNom, _ := reader.ReadString('\n')
	newNom = strings.TrimSpace(newNom)

	fmt.Print("Nouvel Email (laisser vide pour ne pas changer) : ")
	newEmail, _ := reader.ReadString('\n')
	newEmail = strings.TrimSpace(newEmail)

	err := store.Update(&Contact{ID: contactID, Nom: newNom, Email: newEmail})
	if err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Printf("Contact avec l'ID %d mis à jour avec succès.\n", contactID)
	}
}
