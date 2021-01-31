package storage

import (
	"os"

	"github.com/hugosrc/songfy/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Repositories contain all application repositories
type Repositories struct {
	db *gorm.DB
}

// NewRepositories create an repositories
func NewRepositories(env string) (*Repositories, error) {
	var dsn string

	if env == "test" {
		dsn = os.Getenv("DB_DSN_TEST")
	} else {
		dsn = os.Getenv("DB_DSN_PROD")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Repositories{
		db: db,
	}, nil
}

// Close closes database connection
func (repos *Repositories) Close() error {
	db, err := repos.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

// AutoMigrate migrates models to the database
func (repos *Repositories) AutoMigrate() error {
	return repos.db.AutoMigrate(&domain.User{})
}
