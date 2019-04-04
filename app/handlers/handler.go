package handlers

import (
	"github.com/chadiso/golang_tryout/app/models"
	"github.com/jinzhu/gorm"
)

type (
	Context struct {
		DB *gorm.DB
	}
)

func (hc Context) RunMigrations() {
	hc.DB.AutoMigrate(&models.Transaction{})
}
