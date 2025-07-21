package forward

import (
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/db"
	"github.com/davidemaggi/koncierge/internal/repositories"
	"github.com/davidemaggi/koncierge/internal/ui"
	"github.com/davidemaggi/koncierge/internal/wizard"
	"github.com/davidemaggi/koncierge/models"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os"
)

var FwdDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm", "del", "remove"},
	Short:   internal.FORWARD_DELETE_SHORT,
	Long:    internal.FORWARD_DELETE_DESCRIPTION,
	Run:     runDelete,
}

var deleteAll = false

func init() {

	FwdDeleteCmd.Flags().BoolVarP(&deleteAll, "all", "a", false, "If Selected all known forwards will be deleted")

}

func runDelete(cmd *cobra.Command, args []string) {
	_ = cmd
	_ = args
	ui.PrintCommandHeader(internal.FORWARD_DELETE_SHORT, internal.FORWARD_DELETE_DESCRIPTION)

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

	var toDelete []models.ForwardEntity
	if deleteAll {
		toDelete = allForwards
	} else {

		if len(allForwards) == 1 {
			toDelete = allForwards
		} else {

			selectedForwards, ok := wizard.SelectMany(allForwards, "Select forwards to delete", func(f models.ForwardEntity) string {
				return f.GetAsString()
			})

			if !ok || len(selectedForwards) == 0 {
				logger.Attention("No forwards selected.")
				logger.Warn("No forwards selected.")
				os.Exit(0)

			} else {
				toDelete = selectedForwards
			}

		}
	}

	result, _ := pterm.DefaultInteractiveConfirm.Show("Are you really sure you want to delete these forwards?")

	if result {
		for _, dlt := range toDelete {
			err := forwardRepo.Delete(dlt.ID)
			if err != nil {
				logger.Failure("Forward Deletion failed.")
				logger.Error("Forward Deletion failed.", err)

				os.Exit(1)
			}
		}
	}

}
