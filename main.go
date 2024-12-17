package main

import (
	"database/sql"
	"log"
	"os"
	app "ux/app"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Autorise React
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/products", app.ProductPage)
	r.Run(":8080")
}

func init() {
	var err error
	filestat, _ := os.Stat("db.sql")
	if filestat != nil {
		log.Println("Database exists")
		app.Db, err = sql.Open("sqlite3", "./db.sql")
		if err != nil {
			panic(err)
		}
	} else {
		log.Println("Database not exists")
		os.Create("db.sql")
		app.Db, err = sql.Open("sqlite3", "./db.sql")
		if err != nil {
			panic(err)
		}
		app.CreateDb()
		app.FakerProducts()
	}
}
