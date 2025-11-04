package crud

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"crm/memory"
)

var reader = bufio.NewReader(os.Stdin)

func AddContact(store memory.Storer) {
	fmt.Print("Nom : ")
	nom, _ := reader.ReadString('\n')
	nom = strings.TrimSpace(nom)

	fmt.Print("Email : ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	err := store.Add(&memory.Contact{Nom: nom, Email: email})
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	fmt.Println("Contact ajouté avec succès !")
}

func ListContact(store memory.Storer) {
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

func SupContact(store memory.Storer) {
	ListContact(store)
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

func UpdateContact(store memory.Storer) {
	ListContact(store)
	fmt.Print("\nEntrez l'ID du contact à mettre à jour : ")
	var contactID int
	fmt.Scanln(&contactID)

	fmt.Print("Nouveau Nom (laisser vide pour ne pas changer) : ")
	newNom, _ := reader.ReadString('\n')
	newNom = strings.TrimSpace(newNom)

	fmt.Print("Nouvel Email (laisser vide pour ne pas changer) : ")
	newEmail, _ := reader.ReadString('\n')
	newEmail = strings.TrimSpace(newEmail)

	err := store.Update(&memory.Contact{ID: contactID, Nom: newNom, Email: newEmail})
	if err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Printf("Contact avec l'ID %d mis à jour avec succès.\n", contactID)
	}
}
