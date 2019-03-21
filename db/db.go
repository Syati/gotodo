package db

import (
	"fmt"
	"github.com/Syati/gotodo/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Opening a database and save it to `DB`.
func GetDB() *gorm.DB {
	c := config.GetConfig()
	db, err := gorm.Open("mysql", c.Get("db.dsn"))
	if err != nil {
		panic(fmt.Errorf("Fatal db connection error: %s \n", err))
	}

	// some settings
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	return db
}
