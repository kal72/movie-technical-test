package converter

import (
	"movie-technical-test/internal/entity"
	"movie-technical-test/internal/model"
	"movie-technical-test/internal/utils"
)

func MovieToResponse(entity entity.Movie) model.MovieResponse {
	return model.MovieResponse{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Rating:      entity.Rating,
		Image:       entity.Image,
		CreatedAt:   entity.CreatedAt.Format(utils.DateFormat),
		UpdatedAt:   entity.UpdatedAt.Format(utils.DateFormat),
	}
}
