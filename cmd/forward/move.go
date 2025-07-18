package forward

import (
	"fmt"
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/config"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/db"
	"github.com/davidemaggi/koncierge/internal/k8s"
	"github.com/davidemaggi/koncierge/internal/repositories/forwardRepository"
	"github.com/davidemaggi/koncierge/internal/ui"
	"github.com/davidemaggi/koncierge/internal/wizard"
	"github.com/davidemaggi/koncierge/models"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os"
)

var FwdMoveCmd = &cobra.Command{
	Use:     "move",
	Aliases: []string{"mv"},
	Short:   internal.FORWARD_MOVE_SHORT,
	Long:    internal.FORWARD_MOVE_DESCRIPTION,
	Run:     runMove,
}

func init() {

}

func runMove(cmd *cobra.Command, args []string) {
	_ = cmd
	_ = args
	ui.PrintCommandHeader(internal.FORWARD_DELETE_SHORT, internal.FORWARD_DELETE_DESCRIPTION)

	logger := container.App.Logger

	forwardRepo := forwardRepository.NewForwardRepository(db.GetDB())

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

	var toMove []models.ForwardEntity
	if deleteAll {
		toMove = allForwards
	} else {

		if len(allForwards) == 1 {
			toMove = allForwards
		} else {

			selectedForwards, ok := wizard.SelectMany(allForwards, "Select forwards to move", func(f models.ForwardEntity) string {
				return fmt.Sprintf("%s.%s.%s:%d ➡️ localhost:%d", f.ContextName, f.Namespace, f.TargetName, f.PrintPortToForward(), f.LocalPort)
			})

			if !ok || len(selectedForwards) == 0 {
				logger.Attention("No forwards selected.")
				logger.Warn("No forwards selected.")
				os.Exit(0)

			} else {
				toMove = selectedForwards
			}

		}
	}

	kubeService, _ := k8s.ConnectToClusterAndContext(config.KubeConfigFile, config.KubeContext)

	contexts := k8s.GetAllContexts(config.KubeConfigFile)
	current := kubeService.GetCurrentContextAsString()

	moveToCtx, _ := wizard.SelectOne(contexts, "Select the new ctx", func(f string) string {
		return fmt.Sprintf("%s", f)
	}, current)

	result, _ := pterm.DefaultInteractiveConfirm.Show("Are you really sure you want to move these forwards?")

	if result {
		for _, mv := range toMove {

			forwardRepo.MoveToCtx(mv.ID, moveToCtx)

		}
	}

	logger.Success("Moved forwards successfully")

}
