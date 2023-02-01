package database

import (
	"gorm.io/driver/mysql"
    "gorm.io/gorm"
	_ "gorm.io/gorm/logger"
	_ "x-ui/config"
	"x-ui/database/model"
)

var db *gorm.DB

func initUser() error {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		return err
	}
	var count int64
	err = db.Model(&model.User{}).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		user := &model.User{
			Username: "admin",
			Password: "admin",
		}
		return db.Create(user).Error
	}
	return nil
}

func initInbound() error {
	return db.AutoMigrate(&model.Inbound{})
}

func initSetting() error {
	return db.AutoMigrate(&model.Setting{})
}

func InitDB(dbPath string) error {
	dsn := "x_ui_admin:admin@tcp(127.0.0.1:3306)/x_ui"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
    initUser()
    initInbound()
    initSetting()

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func IsNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}
