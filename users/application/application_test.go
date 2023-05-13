package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/solrac97gr/go-test-tables/mocks"
	"github.com/solrac97gr/go-test-tables/users/application"
	"github.com/solrac97gr/go-test-tables/users/domain/models"
	"github.com/solrac97gr/go-test-tables/users/infrastructure/repositories"
	"github.com/stretchr/testify/assert"
)

func TestApplication_CreateUser(t *testing.T) {
	cases := map[string]struct {
		email       string
		password    string
		testSetup   func(*mocks.MockRepository, *mocks.MockValidator)
		assertSetup func(*testing.T, *models.User, string, string, error)
	}{
		"Empty Password [Validation Error]": {
			email:    "mail@car.com",
			password: "",
			testSetup: func(repo *mocks.MockRepository, val *mocks.MockValidator) {
				val.EXPECT().Struct(gomock.Any()).Return(models.ErrInvalidPassword)
			},
			assertSetup: func(t *testing.T, user *models.User, email, password string, err error) {
				assert.Nil(t, user)
				assert.EqualError(t, err, models.ErrInvalidPassword.Error())
			},
		},
		"Empty Email [Validation Error]": {
			email:    "",
			password: "123456",
			testSetup: func(repo *mocks.MockRepository, val *mocks.MockValidator) {
				val.EXPECT().Struct(gomock.Any()).Return(models.ErrInvalidEmail)
			},
			assertSetup: func(t *testing.T, user *models.User, email, password string, err error) {
				assert.Nil(t, user)
				assert.EqualError(t, err, models.ErrInvalidEmail.Error())
			},
		},

		"Error saving [Repository Error]": {
			email: "test@mail.com",
			testSetup: func(repo *mocks.MockRepository, val *mocks.MockValidator) {
				val.EXPECT().Struct(gomock.Any()).Return(nil)
				repo.EXPECT().SaveUser(gomock.Any()).Return(repositories.ErrSavingUser)
			},
			assertSetup: func(t *testing.T, user *models.User, email, password string, err error) {
				assert.Nil(t, user)
				assert.EqualError(t, err, repositories.ErrSavingUser.Error())
			},
		},

		"Valid User [Success]": {
			email:    "test@mail.com",
			password: "123456",
			testSetup: func(repo *mocks.MockRepository, val *mocks.MockValidator) {
				val.EXPECT().Struct(gomock.Any()).Return(nil)
				repo.EXPECT().SaveUser(gomock.Any()).Return(nil)
			},
			assertSetup: func(t *testing.T, user *models.User, email, password string, err error) {
				assert.NotNil(t, user)
				assert.Equal(t, email, user.Email)
				assert.Equal(t, password, user.Password)
				assert.NoError(t, err)
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			// Create a mock controller
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create a mock repository
			repo := mocks.NewMockRepository(ctrl)
			val := mocks.NewMockValidator(ctrl)

			// Setup the mock repository
			if tc.testSetup != nil {
				tc.testSetup(repo, val)
			}

			app := application.NewUserApp(repo, val)
			user, err := app.CreateUser(tc.email, tc.password)

			// Assert the result
			if tc.assertSetup != nil {
				tc.assertSetup(t, user, tc.email, tc.password, err)
			}
		})
	}
}
