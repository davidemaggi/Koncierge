package forward

import (
	"fmt"
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/db"
	"github.com/davidemaggi/koncierge/internal/repositories"
	"github.com/davidemaggi/koncierge/internal/ui"
	"github.com/davidemaggi/koncierge/internal/utils"
	"github.com/davidemaggi/koncierge/models"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var FwdListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "all"},
	Short:   internal.FORWARD_LIST_SHORT,
	Long:    internal.FORWARD_LIST_DESCRIPTION,
	Run:     runList,
}

func init() {

}

func runList(cmd *cobra.Command, args []string) {
	_ = cmd
	_ = args
	ui.PrintCommandHeader(internal.FORWARD_LIST_SHORT, internal.FORWARD_LIST_DESCRIPTION)

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

	tableData1 := pterm.TableData{
		{"N.", "Config", "Context", "Namespace", "Target", "➡️", "Local", "Extra"},
	}

	for id, forward := range allForwards {

		var extras []string

		if len(forward.AdditionalConfigs) > 0 {

			for _, config := range forward.AdditionalConfigs {
				if config.ConfigType == internal.ConfigTypeSecret {
					extras = append(extras, "🔑")
				}
				if config.ConfigType == internal.ConfigTypeMap {
					extras = append(extras, "🔧")

				}
			}

		}
		ex := strings.Join(utils.DistinctStrings(extras), " ")
		tableData1 = append(tableData1, []string{fmt.Sprintf("%d", id+1), forward.KubeConfig.Name, forward.ContextName, forward.Namespace, forward.GetAsShortString(), ex})
	}

	// Create a table with a header and the defined data, then render it
	err = pterm.DefaultTable.WithHasHeader().WithData(tableData1).Render()
	if err != nil {
		logger.Failure("Error Rendering table")
		logger.Error("Error Rendering table", err)
		os.Exit(1)
	}

}
