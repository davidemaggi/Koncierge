package k8s

import (
	stdcontext "context"
	"github.com/davidemaggi/koncierge/internal/config"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/pterm/pterm"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func GetCurrentNamespaceForContext(kubeconfig string, contextName string) string {

	config := clientcmd.GetConfigFromFileOrDie(kubeconfig)
	if config != nil {

		return config.Contexts[contextName].Namespace
	}
	return ""
}

func GetAllNameSpaces(kubeconfig string) ([]string, error) {

	logger := container.App.Logger

	var ret []string

	err := ConnectToCluster(config.KubeConfigFile)

	if err != nil {
		logger.Error("Cannot Connect to cluster")
		return nil, err
	}

	namespaces, err := k8sClient.CoreV1().Namespaces().List(stdcontext.TODO(), metav1.ListOptions{})

	if err != nil {
		logger.Error("Error Retrieving Namespaces")
		return nil, err
	}

	for _, ns := range namespaces.Items {
		ret = append(ret, ns.Name)
	}

	return ret, nil
}

func SetDefaultNamespaceForContext(contextName string, namespace string) error {
	rawConfig := clientcmd.GetConfigFromFileOrDie(k8sConfigPath)
	// Get the context to update
	ctx, exists := rawConfig.Contexts[contextName]
	if !exists {
		k8sLogger.Error("Could not find context " + pterm.Red(contextName) + "in " + pterm.Red(k8sConfigPath))
	}

	// Set the namespace
	ctx.Namespace = namespace

	// Save the updated config
	err := clientcmd.ModifyConfig(clientcmd.NewDefaultPathOptions(), *rawConfig, false)
	if err != nil {
		k8sLogger.Error("Failed to save updated kubeconfig " + pterm.Red(err))

	}

	return nil
}
