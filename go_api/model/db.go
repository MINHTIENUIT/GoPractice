package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
	dsn := "root:nec@tcp(localhost:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Customer{})

	Db = db
}

func Paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case size > 100:
			size = 100
		}

		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
