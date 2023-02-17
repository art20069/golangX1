package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func serveRoutes(r *gin.Engine) {

	r.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"greeting": "Hello Server Default"})
	})

	r.GET("/products/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"greeting": "Hello Server Default"})
	})

	r.PATCH("/heproductsllo/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"greeting": "Hello Server Default"})
	})

	r.DELETE("/products", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"greeting": "Hello Server Default"})
	})

}
