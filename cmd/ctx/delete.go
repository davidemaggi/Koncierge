package ctx

import (
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/config"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/k8s"
	"github.com/davidemaggi/koncierge/internal/ui"
	"github.com/davidemaggi/koncierge/internal/wizard"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os"
)

var contextDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del", "rm", "remove"},
	Short:   internal.CONTEXT_DELETE_SHORT,
	Long:    internal.CONTEXT_DELETE_DESCRIPTION,
	Run:     runDelete,
}

func init() {

}

func runDelete(cmd *cobra.Command, args []string) {
	_ = cmd
	_ = args
	ui.PrintCommandHeader(internal.CONTEXT_DELETE_SHORT, internal.CONTEXT_DELETE_DESCRIPTION)

	logger := container.App.Logger

	TargetContexts := k8s.GetAllContexts(config.KubeConfigFile)
	selectedDelete, ok := wizard.SelectMany(TargetContexts, "Select Contexts to delete", func(s string) string {
		return s
	})

	if !ok {
		logger.Error("Error selecting contexts.", nil)

		os.Exit(1)
	}

	if len(selectedDelete) == 0 {
		logger.Warn("No Contexts selected.")
		logger.Attention("No Contexts selected.")

		os.Exit(0)

	}

	if len(selectedDelete) == len(TargetContexts) {
		logger.Error("You cannot delete all Contexts.", nil)
		logger.Failure("You cannot delete all Contexts.")

		os.Exit(0)

	}

	result, _ := pterm.DefaultInteractiveConfirm.Show("Are you really sure you want to delete these contexts?")

	if result {
		k8s.RemoveContexts(selectedDelete, config.KubeConfigFile)
	}

	logger.Success("Context(s) removed from config File" + pterm.Green(config.KubeConfigFile))

}
