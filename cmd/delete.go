package cmd

import "github.com/spf13/cobra"

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Supprimer un contact à partir de son ID",
	Run: func(cmd *cobra.Command, args []string) {
		handler.DeleteContact()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
