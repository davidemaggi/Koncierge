/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var text = "Koncierge"

		// Convert the text into a format suitable for PTerm
		var letters = putils.LettersFromString(text)

		// Render the text using PTerm's default big text style
		err := pterm.DefaultBigText.WithLetters(letters).Render()
		if err != nil {
			return
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
