package repositories

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository[T any] interface {
	GetByID(id uint) (*T, error)
	GetAll() ([]T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(id uint) error
}

type GormRepository[T any] struct {
	db *gorm.DB
}

func NewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{db: db}
}

func (r *GormRepository[T]) GetByID(id uint) (*T, error) {
	var entity T
	err := r.db.Preload(clause.Associations).First(&entity, id).Error
	return &entity, err
}

func (r *GormRepository[T]) GetAll() ([]T, error) {
	var entities []T
	err := r.db.Preload(clause.Associations).Find(&entities).Error
	return entities, err
}

func (r *GormRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *GormRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *GormRepository[T]) Delete(id uint) error {
	var entity T
	return r.db.Delete(&entity, id).Error
}
