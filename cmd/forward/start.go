/*
Copyright © 2025 Davide Maggi davide.maggi@proton.me

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
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
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// forwardCmd represents the forward command
var FwdStartCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"fwd start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runStart,
}

var startAll = false

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forwardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	FwdStartCmd.Flags().BoolVarP(&startAll, "all", "a", false, "Help message for toggle")

}

func runStart(cmd *cobra.Command, args []string) {
	var (
		stopChans  []chan struct{}
		readyGroup sync.WaitGroup
	)
	logger := container.App.Logger

	forwardRepo := repositories.NewGormRepository[models.ForwardEntity](db.GetDB())

	allForwards, err := forwardRepo.GetAll()

	if err != nil {
		logger.Error("Error Retrieving Forward List")
		os.Exit(0)
	}

	if len(allForwards) == 0 {
		logger.Warn("There are no forward entries in DB")
		os.Exit(0)

	}

	var toStart []models.ForwardEntity
	if startAll {
		toStart = allForwards
	} else {

		if len(allForwards) == 1 {
			toStart = allForwards
		} else {

			selectedForwards, ok := wizard.SelectMany(allForwards, "Select forwards to start", func(f models.ForwardEntity) string {
				return fmt.Sprintf("%d - %s:%d", f.ID, f.TargetName, f.LocalPort)
			})

			if !ok || len(selectedForwards) == 0 {
				log.Println("No forwards selected.")
			} else {
				toStart = selectedForwards
			}

		}
	}

	for _, tmpFwd := range toStart {

		fwd := internal.FromForwardEntity(tmpFwd)

		kubeService, err := k8s.ConnectToClusterAndContext(tmpFwd.KubeConfig.KubeconfigPath, tmpFwd.ContextName)

		fwd.PodName, _ = kubeService.GetFirstPodForService(allForwards[0].Namespace, allForwards[0].TargetName)

		stop, ready, err := kubeService.StartPortForward(fwd)
		if err != nil {
			logger.Error(fmt.Sprintf("Failed to start port forward for %s: %v", fwd.TargetName, err))
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
	logger.Info("All port-forwards are active.")
	logger.Info("Press Ctrl+C to stop...")

	// Wait for signal
	ctx, stopSig := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stopSig()
	<-ctx.Done()

	logger.Info("Shutting down all port-forwards...")

	// Close all stop channels
	for _, stop := range stopChans {
		close(stop)

	}

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
