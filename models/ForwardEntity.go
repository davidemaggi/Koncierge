package models

import (
	"gorm.io/gorm"
)

type ForwardEntity struct {
	gorm.Model

	KubeConfigEntityId uint
	KubeConfig         KubeConfigEntity `gorm:"foreignKey:KubeConfigEntityId"`
	ContextName        string
	Namespace          string
	ForwardType        string
	TargetName         string
	ServicePort        int32
	PodPort            int32
	LocalPort          int32
	AdditionalConfigs  []AdditionalConfigEntity
}

func (f *ForwardEntity) PrintPortToForward() int32 {
	if f.ServicePort != 0 {
		return f.ServicePort
	}
	return f.PodPort
}
