package k8s

import (
	"fmt"
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/container"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *KubeService) GetSecretsInNamespace(namespace string) ([]internal.AdditionalConfigDto, error) {
	var logger = container.App.Logger

	var ret []internal.AdditionalConfigDto

	secrets, err := k.client.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to list secrets"), err)

	}

	for _, secret := range secrets.Items {

		var vals []string

		for key := range secret.Data {

			vals = append(vals, key)

		}

		ret = append(ret, internal.AdditionalConfigDto{
			Name:   secret.Name,
			Values: vals,
		})

	}

	return ret, nil
}

func (k *KubeService) GetSecretValue(namespace string, secretName string, value string) (string, error) {
	var logger = container.App.Logger

	secrets, err := k.client.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to list secrets"), err)

	}

	for _, secret := range secrets.Items {

		if secret.Name == secretName {
			return string(secret.Data[value]), nil
		}

	}

	return "Not Found", nil
}
