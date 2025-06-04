package k8s

import (
	"fmt"
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/container"
	"golang.org/x/net/context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
)

func (k *KubeService) GetFirstPodForService(namespace string, serviceName string) (string, error) {
	// Get the service

	svc, err := k.client.CoreV1().Services(namespace).Get(context.TODO(), serviceName, v1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to get service: %w", err)
	}

	// Extract the selector
	selector := svc.Spec.Selector
	if len(selector) == 0 {
		return "", fmt.Errorf("service has no selector")
	}

	// Convert selector map to label selector string
	selectorString := ""
	for k, v := range selector {
		selectorString += fmt.Sprintf("%s=%s,", k, v)
	}
	selectorString = selectorString[:len(selectorString)-1] // trim trailing comma

	// Get pods matching the selector
	pods, err := k.client.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{
		LabelSelector: selectorString,
	})
	if err != nil {
		return "", fmt.Errorf("failed to list pods: %w", err)
	}
	if len(pods.Items) == 0 {
		return "", fmt.Errorf("no pods found for service")
	}

	// Return the name of the first pod
	return pods.Items[0].Name, nil
}
func (k *KubeService) GetPodsInNamespace(namespace string) []string {

	var ret []string
	logger := container.App.Logger
	pods, err := k.client.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		logger.Error("Failed to list pods: ", err)
		logger.Failure("Failed to list pods")
		os.Exit(1)
	}
	if len(pods.Items) == 0 {
		logger.Error("No pods found", nil)
		logger.Failure("No pods found")
		os.Exit(0)

	}

	for _, item := range pods.Items {
		ret = append(ret, item.Name)
	}

	// Return the name of the first pod
	return ret
}

func (k *KubeService) GetPodPorts(namespace string, podName string) []internal.ServicePortDto {
	var logger = container.App.Logger

	var ret []internal.ServicePortDto

	pod, err := k.client.CoreV1().Pods(namespace).Get(context.TODO(), podName, v1.GetOptions{})
	if err != nil {
		logger.Error("Failed to get pod "+podName+" in namespace "+namespace, err)
		logger.Failure("Failed to get pod " + podName + " in namespace " + namespace)
		os.Exit(1)
	}

	for _, port := range pod.Spec.Containers[0].Ports {

		var tmp internal.ServicePortDto

		tmp.ServicePort = port.ContainerPort
		tmp.Protocol = string(port.Protocol)
		tmp.PodPort = port.ContainerPort
		tmp.ServicePort = port.ContainerPort
		//tmp.Podname = port.
		ret = append(ret, tmp)
	}

	return ret
}
