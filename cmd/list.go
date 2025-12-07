package cmd

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lister tous les contacts",
	Run: func(cmd *cobra.Command, args []string) {
		handler.ListContacts()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
