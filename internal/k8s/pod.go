package k8s

import (
	"fmt"
	"golang.org/x/net/context"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
