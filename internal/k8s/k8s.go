package k8s

import (
	"github.com/davidemaggi/koncierge/internal/container"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeService struct {
	client         *kubernetes.Clientset
	config         *rest.Config
	error          error
	KubeConfigPath string
	CurrentContext string
}

func ConnectToCluster(kubeconfig string) (*KubeService, error) {

	return ConnectToClusterAndContext(kubeconfig, GetCurrentContextAsStringFromConfig(kubeconfig))
}

func ConnectToClusterAndContext(kubeconfig string, contextName string) (*KubeService, error) {
	var logger = container.App.Logger

	configOverrides := &clientcmd.ConfigOverrides{
		CurrentContext: contextName,
	}

	loadingRules := &clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfig}
	ccc := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	k8sConfig, err := ccc.ClientConfig()
	if err != nil {
		return nil, err

	}

	k8sClient, k8sError := kubernetes.NewForConfig(k8sConfig)
	if k8sError != nil {

		logger.Error("Cannot connect to cluster", err)
		return nil, err

	}

	return &KubeService{client: k8sClient, config: k8sConfig, KubeConfigPath: kubeconfig, CurrentContext: contextName, error: nil}, nil

}
func (k *KubeService) GetKubeConfigPath() string {

	return k.KubeConfigPath
}
