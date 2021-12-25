package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Auth struct {
	Id    int    `gorm:"column:id; PRIMARY_KEY" json:"id"`
	Email string `gorm:"column:email" json:"email"`
}

func Connect() error {
	dsn := os.Getenv("PG_CONN")
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Auth{})
}
