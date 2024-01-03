package database

import (
	"fmt"
	"gorm.io/gorm"
	"sync"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB
var once sync.Once

func InitDB(host, user, pass, dbName string, port int) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, user, pass, dbName, port)

	fmt.Println(dsn)
	once.Do(
		func() {

			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				panic(err)
			}

			DB = db
		})
}

func Migrate(models ...any) error {
	return DB.AutoMigrate(models...)
}
