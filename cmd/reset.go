package cmd

import (
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/db"
	"github.com/spf13/cobra"
)

var ResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Delete The Koncierge DB to start over",
	Long:  `Delete the configuration to have a fresh start`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := container.App.Logger

		err := db.Reset()
		if err != nil {

			logger.Error("Error on Koncierge Reset.", err)

		}

	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
