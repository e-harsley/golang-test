package database

import (
	"github.com/go-play/api/model"
	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	db.AutoMigrate(model.Country{})
}
