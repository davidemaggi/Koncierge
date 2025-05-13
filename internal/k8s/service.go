package k8s

import (
	"github.com/davidemaggi/koncierge/internal"
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

func GetServicePorts(namespace, serviceName string) []internal.ServicePortDto {

	var ret []internal.ServicePortDto
	svc, err := k8sClient.CoreV1().Services(namespace).Get(context.TODO(), serviceName, metav1.GetOptions{})
	if err != nil {
		k8sLogger.Error("Failed to get service " + serviceName + " in namespace " + namespace)
		return ret
	}

	for _, port := range svc.Spec.Ports {

		var tmp internal.ServicePortDto

		tmp.ServicePort = port.Port
		tmp.Protocol = string(port.Protocol)
		tmp.PodPort = port.TargetPort.IntVal
		ret = append(ret, tmp)
	}

	return ret
}
