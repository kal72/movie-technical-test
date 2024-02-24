package usecase

import (
	"context"
	"movie-technical-test/internal/entity"
	"movie-technical-test/internal/model"
	"movie-technical-test/internal/model/converter"
	"movie-technical-test/internal/repository"
	"movie-technical-test/internal/utils"
	"movie-technical-test/internal/utils/logger"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type MovieUseCase struct {
	DB              *gorm.DB
	Log             *logger.Logger
	Validate        *validator.Validate
	MovieRepository *repository.MovieRepository
}

func NewMovieUseCase(db *gorm.DB, log *logger.Logger, validate *validator.Validate, movieRepository *repository.MovieRepository) *MovieUseCase {
	return &MovieUseCase{
		DB:              db,
		Log:             log,
		Validate:        validate,
		MovieRepository: movieRepository,
	}
}

func (u *MovieUseCase) AddNew(ctx context.Context, request model.MovieRequest) (fiber.Map, error) {
	tx := u.DB.WithContext(ctx)
	var err error

	err = u.Validate.Struct(request)
	if err != nil {
		return nil, err
	}

	newMovie := entity.Movie{
		Title:       request.Title,
		Description: request.Description,
		Rating:      request.Rating,
		Image:       request.Image,
	}

	err = u.MovieRepository.Create(tx, &newMovie)
	if err != nil {
		u.Log.Error(ctx, err, "failed insert new movie")
		return nil, fiber.ErrInternalServerError
	}

	response := fiber.Map{"id": newMovie.ID}
	return response, nil
}

func (u *MovieUseCase) Update(ctx context.Context, request model.MovieRequest) error {
	tx := u.DB.WithContext(ctx)
	defer tx.Rollback()
	var err error
	var movie entity.Movie

	err = u.Validate.Struct(request)
	if err != nil {
		return err
	}

	err = u.MovieRepository.FindByID(tx, &movie, request.ID)
	if err != nil {
		u.Log.Error(ctx, err, "failed get detail movie")
		return fiber.ErrInternalServerError
	}

	if movie.ID == 0 {
		return fiber.ErrNotFound
	}

	movie.Title = request.Title
	movie.Rating = request.Rating
	movie.Description = request.Description
	movie.Image = request.Image

	err = u.MovieRepository.Update(tx, &movie)
	if err != nil {
		u.Log.Error(ctx, err, "failed update movie")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (u *MovieUseCase) Detail(ctx context.Context, id int) (*model.MovieResponse, error) {
	tx := u.DB.WithContext(ctx)
	var movie entity.Movie

	err := u.MovieRepository.FindByID(tx, &movie, id)
	if err != nil {
		u.Log.Error(ctx, err, "failed get detail movie")
		return nil, fiber.ErrInternalServerError
	}

	if movie.ID == 0 {
		return nil, fiber.ErrNotFound
	}

	response := converter.MovieToResponse(movie)
	return &response, nil
}

func (u *MovieUseCase) List(ctx context.Context, page, size int) ([]model.MovieResponse, *model.PageMetadata, error) {
	tx := u.DB.WithContext(ctx)

	limit, offset := utils.CalculateLimitOffset(page, size)
	movies, count, err := u.MovieRepository.FindAllByOffsetLimit(tx, offset, limit)
	if err != nil {
		u.Log.Error(ctx, err, "failed fetch movies")
		return nil, nil, fiber.ErrInternalServerError
	}

	var responses = make([]model.MovieResponse, len(movies))
	for i, movie := range movies {
		responses[i] = converter.MovieToResponse(movie)
	}

	pageMetaData := utils.CalculatePagination(count, size, page)
	return responses, pageMetaData, nil
}
