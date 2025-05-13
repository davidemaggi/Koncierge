package k8s

import (
	"github.com/davidemaggi/koncierge/internal/container"
	logger "github.com/davidemaggi/koncierge/internal/logger"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var k8sClient *kubernetes.Clientset
var k8sConfig *rest.Config
var k8sConfigPath string
var k8sCurrentContextName string
var k8sLogger *logger.Logger

// var k8sCurrentContext *kubernetes.
var k8sError error

func ConnectToCluster(kubeconfig string) error {
	k8sLogger = container.App.Logger
	k8sConfig, k8sError = clientcmd.BuildConfigFromFlags("", kubeconfig)
	if k8sError != nil {
		return k8sError
	}
	k8sConfigPath = kubeconfig
	k8sCurrentContextName = GetCurrentContextAsString(kubeconfig)

	k8sClient, k8sError = kubernetes.NewForConfig(k8sConfig)
	if k8sError != nil {

		k8sLogger.Error("Cannot connect to cluster")
		return k8sError
	}

	return nil
}
