package db

import (
	"fmt"
	"log"

	configpkg "github.com/dkostenko/gin-server-example/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB gorm.DB

func init() {
	dbConfig := fmt.Sprintf("user=%s dbname=%s sslmode=disable", configpkg.DbUser, configpkg.DbName)

	db, err := gorm.Open("postgres", dbConfig)

	if err != nil {
		log.Printf("type: %T, value: %v, readable value: %#v", err, err, err)
		panic(err)
	}

	DB = db
}
