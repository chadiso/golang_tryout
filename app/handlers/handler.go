package handlers

import "github.com/jinzhu/gorm"

type (
	Context struct {
		DB *gorm.DB
	}
)