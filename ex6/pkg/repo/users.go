package repo

import (
	"github.com/datnguyennnx/go-23/ex6/pkg/repo/user"
	"gorm.io/gorm"
)

// Repo is the interface that wraps the basic methods for a repository.
type UserRepo interface {
	DB() *gorm.DB
	AutoMigrate(models ...interface{}) error
	User() user.UserRepo
}
