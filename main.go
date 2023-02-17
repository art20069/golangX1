package main

import (
	"codezard-pos/db"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loadding file env")
	}
	db.ConnectDB()

	os.MkdirAll("uploads/products", 0755)
	r := gin.Default()
	serveRoutes(r)
	r.Run()
}
