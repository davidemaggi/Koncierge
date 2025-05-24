package forward

import (
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/davidemaggi/koncierge/internal/db"
	"github.com/davidemaggi/koncierge/internal/k8s"
	"github.com/davidemaggi/koncierge/internal/repositories/forwardRepository"
	"github.com/davidemaggi/koncierge/internal/ui"
	"github.com/davidemaggi/koncierge/internal/wizard"
	"github.com/spf13/cobra"
	"os"
)

var FwdAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"fwd add"},
	Short:   internal.FORWARD_ADD_SHORT,
	Long:    internal.FORWARD_ADD_DESCRIPTION,
	Run:     runAdd,
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

	_ = cmd
	_ = args

	logger := container.App.Logger

	fwdRepo := forwardRepository.NewForwardRepository(db.GetDB())

	ui.PrintCommandHeader(internal.FORWARD_ADD_SHORT, internal.FORWARD_ADD_DESCRIPTION)

	fwd := wizard.BuildForward()

	done := false

	for !done {

		addConfig, ok := wizard.SelectOne([]string{internal.BoolYes, internal.BoolNo}, "Do you want to add an additional config?", func(t string) string {
			return t
		}, internal.BoolNo)

		if !ok {
			os.Exit(1)
		}

		if addConfig == internal.BoolNo {
			done = true
			continue

		}

		addType, ok := wizard.SelectOne([]string{internal.ConfigTypeMap, internal.ConfigTypeSecret}, "Which kind of config?", func(t string) string {
			return t

		}, internal.ConfigTypeSecret)

		if !ok {
			os.Exit(1)
		}

		kubeService, err := k8s.ConnectToClusterAndContext(fwd.KubeconfigPath, fwd.ContextName)

		if err != nil {
			os.Exit(1)
		}

		var confs []internal.AdditionalConfigDto

		if addType == internal.ConfigTypeSecret {

			confs, err = kubeService.GetSecretsInNamespace(fwd.Namespace)
			if err != nil {
				os.Exit(1)
			}
		}

		if addType == internal.ConfigTypeMap {
			confs, err = kubeService.GetConfigMapsInNamespace(fwd.Namespace)
			if err != nil {
				os.Exit(1)
			}
		}

		SelectConf, ok := wizard.SelectOne(confs, "Which one", func(dto internal.AdditionalConfigDto) string {
			return dto.Name
		}, "")

		if !ok {
			os.Exit(1)
		}

		SelectVals, ok := wizard.SelectMany(SelectConf.Values, "Select Values", func(s string) string {
			return s
		})

		if !ok {
			os.Exit(1)
		}

		fwd.AdditionalConfigs = append(fwd.AdditionalConfigs, internal.AdditionalConfigDto{
			Name:       SelectConf.Name,
			ConfigType: addType,
			Values:     SelectVals,
		})

	}
	fwdRepo.CreateFromDto(fwd)

	logger.Success("Forward Created!")

}
