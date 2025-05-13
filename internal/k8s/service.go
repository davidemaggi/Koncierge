package k8s

import (
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetServicesInNamespace(namespace string) []string {

	var ret []string

	services, err := k8sClient.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		k8sLogger.Error("Failed to list services in namespace " + namespace)
	}

	for _, svc := range services.Items {
		ret = append(ret, svc.Name)
	}

	return ret
}
