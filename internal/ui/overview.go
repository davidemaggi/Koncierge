package ui

import (
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/pterm/pterm"
	"os"
)

func PrintForwardOverview(fwd internal.ForwardDto, configs map[string]string) {

	var lg = container.App.Logger

	// Define a map with interesting stuff
	overview := map[string]any{
		fwd.PodName: fwd.TargetPort,
		"localhost": fwd.LocalPort,
	}

	// Log a debug level message with arguments from the map
	lg.Get().Info("Forwarding "+pterm.LightMagenta(fwd.TargetName)+" ("+fwd.ForwardType+")", lg.Get().ArgsFromMap(overview))

	if len(fwd.AdditionalConfigs) != 0 {
		lg.Get().Info("The following Configurations are linked to the forward")

		tableData1 := pterm.TableData{
			{"Type", "Name", "Value"},
		}

		for _, additionalConf := range fwd.AdditionalConfigs {

			for _, value := range additionalConf.Values {

				tmpStr := additionalConf.Name + "." + value

				tableData1 = append(tableData1, []string{additionalConf.ConfigType, tmpStr, configs[tmpStr]})
			}

		}

		// Create a table with a header and the defined data, then render it
		err := pterm.DefaultTable.WithHasHeader().WithData(tableData1).Render()
		if err != nil {
			os.Exit(1)
		}

	}

}
