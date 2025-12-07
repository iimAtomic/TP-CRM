package cmd

import "github.com/spf13/cobra"

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Ajouter un contact",
	Run: func(cmd *cobra.Command, args []string) {
		handler.AddContact()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	
}
