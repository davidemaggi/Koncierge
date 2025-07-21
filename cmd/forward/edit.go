package forward

import (
	"fmt"
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/db"
	"github.com/davidemaggi/koncierge/internal/repositories/forwardRepository"
	"github.com/davidemaggi/koncierge/internal/ui"
	"github.com/davidemaggi/koncierge/internal/wizard"
	"github.com/davidemaggi/koncierge/models"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var FwdEditCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e", "update", "upd"},
	Short:   internal.FORWARD_EDIT_SHORT,
	Long:    internal.FORWARD_EDIT_DESCRIPTION,
	Run:     runEdit,
}

func init() {

}

func runEdit(cmd *cobra.Command, args []string) {
	_ = cmd
	_ = args
	ui.PrintCommandHeader(internal.FORWARD_EDIT_SHORT, internal.FORWARD_EDIT_DESCRIPTION)

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

	selectedForward, ok := wizard.SelectOne(allForwards, "Select forward to Edit", func(f models.ForwardEntity) string {
		return f.GetAsString()
	}, "")

	if !ok {
		logger.Attention("No forwards selected.")
		logger.Warn("No forwards selected.")
		os.Exit(0)

	}

	//kubeService, _ := k8s.ConnectToClusterAndContext(config.KubeConfigFile, config.KubeContext)

	//contexts := k8s.GetAllContexts(config.KubeConfigFile)
	//current := kubeService.GetCurrentContextAsString()

	//moveToCtx, _ := wizard.SelectOne(contexts, "Select the new ctx", func(f string) string {
	//	return fmt.Sprintf("%s", f)
	//}, current)

	tmpLocal, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue(fmt.Sprintf("%d", selectedForward.LocalPort)).WithDefaultText("Insert the Local Port").Show()

	if val, err := strconv.ParseInt(tmpLocal, 10, 32); err == nil {
		selectedForward.LocalPort = int32(val)

	} else {
		logger.Error("Failed to parse local port number", err)
		os.Exit(1)
	}

	result, _ := pterm.DefaultInteractiveConfirm.Show("Are you really sure you want to update this forwards?")

	if result {

		err := forwardRepo.Update(&selectedForward)
		if err != nil {
			logger.Failure("Error Updating Forward")
			logger.Error("Error Updating Forward.", err)
			os.Exit(1)
		}

	}

	logger.Success("Operation Completed!")

}
