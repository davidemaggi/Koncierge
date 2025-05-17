package models

import (
	"gorm.io/gorm"
)

type AdditionalConfigEntity struct {
	gorm.Model

	ForwardEntityId uint
	Forward         ForwardEntity `gorm:"foreignKey:ForwardEntityId"`
	Name            string
	ConfigType      string
	Values          StringArray
}
