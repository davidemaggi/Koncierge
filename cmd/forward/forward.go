package forward

import (
	"github.com/davidemaggi/koncierge/internal/config"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/k8s"
	"github.com/davidemaggi/koncierge/internal/ui"
	"github.com/davidemaggi/koncierge/internal/wizard"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"syscall"
)

var FwdCmd = &cobra.Command{
	Use:     "forward",
	Aliases: []string{"fwd"},
	Short:   "Just start a new forward",
	Long:    `A wizard driven port-forward nobody wants to remember commands.`,
	Run:     runCommand,
}

func init() {
	FwdCmd.AddCommand(FwdAddCmd)
	FwdCmd.AddCommand(FwdStartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forwardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forwardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runCommand(cmd *cobra.Command, args []string) {

	_ = cmd
	_ = args

	logger := container.App.Logger
	kubeService, err := k8s.ConnectToCluster(config.KubeConfigFile)

	if err != nil {
		logger.Error("Cannot Connect to cluster", err)
		return
	}

	fwd := wizard.BuildForward()

	stop, ready, err := kubeService.StartPortForward(fwd)
	if err != nil {
		logger.Error("Error starting port forward", err)
		os.Exit(1)
	}

	<-ready
	ui.PrintForwardOverview(fwd, nil)

	logger.Info("Press Ctrl+C to stop...")

	ctx, stopSig := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopSig()

	// Wait for Ctrl+C
	<-ctx.Done()

	logger.Info("Shutting down port-forward...")

	close(stop)

}
