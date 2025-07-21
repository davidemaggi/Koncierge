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

	err := db.GetDB().Preload("AdditionalConfigs").First(&existingFwd, searchFwd).Error

	if err != nil {
		logger.Failure("Cannot Find Forward")

		logger.Error("Cannot Find Forward", err)
		os.Exit(1)

	}

	cp := models.ForwardEntity{
		KubeConfigEntityId: existingFwd.KubeConfigEntityId,
		ContextName:        moveCtx,
		Namespace:          existingFwd.Namespace,
		ForwardType:        existingFwd.ForwardType,
		TargetName:         existingFwd.TargetName,
		ServicePort:        existingFwd.ServicePort,
		PodPort:            existingFwd.PodPort,
		LocalPort:          existingFwd.LocalPort,
	}

	if r.Exists(cp) {
		logger.Attention("The same forward already exists: " + cp.GetAsString())
		logger.Warn("The same forward already exists")
		//os.Exit(1)
	} else {

		db.GetDB().Create(&cp)

		for _, config := range existingFwd.AdditionalConfigs {
			cpConf := models.AdditionalConfigEntity{
				ForwardEntityId: cp.ID,
				Name:            config.Name,
				ConfigType:      config.ConfigType,
				Values:          config.Values,
			}
			db.GetDB().Create(&cpConf)
		}
		logger.Success("Forward Copied Successfully: " + cp.GetAsString())

	}
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

	if r.Exists(existingFwd) {
		logger.Attention("The same forward already exists: " + existingFwd.GetAsString())
		logger.Warn("The same forward already exists")
		//os.Exit(1)
	} else {

		db.GetDB().Save(existingFwd)
		logger.Success("Forward Moved Successfully: " + existingFwd.GetAsString())

	}
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

	newFwd := &models.ForwardEntity{
		KubeConfigEntityId: existingConfig.ID,
		ContextName:        fwd.ContextName,
		Namespace:          fwd.Namespace,
		ForwardType:        fwd.ForwardType,
		TargetName:         fwd.TargetName,
		ServicePort:        fwd.ServicePort,
		PodPort:            fwd.PodPort,
		LocalPort:          fwd.LocalPort,
	}

	for _, config := range fwd.AdditionalConfigs {
		newFwd.AdditionalConfigs = append(newFwd.AdditionalConfigs, models.AdditionalConfigEntity{
			Name: config.Name, Values: config.Values, ConfigType: config.ConfigType,
		})
	}

	if r.Exists(*newFwd) {
		logger.Attention("The same forward already exists: " + newFwd.GetAsString())
		logger.Warn("The same forward already exists")
		os.Exit(1)
	}

	logger.Info("Creating Forward")
	err = db.GetDB().Create(newFwd).Error

	if err != nil {
		logger.Failure("Cannot create Forward")

		logger.Error("Cannot create Forward", err)
		os.Exit(1)

	} else {
		logger.Info("Forward created")
	}

}

func (r *GormForwardRepository) Exists(entity models.ForwardEntity) bool {
	var count int64

	db.GetDB().
		Model(&models.ForwardEntity{}).
		Where("kube_config_entity_id = ? AND context_name = ? AND target_name = ? AND service_port = ? AND pod_port = ?",
			entity.KubeConfigEntityId,
			entity.ContextName,
			entity.TargetName,
			entity.ServicePort,
			entity.PodPort,
		).Where("id <> ?", entity.ID).Count(&count)

	return count > 0
}
