package repo

import (
	"github.com/datnguyennnx/go-23/ex6/pkg/repo/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type pgRepo struct {
	db       *gorm.DB
	userRepo user.UserRepo
}

func NewPGRepo(connectionStr string) UserRepo {
	db, err := gorm.Open(postgres.Open(connectionStr), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &pgRepo{
		db:       db,
		userRepo: user.NewUserRepo(db),
	}
}

func (r pgRepo) AutoMigrate(models ...interface{}) error {
	for idx := range models {
		if err := r.db.AutoMigrate(models[idx]); err != nil {
			return err
		}
	}
	return nil
}

func (r pgRepo) User() user.UserRepo {
	return r.userRepo
}

func (r pgRepo) DB() *gorm.DB {
	return r.db
}
