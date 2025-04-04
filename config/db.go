package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var err error

func DatbaseInit() {
	host := "localhost"
	user := "root"
	password := "secret"
	dbName := "go-rest-api"
	port := 5432

	//creating the dsn
	dsn := fmt.Sprintf("host=%s,user=%s,password=%s,db-name=%s,port=%s", host, user, password, dbName, port)

	//connecting to the database
	database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}
func DB() *gorm.DB {
	return database
}
