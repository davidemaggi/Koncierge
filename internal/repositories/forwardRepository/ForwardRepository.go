package forwardRepository

import (
	"github.com/davidemaggi/koncierge/internal"
	"github.com/davidemaggi/koncierge/internal/db"
	logger2 "github.com/davidemaggi/koncierge/internal/logger"
	"github.com/davidemaggi/koncierge/internal/repositories"
	"github.com/davidemaggi/koncierge/models"
	"gorm.io/gorm"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

type ForwardRepository interface {
	repositories.Repository[models.ForwardEntity]
	CreateFromDto(fwd internal.ForwardDto)
	MoveToCtx(id uint, moveCtx string)
}

type GormForwardRepository struct {
	*repositories.GormRepository[models.ForwardEntity]
}

func NewForwardRepository(db *gorm.DB) *GormForwardRepository {
	return &GormForwardRepository{repositories.NewGormRepository[models.ForwardEntity](db)}
}

func (r *GormForwardRepository) CopyToCtx(ida uint, moveCtx string) {
	var existingFwd models.ForwardEntity
	var searchFwd models.ForwardEntity
	var logger = logger2.NewLogger(true)

	searchFwd.ID = ida

	err := db.GetDB().First(&existingFwd, searchFwd).Error

	if err != nil {
		logger.Failure("Cannot Find Forward")

		logger.Error("Cannot Find Forward", err)
		os.Exit(1)

	}

	var addConfigs []models.AdditionalConfigEntity

	for _, config := range existingFwd.AdditionalConfigs {
		addConfigs = append(addConfigs, models.AdditionalConfigEntity{
			ForwardEntityId: config.ForwardEntityId,
			Name:            config.Name,
			ConfigType:      config.ConfigType,
			Values:          config.Values,
		})
	}

	cp := models.ForwardEntity{
		KubeConfigEntityId: existingFwd.KubeConfigEntityId,
		ContextName:        moveCtx,
		Namespace:          existingFwd.Namespace,
		ForwardType:        existingFwd.ForwardType,
		TargetName:         existingFwd.TargetName,
		TargetPort:         existingFwd.TargetPort,
		LocalPort:          existingFwd.LocalPort,
		AdditionalConfigs:  addConfigs,
	}

	db.GetDB().Create(&cp)

}

func (r *GormForwardRepository) MoveToCtx(ida uint, moveCtx string) {
	var existingFwd models.ForwardEntity
	var searchFwd models.ForwardEntity
	var logger = logger2.NewLogger(true)

	searchFwd.ID = ida

	err := db.GetDB().First(&existingFwd, searchFwd).Error

	if err != nil {
		logger.Failure("Cannot Find Forward")

		logger.Error("Cannot Find Forward", err)
		os.Exit(1)

	}

	existingFwd.ContextName = moveCtx

	db.GetDB().Save(existingFwd)

}

func (r *GormForwardRepository) CreateFromDto(fwd internal.ForwardDto) {

	var logger = logger2.NewLogger(true)

	var existingConfig models.KubeConfigEntity

	err := db.GetDB().First(&existingConfig, models.KubeConfigEntity{KubeconfigPath: fwd.KubeconfigPath}).Error
	defaultFile := ""
	if home := homedir.HomeDir(); home != "" {
		defaultFile = filepath.Join(home, ".kube", "config")
	}

	if err != nil {
		newKc := &models.KubeConfigEntity{
			Name:           filepath.Base(fwd.KubeconfigPath),
			KubeconfigPath: fwd.KubeconfigPath,
			IsDefault:      fwd.KubeconfigPath == defaultFile,
		}

		db.GetDB().Create(newKc)

		err = db.GetDB().First(&existingConfig, models.KubeConfigEntity{KubeconfigPath: fwd.KubeconfigPath}).Error
		if err != nil {

			logger.Error("Error creating new KubeConfig", err)
			os.Exit(1)
		}
	}

	var existingForward models.ForwardEntity

	err = db.GetDB().First(&existingForward, models.ForwardEntity{ContextName: fwd.ContextName, TargetName: fwd.TargetName, TargetPort: fwd.TargetPort, KubeConfigEntityId: existingConfig.ID}).Error

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
			logger.Failure("Cannot create Forward")

			logger.Error("Cannot create Forward", err)
			os.Exit(1)

		} else {
			logger.Info("Forward created")
		}

	} else {
		logger.Attention("The forward entity already exists")
		logger.Warn("The forward entity already exists")
		os.Exit(1)
	}

}
