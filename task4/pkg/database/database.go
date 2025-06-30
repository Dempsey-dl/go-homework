package database

import (
	"blog/internal/config"
	"blog/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.Database.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{}); err != nil {
		return nil, err
	}
	return db, nil
}
