package k8s

import (
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/container"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
)

func (k *KubeService) GetServicesInNamespace(namespace string) []string {

	var logger = container.App.Logger
	var ret []string

	services, err := k.client.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error("Failed to list services in namespace "+namespace, err)
		os.Exit(1)
	}

	for _, svc := range services.Items {
		ret = append(ret, svc.Name)
	}

	return ret
}

func (k *KubeService) GetServicePorts(namespace string, serviceName string) []internal.ServicePortDto {
	var logger = container.App.Logger

	var ret []internal.ServicePortDto

	svc, err := k.client.CoreV1().Services(namespace).Get(context.TODO(), serviceName, metav1.GetOptions{})
	if err != nil {
		logger.Error("Failed to get service "+serviceName+" in namespace "+namespace, err)
		os.Exit(1)
	}

	for _, port := range svc.Spec.Ports {

		var tmp internal.ServicePortDto

		tmp.ServicePort = port.Port
		tmp.Protocol = string(port.Protocol)
		tmp.PodPort = port.TargetPort.IntVal
		//tmp.Podname = port.
		ret = append(ret, tmp)
	}

	//pod, _ := k.GetFirstPodForService(namespace, serviceName)

	return ret

	//podPorts := k.GetPodPorts(namespace, pod)

	//return MergeServicePortDtos(ret, podPorts)
}
