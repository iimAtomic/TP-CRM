package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"

	"crm/service"
)

var reader = bufio.NewReader(os.Stdin)

type ContactHandler struct {
	service *service.ContactService
}

func NewContactHandler(s *service.ContactService) *ContactHandler {
	return &ContactHandler{service: s}
}

func (h *ContactHandler) AddContact() {
	fmt.Print("Nom : ")
	nom, _ := reader.ReadString('\n')

	fmt.Print("Email : ")
	email, _ := reader.ReadString('\n')

	err := h.service.Add(strings.TrimSpace(nom), strings.TrimSpace(email))
	if err != nil {
		fmt.Println("Erreur :", err)
		return
	}
	fmt.Println("Contact ajouté avec succès !")
}

func (h *ContactHandler) ListContacts() {
	contacts, _ := h.service.List()

	if len(contacts) == 0 {
		fmt.Println("Aucun contact.")
		return
	}

	for _, c := range contacts {
		fmt.Printf("ID: %d - %s (%s)\n", c.ID, c.Nom, c.Email)
	}
}

func (h *ContactHandler) DeleteContact() {
	fmt.Print("ID du contact à supprimer : ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	err := h.service.Delete(id)
	if err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Println("Contact supprimé.")
	}
}

func (h *ContactHandler) UpdateContact() {
	fmt.Print("ID du contact : ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Nouveau nom : ")
	nom, _ := reader.ReadString('\n')

	fmt.Print("Nouvel email : ")
	email, _ := reader.ReadString('\n')

	err := h.service.Update(id, strings.TrimSpace(nom), strings.TrimSpace(email))
	if err != nil {
		fmt.Println("Erreur :", err)
	} else {
		fmt.Println("Mise à jour réussie.")
	}
}
