package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	
)

var (
	db * gorm.DB
)


// func GetDBConObj(config config.Config) *gorm.DB {
// 	dbHost := config.GetString(`database.host`)
// 	dbPort := config.GetString(`database.port`)
// 	dbUser := config.GetString(`database.user`)
// 	dbPass := config.GetString(`database.pass`)
// 	dbName := config.GetString(`database.name`)
// 	db, err := gorm.Open("postgres",
// 		"host="+dbHost+" port="+dbPort+" user="+dbUser+" dbname="+dbName+" password="+dbPass+" sslmode=disable")
// 	if err != nil {
// 		println("Database connection established failed ==>", err.Error())
// 	} else {
// 		println("Database connection established successfully")
// 	}
// 	//return db
// }

func Connect() {
	// Please define your user name and password for my sql.

	dsn := "host=localhost port=5432 user=postgres dbname=test1 sslmode=disable password=nineleaps"
	d, err := gorm.Open("postgres", dsn)

	// d, err := gorm.Open("mysql", "root:root@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	 db = d
}

func GetDB() *gorm.DB {
	return db
}