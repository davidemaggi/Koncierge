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
	TargetPort         int32
	LocalPort          int32
	AdditionalConfigs  []AdditionalConfigEntity
}
