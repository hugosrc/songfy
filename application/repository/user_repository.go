package repository

import "github.com/hugosrc/songfy/domain"

// UserRepository repository interface
type UserRepository interface {
	Save(user *domain.User) (*domain.User, error)
}
