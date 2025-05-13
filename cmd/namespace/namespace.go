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
package namespace

import (
	"github.com/davidemaggi/koncierge/internal/config"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/k8s"
	"github.com/pterm/pterm"

	"github.com/spf13/cobra"
)

// namespaceCmd represents the namespace command
var NamespaceCmd = &cobra.Command{
	Use:     "namespace",
	Aliases: []string{"ns"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runCommand,
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// namespaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// namespaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runCommand(cmd *cobra.Command, args []string) {

	logger := container.App.Logger
	err := k8s.ConnectToCluster(config.KubeConfigFile)

	if err != nil {
		logger.Error("Cannot Connect to cluster")
		return
	}

	spaces, err := k8s.GetAllNameSpaces(config.KubeConfigFile)

	if err != nil {
		return
	}

	selectedOption := ""
	current := k8s.GetCurrentNamespaceForContext(config.KubeConfigFile, k8s.GetCurrentContextAsString(config.KubeConfigFile))

	if len(spaces) == 0 {
		logger.Info("No namespace available in " + pterm.Green(config.KubeConfigFile))

		// Display the selected option to the user with a green color for emphasis
	}

	if len(spaces) == 1 {
		selectedOption = spaces[0]
		logger.Info("Only " + pterm.Green("one") + " namespace is available")

		// Display the selected option to the user with a green color for emphasis
	}

	if len(spaces) >= 2 {
		if current == "" {
			current = spaces[0]
		}
		selectedOption, _ = pterm.DefaultInteractiveSelect.WithOptions(spaces).WithDefaultOption(current).Show()

	}

	if selectedOption == current {
		logger.Warn("Selected and Current namespace are the same " + pterm.Yellow("Skipping"))

	}

	logger.Info("Switching to " + pterm.Green(selectedOption))
	err = k8s.SetDefaultNamespaceForContext(k8s.GetCurrentContextAsString(config.KubeConfigFile), selectedOption)

	if err != nil {
		return
	}

}
