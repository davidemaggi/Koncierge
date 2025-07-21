package ui

import (
	"fmt"
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/version"
	"github.com/pterm/pterm"
	"os"
)

func PrintCommandHeader(name string, description string) {

	var lg = container.App.Logger
	lg.Get().Trace("Koncierge: " + version.Version() + " - " + version.Name())
	lg.Get().Trace(name + ": " + description)

}

func PrintCurrentStatus(currentConfig string, currentCtx string, currentNs string) {

	pterm.Println("📄 Current KubeConfig: " + pterm.Green(currentConfig))
	pterm.Println("🗃️ Current Context: " + pterm.Green(currentCtx))
	pterm.Println("🪪 Default Namespace: " + pterm.Green(currentNs))

}

func PrintForwardOverview(fwd internal.ForwardDto, configs map[string]string) {

	var lg = container.App.Logger

	// Define a map with interesting stuff

	// Log a debug level message with arguments from the map

	//pterm.DefaultBasicText.Println(fmt.Sprintf("%s.%s.%s:%s ➡️ localhost:%s", pterm.Gray(fwd.ContextName), pterm.Gray(fwd.Namespace), fwd.TargetName, pterm.Green(fwd.TargetPort), pterm.LightBlue(fwd.LocalPort)))

	tableData := pterm.TableData{
		{"KubeConfig", fwd.KubeconfigPath},
		{"Context", fwd.ContextName},
		{"Namespace", fwd.Namespace},
		{"Remote", fmt.Sprintf("%s:%s", fwd.TargetName, pterm.Green(fwd.PrintPortToForward()))},
		{"Local", fmt.Sprintf("localhost:%s", pterm.LightBlue(fwd.LocalPort))},
	}

	//_ = pterm.DefaultTable.WithBoxed().WithData(tableData).Render()

	if len(fwd.AdditionalConfigs) != 0 {
		tableData = append(tableData, []string{"", ""})

		for _, additionalConf := range fwd.AdditionalConfigs {

			for _, value := range additionalConf.Values {

				tmpStr := additionalConf.Name + "." + value
				tmpStrName := pterm.Gray(additionalConf.Name+".") + value
				icon := "🚽"

				switch additionalConf.ConfigType {
				case internal.ConfigTypeSecret:
					icon = "🔑"
					break
				case internal.ConfigTypeMap:
					icon = "🔧"
					break
				default:
					icon = "🚽"
					break
				}

				tableData = append(tableData, []string{fmt.Sprintf("%s %s", icon, tmpStrName), configs[tmpStr]})
			}

		}

		// Create a table with a header and the defined data, then render it

	}
	err := pterm.DefaultTable.WithBoxed().WithData(tableData).Render()
	if err != nil {
		os.Exit(1)
	}
	_ = lg
}
