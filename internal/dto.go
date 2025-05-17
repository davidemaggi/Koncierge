package internal

import (
	"github.com/davidemaggi/koncierge/models"
)

type ForwardDto struct {
	KubeconfigPath    string
	ContextName       string
	Namespace         string
	ForwardType       string
	TargetName        string
	PodName           string
	TargetPort        int32
	LocalPort         int32
	AdditionalConfigs []AdditionalConfigDto
}

func FromForwardEntity(entity models.ForwardEntity) ForwardDto {

	var ret ForwardDto

	ret.TargetPort = entity.TargetPort
	ret.TargetName = entity.TargetName
	ret.ForwardType = entity.ForwardType
	ret.ContextName = entity.ContextName
	ret.Namespace = entity.Namespace
	//ret.KubeconfigPath = entity.
	ret.LocalPort = entity.LocalPort

	for _, config := range entity.AdditionalConfigs {
		ret.AdditionalConfigs = append(ret.AdditionalConfigs, AdditionalConfigDto{
			Name:       config.Name,
			Values:     config.Values,
			ConfigType: config.ConfigType,
		})
	}

	return ret

}

type AdditionalConfigDto struct {
	Name       string
	ConfigType string
	Values     []string
}

type ServicePortDto struct {
	Protocol    string
	ServicePort int32
	PodPort     int32
	PodName     string
}

const (
	ForwardPod     = "📦 Pod"
	ForwardService = "🌐 Service"

	ConfigTypeSecret = "🔑 Secret"
	ConfigTypeMap    = "🔧 ConfigMap"

	BoolYes = "✅ Yes"
	BoolNo  = "🛑 No"
)
