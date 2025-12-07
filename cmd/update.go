package cmd

import "github.com/spf13/cobra"

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Mettre à jour un contact existant",
	Run: func(cmd *cobra.Command, args []string) {
		handler.UpdateContact()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
