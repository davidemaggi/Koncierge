package forward

import (
	"fmt"
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/db"
	"github.com/davidemaggi/koncierge/internal/k8s"
	"github.com/davidemaggi/koncierge/internal/repositories"
	"github.com/davidemaggi/koncierge/internal/ui"
	"github.com/davidemaggi/koncierge/internal/wizard"
	"github.com/davidemaggi/koncierge/models"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var FwdStartCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"start"},
	Short:   internal.FORWARD_START_SHORT,
	Long:    internal.FORWARD_START_DESCRIPTION,
	Run:     runStart,
}

var startAll = false
var fwdIds []int

func init() {

	FwdStartCmd.Flags().BoolVarP(&startAll, "all", "a", false, "If Selected all known forwards will be started")
	FwdStartCmd.Flags().IntSliceVarP(&fwdIds, "number", "n", []int{}, "List of Forward Ids (comma separated)")
}

func runStart(cmd *cobra.Command, args []string) {
	_ = cmd
	_ = args
	ui.PrintCommandHeader(internal.FORWARD_START_SHORT, internal.FORWARD_START_DESCRIPTION)

	var (
		stopChans  []chan struct{}
		readyGroup sync.WaitGroup
	)
	logger := container.App.Logger

	forwardRepo := repositories.NewGormRepository[models.ForwardEntity](db.GetDB())

	allForwards, err := forwardRepo.GetAll()

	if err != nil {
		logger.Failure("Error Retrieving Forward List")
		logger.Error("Error Retrieving Forward List", err)
		os.Exit(1)
	}

	if len(allForwards) == 0 {
		logger.Attention("There are no forward entries in DB")
		logger.Warn("There are no forward entries in DB")
		os.Exit(1)

	}

	var toStart []models.ForwardEntity

	if len(fwdIds) > 0 {

		toStart = selectByIndices(allForwards, fwdIds)

	} else if startAll {
		toStart = allForwards
	} else {

		if len(allForwards) == 1 {
			toStart = allForwards
		} else {

			selectedForwards, ok := wizard.SelectMany(allForwards, "Select forwards to start", func(f models.ForwardEntity) string {
				return f.GetAsString()
			})

			if !ok || len(selectedForwards) == 0 {
				logger.Attention("No forwards selected.")
				logger.Warn("No forwards selected.")
				os.Exit(0)

			} else {
				toStart = selectedForwards
			}

		}
	}
	// Create two spinners with their own writers

	for _, tmpFwd := range toStart {

		fwd := internal.FromForwardEntity(tmpFwd)

		kubeService, err := k8s.ConnectToClusterAndContext(tmpFwd.KubeConfig.KubeconfigPath, tmpFwd.ContextName)

		fwd.PodName, _ = kubeService.GetFirstPodForService(tmpFwd.Namespace, tmpFwd.TargetName)

		stop, ready, err := kubeService.StartPortForward(fwd)
		if err != nil {
			logger.Failure(fmt.Sprintf("Failed to start port forward for %s", fwd.TargetName))
			logger.Error(fmt.Sprintf("Failed to start port forward for %s", fwd.TargetName), err)
			os.Exit(1)
		}

		stopChans = append(stopChans, stop)

		// Wait for each port-forward to be ready
		readyGroup.Add(1)
		go func(f internal.ForwardDto, r chan struct{}) {
			<-r

			ui.PrintForwardOverview(f, GetValuesForAdditionalConfigs(kubeService, f))

			readyGroup.Done()
		}(fwd, ready)
	}

	// Wait for all port-forwards to be ready
	readyGroup.Wait()
	logger.Success("All port-forwards are active.")
	pterm.DefaultBasicText.Println("⌨️ Press Ctrl+C to stop...")

	// Wait for signal
	ctx, stopSig := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopSig()
	<-ctx.Done()

	pterm.DefaultBasicText.Println("🔌 Shutting down port-forward...")

	// Close all stop channels
	for _, stop := range stopChans {

		close(stop)

	}

}

func selectByIndices[T any](items []T, indices []int) []T {
	var selected []T
	for _, idx := range indices {
		if idx > 0 && idx <= len(items) {
			selected = append(selected, items[idx-1])
		}
	}
	return selected
}

func GetValuesForAdditionalConfigs(kubeService *k8s.KubeService, dto internal.ForwardDto) map[string]string {

	ret := make(map[string]string)
	for _, config := range dto.AdditionalConfigs {

		for _, str := range config.Values {

			var val string

			if config.ConfigType == internal.ConfigTypeSecret {
				val, _ = kubeService.GetSecretValue(dto.Namespace, config.Name, str)
			}

			if config.ConfigType == internal.ConfigTypeMap {
				val, _ = kubeService.GetConfigMapValue(dto.Namespace, config.Name, str)
			}

			ret[config.Name+"."+str] = val
		}
	}
	return ret
}
