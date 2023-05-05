package ports

import "github.com/solrac97gr/go-test-tables/users/domain/models"

type Application interface {
	// CreateUser creates a new user
	CreateUser(email, password string) (*models.User, error)
}

type Repository interface {
	// SaveUser saves a user
	SaveUser(user *models.User) error
}
