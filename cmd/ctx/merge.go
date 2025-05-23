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

var configMergeCmd = &cobra.Command{
	Use:     "merge",
	Aliases: []string{},
	Short:   internal.CONFIG_MERGE_SHORT,
	Long:    internal.CONFIG_MERGE_DESCRIPTION,
	Run:     runMerge,
}

var sourceConfig string

func init() {

	configMergeCmd.Flags().StringVar(&sourceConfig, "sourceConfig", "sc", "The Source KubeConfig file")

}

func runMerge(cmd *cobra.Command, args []string) {
	_ = cmd
	_ = args
	ui.PrintCommandHeader(internal.CONFIG_MERGE_SHORT, internal.CONFIG_MERGE_DESCRIPTION)

	logger := container.App.Logger

	TargetContexts := k8s.GetAllContexts(config.KubeConfigFile)
	SourceContexts := k8s.GetAllContexts(sourceConfig)

	selectedSource, ok := wizard.SelectMany(SourceContexts, "Select Contexts to merge", func(s string) string {
		return s
	})

	if !ok {
		logger.Error("Error selecting contexts.", nil)
		os.Exit(1)
	}

	if len(selectedSource) == 0 {
		logger.Warn("No Contexts selected.")
		os.Exit(0)

	}

	result, _ := pterm.DefaultInteractiveConfirm.Show("Are you really sure you want to merge these contexts?")

	if result {
		k8s.MergeContexts(selectedSource, sourceConfig, config.KubeConfigFile)
	}

	_ = logger
	_ = TargetContexts
	_ = SourceContexts

}
