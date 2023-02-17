package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	os.MkdirAll("uploads/products", 0755)
	r := gin.Default()
	serveRoutes(r)
	r.Run()

}
