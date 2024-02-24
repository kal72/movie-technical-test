package usecase

import (
	"context"
	"movie-technical-test/internal/entity"
	"movie-technical-test/internal/model"
	"movie-technical-test/internal/repository"
	log "movie-technical-test/internal/utils/logger"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbMock(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	sqldb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	mock.MatchExpectationsInOrder(false)
	gormdb, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "sqlmock_db_0",
		DriverName:                "mysql",
		Conn:                      sqldb,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		t.Fatal(err)
	}

	return gormdb, mock
}

func TestMovieUseCase_AddNew(t *testing.T) {
	// Setup
	mockDB, dbmock := DbMock(t)
	dbmock.MatchExpectationsInOrder(false)
	dbmock.ExpectBegin()
	dbmock.ExpectExec("INSERT INTO `movies`").WillReturnResult(sqlmock.NewResult(1, 1))
	dbmock.ExpectCommit()

	mockLog := log.New(logrus.StandardLogger(), "app-unit-test")
	mockValidate := validator.New()
	mockMovieRepo := repository.NewMovieRepository()

	uc := NewMovieUseCase(mockDB, mockLog, mockValidate, mockMovieRepo)

	testCases := []struct {
		name           string
		request        model.MovieCreateRequest
		expectedError  error
		expectedOutput fiber.Map
	}{
		{
			name: "ValidRequest",
			request: model.MovieCreateRequest{
				Title:       "Test Movie",
				Description: "Test Description",
				Rating:      5.0,
				Image:       "test.jpg",
			},
			expectedError:  nil,
			expectedOutput: map[string]interface{}{"id": uint(1)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Invoke the method under test
			response, err := uc.AddNew(context.Background(), tc.request)

			// Assertions
			if tc.expectedError != nil {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expectedOutput, response)
			}
		})
	}
}

func TestMovieUseCase_List(t *testing.T) {
	// Mock data
	rowsCount := sqlmock.NewRows([]string{"count"}).AddRow(2)
	rowsData := sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
		AddRow(1, "Movie 1", "Description 1", 4.5, "image1.jpg", nil, nil).
		AddRow(2, "Movie 2", "Description 2", 3.8, "image2.jpg", nil, nil)

	// Setup
	mockDB, dbmock := DbMock(t)
	dbmock.MatchExpectationsInOrder(false)
	dbmock.ExpectQuery("SELECT count(.+) FROM `movies`").WillReturnRows(rowsCount)
	dbmock.ExpectQuery("SELECT (.+) FROM `movies`").WillReturnRows(rowsData)

	mockLog := log.New(logrus.StandardLogger(), "app-unit-test")
	mockValidate := validator.New()
	mockMovieRepo := repository.NewMovieRepository()

	uc := NewMovieUseCase(mockDB, mockLog, mockValidate, mockMovieRepo)

	// Test cases
	testCases := []struct {
		name           string
		page           int
		size           int
		expectedError  error
		expectedOutput []model.MovieResponse
	}{
		{
			name:          "ValidRequest",
			page:          1,
			size:          10,
			expectedError: nil,
			expectedOutput: []model.MovieResponse{
				{ID: 1, Title: "Movie 1", Description: "Description 1", Rating: 4.5, Image: "image1.jpg", CreatedAt: "01-01-0001", UpdatedAt: "01-01-0001"},
				{ID: 2, Title: "Movie 2", Description: "Description 2", Rating: 3.8, Image: "image2.jpg", CreatedAt: "01-01-0001", UpdatedAt: "01-01-0001"},
			},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Invoke the method under test
			responses, _, err := uc.List(context.Background(), tc.page, tc.size)

			// Assertions
			if tc.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOutput, responses)
			}
		})
	}
}

func TestMovieUseCase_Detail(t *testing.T) {
	// Mock data
	mockMovie := entity.Movie{
		ID:          1,
		Title:       "Test Movie",
		Description: "Test Description",
		Rating:      4.5,
		Image:       "test.jpg",
	}

	// Setup
	mockDB, dbmock := DbMock(t)
	dbmock.MatchExpectationsInOrder(false)
	dbmock.ExpectQuery("SELECT (.+) FROM movies WHERE id = ?").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "rating", "image"}).
			AddRow(mockMovie.ID, mockMovie.Title, mockMovie.Description, mockMovie.Rating, mockMovie.Image))

	mockLog := log.New(logrus.StandardLogger(), "app-unit-test")
	mockValidate := validator.New()
	mockMovieRepo := repository.NewMovieRepository()

	uc := NewMovieUseCase(mockDB, mockLog, mockValidate, mockMovieRepo)

	// Test cases
	testCases := []struct {
		name           string
		id             int
		expectedError  error
		expectedOutput *model.MovieResponse
	}{
		{
			name:           "ValidRequest",
			id:             1,
			expectedError:  nil,
			expectedOutput: &model.MovieResponse{ID: 1, Title: "Test Movie", Description: "Test Description", Rating: 4.5, Image: "test.jpg", CreatedAt: "01-01-0001", UpdatedAt: "01-01-0001"},
		},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Invoke the method under test
			response, err := uc.Detail(context.Background(), tc.id)

			// Assertions
			if tc.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedOutput, response)
			}
		})
	}
}
