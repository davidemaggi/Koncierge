package namespace

import (
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

var NsCmd = &cobra.Command{
	Use:     "namespace",
	Aliases: []string{"ns"},
	Short:   internal.NAMESPACE_SHORT,
	Long:    internal.NAMESPACE_DESCRIPTION,
	Run:     runCommand,
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// namespaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// namespaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runCommand(cmd *cobra.Command, args []string) {

	_ = cmd
	_ = args
	ui.PrintCommandHeader(internal.NAMESPACE_SHORT, internal.NAMESPACE_DESCRIPTION)

	logger := container.App.Logger
	kubeService, err := k8s.ConnectToClusterAndContext(config.KubeConfigFile, config.KubeContext)

	if err != nil {
		logger.Error("Cannot Connect to cluster", err)
		os.Exit(1)
	}

	spaces, err := kubeService.GetAllNameSpaces()

	if err != nil {
		logger.Error("Error retrieving namespaces", err)
		os.Exit(1)
	}

	current := k8s.GetCurrentNamespaceForContext(config.KubeConfigFile, config.KubeContext)

	selNamespace, _ := wizard.SelectOne(spaces, "Select a namespace", func(f string) string {
		return f
	}, current)
	newNs := selNamespace

	logger.Info("Switching to " + pterm.Green(newNs))
	err = k8s.SetDefaultNamespaceForContext(config.KubeConfigFile, k8s.GetCurrentContextAsStringFromConfig(config.KubeConfigFile), newNs)

	if err != nil {
		os.Exit(1)
	}

}
