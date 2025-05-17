package k8s

import (
	stdcontext "context"
	"github.com/davidemaggi/koncierge/internal/container"
	"github.com/pterm/pterm"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func GetCurrentNamespaceForContext(kubeconfig string, contextName string) string {

	config := clientcmd.GetConfigFromFileOrDie(kubeconfig)
	if config != nil {

		return config.Contexts[contextName].Namespace
	}
	return ""
}

func (k *KubeService) GetAllNameSpaces() ([]string, error) {

	logger := container.App.Logger

	var ret []string

	namespaces, err := k.client.CoreV1().Namespaces().List(stdcontext.TODO(), metav1.ListOptions{})

	if err != nil {
		logger.Error("Error Retrieving Namespaces")
		os.Exit(1)

	}

	for _, ns := range namespaces.Items {
		ret = append(ret, ns.Name)
	}

	return ret, nil
}

func SetDefaultNamespaceForContext(kubeConfig string, contextName string, namespace string) error {
	var logger = container.App.Logger
	rawConfig := clientcmd.GetConfigFromFileOrDie(kubeConfig)
	// Get the context to update
	ctx, exists := rawConfig.Contexts[contextName]
	if !exists {
		logger.Error("Could not find context " + pterm.Red(contextName) + "in " + pterm.Red(kubeConfig))
		os.Exit(1)

	}

	// Set the namespace
	ctx.Namespace = namespace

	// Save the updated config
	err := clientcmd.ModifyConfig(clientcmd.NewDefaultPathOptions(), *rawConfig, false)
	if err != nil {
		logger.Error("Failed to save updated kubeconfig " + pterm.Red(err))
		os.Exit(1)

	}

	return nil
}
