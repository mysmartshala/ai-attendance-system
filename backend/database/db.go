package database

import (
	"fmt"
	"ai-attendance-system/config"
	"ai-attendance-system/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&models.Student{},
		&models.Attendance{},
		&models.Teacher{},
	)

	return db, nil
}
