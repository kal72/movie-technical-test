package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type Repository[T any] struct{}

func (r *Repository[T]) Create(db *gorm.DB, entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) Update(db *gorm.DB, entity *T) error {
	return db.Save(entity).Error
}

func (r *Repository[T]) Delete(db *gorm.DB, entity *T) error {
	return db.Delete(entity).Error
}

func (r *Repository[T]) Count(db *gorm.DB) (int64, error) {
	var total int64
	err := db.Model(new(T)).Count(&total).Error
	return total, err
}

func (r *Repository[T]) FindBy(db *gorm.DB, entity *T, column string, value any) error {
	return db.Where(fmt.Sprintf("%s = ?", column), value).Take(entity).Error
}
