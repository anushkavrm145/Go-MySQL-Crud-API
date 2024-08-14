package main

import (
	"fmt"

	// "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urlDSN = "root:Pass%40accolite%243@tcp(localhost:3306)/mydb"
var err error

func DataMigration() {
	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Connection failed")
	}
	Database.AutoMigrate(&Employee{})
}
