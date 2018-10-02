package repositories

import "github.com/jinzhu/gorm"

type UserRepo struct {
	dbCtx *gorm.DB
}

func NewUserRepo(dbCtx *gorm.DB) *UserRepo {
	return &UserRepo{
		dbCtx: dbCtx,
	}
}