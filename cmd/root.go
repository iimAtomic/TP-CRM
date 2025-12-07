package cmd

import (
	"fmt"
	"os"

	"crm/handlers"
	"crm/repository/json"
	"crm/service"

	"github.com/spf13/cobra"
)


var repo *json.JSONContactRepository
var svc *service.ContactService
var handler *handlers.ContactHandler

var rootCmd = &cobra.Command{
	Use:   "crm",
	Short: "CRM CLI - Gestion des contacts",
	Long:  `CRM en ligne de commande permettant d'ajouter, lister, mettre à jour et supprimer des contacts.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tapez --help pour voir les commandes disponibles.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Initialisation globale
	repo = json.NewJSONContactRepository()
	svc = service.NewContactService(repo)
	handler = handlers.NewContactHandler(svc)

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Mode verbose")
}
