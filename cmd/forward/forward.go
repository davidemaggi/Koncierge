package forward

import (
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/config"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/k8s"
	"github.com/davidemaggi/koncierge/internal/ui"
	"github.com/davidemaggi/koncierge/internal/wizard"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"syscall"
)

var FwdCmd = &cobra.Command{
	Use:     "forward",
	Aliases: []string{"fwd"},
	Short:   internal.FORWARD_SHORT,
	Long:    internal.FORWARD_DESCRIPTION,
	Run:     runCommand,
}

func init() {
	FwdCmd.AddCommand(FwdAddCmd)
	FwdCmd.AddCommand(FwdStartCmd)
	FwdCmd.AddCommand(FwdDeleteCmd)
	FwdCmd.AddCommand(FwdListCmd)
	FwdCmd.AddCommand(FwdMoveCmd)
	FwdCmd.AddCommand(FwdCopyCmd)
	FwdCmd.AddCommand(FwdEditCmd)

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
	ui.PrintCommandHeader(internal.FORWARD_SHORT, internal.FORWARD_DESCRIPTION)

	logger := container.App.Logger
	kubeService, err := k8s.ConnectToCluster(config.KubeConfigFile)

	if err != nil {
		logger.Failure("Cannot Connect to cluster")
		logger.Error("Cannot Connect to cluster", err)
		return
	}

	fwd := wizard.BuildForward()

	stop, ready, err := kubeService.StartPortForward(fwd)
	if err != nil {
		logger.Failure("Error starting port forward")
		logger.Error("Error starting port forward", err)
		os.Exit(1)
	}

	<-ready
	ui.PrintForwardOverview(fwd, nil)

	pterm.DefaultBasicText.Println("⌨️ Press Ctrl+C to stop...")

	ctx, stopSig := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopSig()

	// Wait for Ctrl+C
	<-ctx.Done()

	pterm.DefaultBasicText.Println("🔌 Shutting down port-forward...")

	close(stop)

}
