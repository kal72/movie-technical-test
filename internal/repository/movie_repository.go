package repository

import (
	"movie-technical-test/internal/entity"

	"gorm.io/gorm"
)

type MovieRepository struct {
	Repository[entity.Movie]
}

func NewMovieRepository() *MovieRepository {
	return &MovieRepository{}
}

func (r *MovieRepository) FindAllByOffsetLimit(tx *gorm.DB, offset, limit int) ([]entity.Movie, int64, error) {
	var movies []entity.Movie
	var count int64

	err := tx.Model(&entity.Movie{}).Count(&count).Limit(limit).Offset(offset).Find(&movies).Error
	if err != nil {
		return nil, count, err
	}

	return movies, count, nil
}
