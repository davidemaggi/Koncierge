package forwardRepository

import (
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/db"
	logger2 "github.com/davidemaggi/koncierge/internal/logger"
	"github.com/davidemaggi/koncierge/internal/repositories"
	"github.com/davidemaggi/koncierge/models"
	"gorm.io/gorm"
)

type ForwardRepository interface {
	repositories.Repository[models.ForwardEntity]
	CreateFromDto(fwd internal.ForwardDto)
}

type GormForwardRepository struct {
	*repositories.GormRepository[models.ForwardEntity]
}

func NewForwardRepository(db *gorm.DB) *GormForwardRepository {
	return &GormForwardRepository{repositories.NewGormRepository[models.ForwardEntity](db)}
}
func (r *GormForwardRepository) CreateFromDto(fwd internal.ForwardDto) {

	var logger = logger2.NewLogger(true)

	var existingConfig models.KubeConfigEntity

	err := db.GetDB().First(&existingConfig, models.KubeConfigEntity{KubeconfigPath: fwd.KubeconfigPath}).Error

	if err != nil {
		newKc := &models.KubeConfigEntity{
			Name:           "",
			KubeconfigPath: fwd.KubeconfigPath,
		}

		db.GetDB().Create(newKc)

		err = db.GetDB().First(&existingConfig, models.KubeConfigEntity{KubeconfigPath: fwd.KubeconfigPath}).Error
		if err != nil {

			logger.Error("Error creating new KubeConfig")

		}
	}

	var existingForward models.ForwardEntity

	err = db.GetDB().First(&existingForward, models.ForwardEntity{TargetName: fwd.TargetName, TargetPort: fwd.TargetPort, KubeConfigEntityId: existingConfig.ID}).Error

	newFwd := &models.ForwardEntity{
		KubeConfigEntityId: existingConfig.ID,
		ContextName:        fwd.ContextName,
		Namespace:          fwd.Namespace,
		ForwardType:        fwd.ForwardType,
		TargetName:         fwd.TargetName,
		TargetPort:         fwd.TargetPort,
		LocalPort:          fwd.LocalPort,
	}

	for _, config := range fwd.AdditionalConfigs {
		newFwd.AdditionalConfigs = append(newFwd.AdditionalConfigs, models.AdditionalConfigEntity{
			Name: config.Name, Values: config.Values, ConfigType: config.ConfigType,
		})
	}

	if err != nil {
		logger.Info("Creating Forward")
		err = db.GetDB().Create(newFwd).Error

		if err != nil {
			logger.Error("Cannot create Forward")

		} else {
			logger.Info("Forward created")
		}

	} else {
		logger.Warn("The forward entity already exists")

	}

}
