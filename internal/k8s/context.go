package k8s

import (
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/pterm/pterm"
	"k8s.io/client-go/tools/clientcmd"
)

func GetCurrentContextAsString(kubeconfig string) string {

	config := clientcmd.GetConfigFromFileOrDie(kubeconfig)
	if config != nil {

		return config.CurrentContext
	}
	return ""
}

func GetAllContexts(kubeconfig string) []string {

	logger := container.App.Logger

	config := clientcmd.GetConfigFromFileOrDie(kubeconfig)

	logger.Info("Retrieving Contexts from " + pterm.LightMagenta(kubeconfig))

	if config != nil {

		logger.Info("Current Context is " + pterm.LightMagenta(config.CurrentContext))

		contextNames := make(map[string]any)

		for key, ctx := range config.Contexts {
			contextNames[key] = "Current Namespace: " + pterm.LightMagenta(ctx.Namespace)
		}

		logger.MoreInfo("Other Contexts", contextNames)

		var options []string

		// Generate 100 options and add them to the options slice
		for key, _ := range config.Contexts {
			options = append(options, key)
		}

		return options

	} else {
		logger.Error("Cannot retrieve a valid config from " + pterm.LightMagenta(kubeconfig))
	}
	return nil
}

func SwitchContext(ctx, kubeconfig string) (err error) {
	rawConfig := clientcmd.GetConfigFromFileOrDie(kubeconfig)
	logger := container.App.Logger

	if rawConfig.Contexts[ctx] == nil {
		logger.Error("Context " + pterm.LightRed(kubeconfig) + " doesn't exists.")
		return
	}
	rawConfig.CurrentContext = ctx
	err = clientcmd.ModifyConfig(clientcmd.NewDefaultPathOptions(), *rawConfig, true)

	if err != nil {

		logger.Error("Context " + pterm.LightRed(kubeconfig) + " cannot be changed.")
		return
	}

	k8sCurrentContextName = GetCurrentContextAsString(kubeconfig)

	return
}
