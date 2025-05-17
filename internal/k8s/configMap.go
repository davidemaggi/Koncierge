package k8s

import (
	"fmt"
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/container"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *KubeService) GetConfigMapsInNamespace(namespace string) ([]internal.AdditionalConfigDto, error) {
	var logger = container.App.Logger

	var ret []internal.AdditionalConfigDto

	maps, err := k.client.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to list config maps"), err)

	}

	for _, m := range maps.Items {

		var vals []string

		for key := range m.Data {

			vals = append(vals, key)

		}

		ret = append(ret, internal.AdditionalConfigDto{
			Name:   m.Name,
			Values: vals,
		})

	}

	return ret, nil
}
func (k *KubeService) GetConfigMapValue(namespace string, mapName string, value string) (string, error) {
	var logger = container.App.Logger

	maps, err := k.client.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to list Config Map"), err)

	}

	for _, m := range maps.Items {

		if m.Name == mapName {
			return m.Data[value], nil
		}

	}

	return "Not Found", nil
}
