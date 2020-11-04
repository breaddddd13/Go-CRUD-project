package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Student :define format
type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Major  string `json:"major"`
	School string `json:"school"`
}

var router *gin.Engine
var db *gorm.DB
var err error

func main() {
	dsn := "host=localhost user=sonicwang password=0000 dbname=local port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	db.AutoMigrate(&Student{})

	router = gin.Default()
	initializeRoutes()
	router.Run()
}
