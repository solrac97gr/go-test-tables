package repositories

import "github.com/solrac97gr/go-test-tables/users/domain/models"

type FakeStorage struct {
	DB map[int]*models.User
}

func NewFakeStorage() *FakeStorage {
	return &FakeStorage{
		DB: make(map[int]*models.User),
	}
}

var (
	ErrSavingUser = models.ErrSavingUser
)

func (s *FakeStorage) SaveUser(user *models.User) error {
	s.DB[user.ID] = user
	return nil
}
