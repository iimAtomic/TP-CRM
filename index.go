package main

import (
	"fmt"
	"os"
)

var contacts []Contact

var nextID int = 1

type Contact struct {
	ID    int
	Nom   string
	Email string
}

func main() {

	for {
		fmt.Println("\n****** Faites un choix ******** ")
		fmt.Println("1- Ajouter un contact")
		fmt.Println("2- Lister les contacts")
		fmt.Println("3- Supprimer un contact")
		fmt.Println("4- Mettre à jour un contact")
		fmt.Println("5- Quitter l'application")
		fmt.Println(" ********\n")

		var Choix int
		fmt.Print("Votre choix : ")
		_, err := fmt.Scanln(&Choix)
		if err != nil {
			fmt.Println("Erreur de saisie. Veuillez entrer un numéro valide.")
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
			os.Exit(0)
		default:
			fmt.Println("Choix invalide. Veuillez entrer un numéro entre 1 et 5.")
		}
	}
}

func addContact() {
	var Nom string
	var Email string

	fmt.Println("\n******** Entrez votre contact ********")

	fmt.Print("Nom : ")
	_, _ = fmt.Scanln(&Nom)

	fmt.Print("Email : ")
	_, _ = fmt.Scanln(&Email)

	newContact := Contact{
		ID:    nextID,
		Nom:   Nom,
		Email: Email,
	}

	contacts = append(contacts, newContact)

	nextID++

	fmt.Printf("\nContact ajouté avec succès")
}

func listContact() {
	if len(contacts) == 0 {
		fmt.Println("\nLa liste de contacts est vide.")
		return
	}

	fmt.Println("\nListe des Contacts")

	for _, contact := range contacts {
		fmt.Printf("ID: %d, Nom: %s, Email: %s\n", contact.ID, contact.Nom, contact.Email)
	}

	fmt.Println("\n")
}

func supContact() {
	if len(contacts) == 0 {
		fmt.Println("\nImpossible de supprimer, la liste de contacts est vide.")
		return
	}

	listContact()

	var contactID int
	fmt.Print("\nVeuillez entrer l'ID du contact à supprimer : ")
	_, err := fmt.Scanln(&contactID)
	if err != nil {
		fmt.Println("Erreur de saisie. ID invalide.")
		return
	}

	for i, contact := range contacts {
		if contact.ID == contactID {
			contacts = append(contacts[:i], contacts[i+1:]...)
			fmt.Printf("Contact avec l'ID %d supprimé avec succès.\n", contactID)
			return
		}
	}

	fmt.Printf("Aucun contact trouvé avec l'ID %d.\n", contactID)
}

func updateContact() {
	if len(contacts) == 0 {
		fmt.Println("\nImpossible de mettre à jour, la liste de contacts est vide.")
		return
	}

	listContact()

	var contactID int
	fmt.Print("\nEntrez l'ID du contact à mettre à jour : ")
	_, _ = fmt.Scanln(&contactID)

	for i := range contacts {
		if contacts[i].ID == contactID {
			var newNom string
			var newEmail string

			fmt.Printf("Mise à jour du contact ID: %d (Nom actuel: %s, Email actuel: %s)\n", contacts[i].ID, contacts[i].Nom, contacts[i].Email)

			fmt.Print("Entrez le nouveau Nom (laisser vide pour ne pas changer) : ")
			_, _ = fmt.Scanln(&newNom)
			if newNom != "" {
				contacts[i].Nom = newNom
			}

			fmt.Print("Entrez le nouvel Email (laisser vide pour ne pas changer) : ")
			_, _ = fmt.Scanln(&newEmail)
			if newEmail != "" {
				contacts[i].Email = newEmail
			}

			fmt.Printf("Contact avec l'ID %d mis à jour avec succès.\n", contactID)
			return
		}
	}

	fmt.Printf("Aucun contact trouvé avec l'ID %d.\n", contactID)
}
