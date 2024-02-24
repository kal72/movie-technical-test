package usecase

import (
	"context"
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
