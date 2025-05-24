package ctx

import (
	"fmt"
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/config"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/k8s"
	"github.com/davidemaggi/koncierge/internal/ui"
	"github.com/davidemaggi/koncierge/internal/wizard"
	"github.com/pterm/pterm"
	"os"

	"github.com/spf13/cobra"
)

var CtxCmd = &cobra.Command{
	Use:     "context",
	Aliases: []string{"ctx"},
	Short:   internal.CONTEXT_SHORT,
	Long:    internal.CONTEXT_DESCRIPTION,
	Run:     runCommand,
}

func init() {
	CtxCmd.AddCommand(contextmergeCmd)
	CtxCmd.AddCommand(contextDeleteCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// contextCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// contextCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runCommand(cmd *cobra.Command, args []string) {

	_ = cmd
	_ = args
	ui.PrintCommandHeader(internal.CONTEXT_SHORT, internal.CONTEXT_DESCRIPTION)

	logger := container.App.Logger

	kubeService, _ := k8s.ConnectToClusterAndContext(config.KubeConfigFile, config.KubeContext)

	contexts := k8s.GetAllContexts(config.KubeConfigFile)
	current := kubeService.GetCurrentContextAsString()

	newCtx, _ := wizard.SelectOne(contexts, "Select the new ctx", func(f string) string {
		return fmt.Sprintf("%s", f)
	}, current)

	logger.Info("Switching to " + pterm.Green(newCtx))
	err := k8s.SwitchContext(newCtx, config.KubeConfigFile)

	if err != nil {
		logger.Failure("Error switching to " + pterm.Red(newCtx))
		logger.Error("Error switching to "+pterm.Red(newCtx), err)
		os.Exit(1)
	}
	logger.Success("Switched to " + pterm.LightGreen(newCtx))

}
