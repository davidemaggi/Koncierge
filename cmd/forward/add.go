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
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/db"
	"github.com/davidemaggi/koncierge/internal/k8s"
	"github.com/davidemaggi/koncierge/internal/repositories/forwardRepository"
	"github.com/davidemaggi/koncierge/internal/wizard"
	"github.com/spf13/cobra"
)

// forwardCmd represents the forward command
var FwdAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"fwd add"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runAdd,
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// forwardCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// forwardCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runAdd(cmd *cobra.Command, args []string) {

	//logger := container.App.Logger
	fwdRepo := forwardRepository.NewForwardRepository(db.GetDB())

	fwd := wizard.BuildForward()

	done := false

	for !done {

		addConfig, _ := wizard.SelectOne([]string{internal.BoolYes, internal.BoolNo}, "Do you want to add an additional config?", func(t string) string {
			return t
		}, internal.BoolNo)

		if addConfig == internal.BoolNo {
			done = true
			continue

		}

		addType, _ := wizard.SelectOne([]string{internal.ConfigTypeMap, internal.ConfigTypeSecret}, "Which kind of config?", func(t string) string {
			return t

		}, internal.ConfigTypeSecret)

		kubeService, _ := k8s.ConnectToClusterAndContext(fwd.KubeconfigPath, fwd.ContextName)

		var confs []internal.AdditionalConfigDto

		if addType == internal.ConfigTypeSecret {

			confs, _ = kubeService.GetSecretsInNamespace(fwd.Namespace)

		}

		if addType == internal.ConfigTypeMap {
			confs, _ = kubeService.GetConfigMapsInNamespace(fwd.Namespace)

		}

		SelectConf, _ := wizard.SelectOne(confs, "Which one", func(dto internal.AdditionalConfigDto) string {
			return dto.Name
		}, "")

		SelectVals, _ := wizard.SelectMany(SelectConf.Values, "Select Values", func(s string) string {
			return s
		})

		fwd.AdditionalConfigs = append(fwd.AdditionalConfigs, internal.AdditionalConfigDto{
			Name:       SelectConf.Name,
			ConfigType: addType,
			Values:     SelectVals,
		})

	}
	fwdRepo.CreateFromDto(fwd)
}
