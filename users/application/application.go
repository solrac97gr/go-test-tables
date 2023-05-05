package application

import (
	"github.com/solrac97gr/go-test-tables/users/domain/models"
	"github.com/solrac97gr/go-test-tables/users/domain/ports"
	val "github.com/solrac97gr/go-test-tables/validator/domain/ports"
)

type UserApp struct {
	UserRepo  ports.Repository
	Validator val.Validator
}

func NewUserApp(repo ports.Repository, val val.Validator) *UserApp {
	return &UserApp{
		UserRepo:  repo,
		Validator: val,
	}
}

func (app *UserApp) CreateUser(email, password string) (*models.User, error) {
	user := &models.User{Email: email, Password: password}

	err := app.Validator.Struct(user)
	if err != nil {
		return nil, err
	}

	err = app.UserRepo.SaveUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
