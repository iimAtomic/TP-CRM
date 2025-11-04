package main

import (
	"bufio"
	"fmt"
	"os"

	"crm/crud"
	"crm/memory"
)

var reader = bufio.NewReader(os.Stdin)
var storeInstance memory.Storer

func main() {
	storeInstance = memory.NewMemoryStore()

	for {
		fmt.Println("\n****** Faites un choix ******** ")
		fmt.Println("1- Ajouter un contact")
		fmt.Println("2- Lister les contacts")
		fmt.Println("3- Supprimer un contact")
		fmt.Println("4- Mettre à jour un contact")
		fmt.Println("5- Quitter l'application")
		fmt.Println(" ********\n")

		fmt.Print("Votre choix : ")
		var choix int
		_, err := fmt.Scanln(&choix)
		if err != nil {
			fmt.Println("Erreur de saisie. Veuillez entrer un numéro valide.")
			reader.ReadString('\n')
			continue
		}

		switch choix {
		case 1:
			crud.AddContact(storeInstance)
		case 2:
			crud.ListContact(storeInstance)
		case 3:
			crud.SupContact(storeInstance)
		case 4:
			crud.UpdateContact(storeInstance)
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Choix invalide. Veuillez entrer un numéro entre 1 et 5.")
		}
	}
}
