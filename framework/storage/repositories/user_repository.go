package repositories

import (
	"github.com/hugosrc/songfy/domain"
	"gorm.io/gorm"
)

// UserRepository repository implementation
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Save inserts a new user into the database
func (repo *UserRepository) Save(user *domain.User) (*domain.User, error) {
	err := repo.db.Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
