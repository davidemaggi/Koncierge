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
	ServicePort       int32
	PodPort           int32
	LocalPort         int32
	AdditionalConfigs []AdditionalConfigDto
}

func (f *ForwardDto) PortForForward() int32 {
	if f.PodPort != 0 {
		return f.PodPort
	}
	return f.ServicePort
}

func (f *ForwardDto) PrintPortToForward() int32 {
	if f.ServicePort != 0 {
		return f.ServicePort
	}
	return f.PodPort
}

func FromForwardEntity(entity models.ForwardEntity) ForwardDto {

	var ret ForwardDto

	ret.ServicePort = entity.ServicePort
	ret.PodPort = entity.PodPort
	ret.TargetName = entity.TargetName
	ret.ForwardType = entity.ForwardType
	ret.ContextName = entity.ContextName
	ret.Namespace = entity.Namespace
	ret.KubeconfigPath = entity.KubeConfig.KubeconfigPath
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

type ServicePodPortKey struct {
	ServicePort int32
	PodPort     int32
}
