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

func (r *MovieRepository) FindAllByOffsetLimit(tx *gorm.DB, offset, limit int) ([]entity.Movie, error) {
	var movies []entity.Movie

	err := tx.Limit(limit).Offset(offset).Find(&movies).Error
	if err != nil {
		return nil, err
	}

	return movies, nil
}
