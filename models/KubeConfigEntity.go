package models

import "gorm.io/gorm"

type KubeConfigEntity struct {
	gorm.Model
	KubeconfigPath string
	Name           string
	IsDefault      bool
	Forwards       []ForwardEntity
}
