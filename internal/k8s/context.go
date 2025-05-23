package k8s

import (
	"fmt"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/pterm/pterm"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"os"
)

func GetCurrentContextAsStringFromConfig(kubeconfig string) string {

	config := clientcmd.GetConfigFromFileOrDie(kubeconfig)
	if config != nil {

		return config.CurrentContext
	}
	return ""
}

func (k *KubeService) GetCurrentContextAsString() string {

	return k.CurrentContext
}

func GetAllContexts(kubeconfig string) []string {

	logger := container.App.Logger

	config := clientcmd.GetConfigFromFileOrDie(kubeconfig)

	logger.Trace("Retrieving Contexts from " + pterm.LightMagenta(kubeconfig))

	if config != nil {

		logger.Trace("Current Context is " + pterm.LightMagenta(config.CurrentContext))

		contextNames := make(map[string]any)

		for key, ctx := range config.Contexts {
			contextNames[key] = "Current Namespace: " + pterm.LightMagenta(ctx.Namespace)
		}

		logger.MoreTrace("Other Contexts", contextNames)

		var options []string

		// Generate 100 options and add them to the options slice
		for key := range config.Contexts {
			options = append(options, key)
		}

		return options

	} else {
		logger.Error("Cannot retrieve a valid config from "+pterm.LightMagenta(kubeconfig), nil)
		os.Exit(1)

	}
	return nil
}

func SwitchContext(ctx, kubeconfig string) (err error) {
	rawConfig := clientcmd.GetConfigFromFileOrDie(kubeconfig)
	logger := container.App.Logger

	if rawConfig.Contexts[ctx] == nil {
		logger.Error("Context "+pterm.LightRed(kubeconfig)+" doesn't exists.", nil)
		os.Exit(1)

	}
	rawConfig.CurrentContext = ctx
	err = clientcmd.ModifyConfig(clientcmd.NewDefaultPathOptions(), *rawConfig, true)

	if err != nil {

		logger.Error("Context "+pterm.LightRed(kubeconfig)+" cannot be changed.", err)
		os.Exit(1)
	}

	return nil
}

func MergeContexts(contextsToCopy []string, fromPath string, toPath string) {

	logger := container.App.Logger

	sourceConfig, err := clientcmd.LoadFromFile(fromPath)
	if err != nil {
		logger.Error("Cannot load Source config file: "+fromPath, err)
	}

	targetConfig, err := clientcmd.LoadFromFile(toPath)
	if err != nil {
		// If target config doesn't exist, initialize a new one
		logger.Warn("Target Config doesn't exist, creating: " + fromPath)

		if os.IsNotExist(err) {
			targetConfig = api.NewConfig()
		} else {
			logger.Error("Cannot load Target config file: "+fromPath, err)
		}
	}

	for _, ctxName := range contextsToCopy {
		ctx, ok := sourceConfig.Contexts[ctxName]
		if !ok {
			logger.Warn(fmt.Sprintf("Context %q not found in source config\n", ctxName))

			continue
		}

		clusterName := ctx.Cluster
		authInfoName := ctx.AuthInfo

		// Copy context
		targetConfig.Contexts[ctxName] = ctx

		// Copy cluster
		if cluster, ok := sourceConfig.Clusters[clusterName]; ok {
			targetConfig.Clusters[clusterName] = cluster
		} else {
			logger.Warn(fmt.Sprintf("Warning: Cluster %q for context %q not found\n", clusterName, ctxName))
		}

		// Copy authInfo
		if authInfo, ok := sourceConfig.AuthInfos[authInfoName]; ok {
			targetConfig.AuthInfos[authInfoName] = authInfo
		} else {
			logger.Warn(fmt.Sprintf("Warning: AuthInfo %q for context %q not found\n", authInfoName, ctxName))

		}
	}

	if err := clientcmd.WriteToFile(*targetConfig, toPath); err != nil {
		logger.Error("Error Saving Target Config", err)
		os.Exit(1)
	}
	logger.Success("Contexts copied successfully.")

}
