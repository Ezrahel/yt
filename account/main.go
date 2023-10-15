package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server... ")
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		c.JSON(http.StatusOK, gin.H[
			"Hello": "Ajay"
		])
	})
	svc:=&http.Serve{
		addr: "8080",
		Handler: router,
	}
}
