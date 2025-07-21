package models

import (
	"fmt"
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
	AdditionalConfigs  []AdditionalConfigEntity `gorm:"foreignKey:ForwardEntityId"`
}

func (f *ForwardEntity) PrintPortToForward() int32 {
	if f.ServicePort != 0 {
		return f.ServicePort
	}
	return f.PodPort
}

func (f *ForwardEntity) GetAsString() string {
	return fmt.Sprintf("%s.%s.%s:%d ➡️ localhost:%d", f.ContextName, f.Namespace, f.TargetName, f.PrintPortToForward(), f.LocalPort)
}
