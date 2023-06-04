package main

import (
	"log"
	"net/http"

	"github.com/aishmn/golang-testcrud/initializers"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Couldnot load env variables", err)
	}
	initializers.ConnectDB(&config)
	server = gin.Default()

}

func main() {

	config, err := initializers.LoadConfig(".")

	if err != nil {
		log.Fatal("? Couldnot load env variables", err)
	}

	router := server.Group("/api")

	router.GET("/healthcheck", func(ctx *gin.Context) {
		message := "Welcome to CRUD GO LANG"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	log.Fatal(server.Run(":" + config.ServerPort))
}
