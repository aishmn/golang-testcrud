package main

import (
	"log"
	"net/http"

	"github.com/aishmn/golang-testcrud/controllers"
	"github.com/aishmn/golang-testcrud/initializers"
	"github.com/aishmn/golang-testcrud/routes"
	"github.com/gin-gonic/gin"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	PostController      controllers.PostController
	PostRouteController routes.PostRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Couldnot load env variables", err)
	}
	initializers.ConnectDB(&config)
	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	PostController = controllers.NewPostController(initializers.DB)
	PostRouteController = routes.NewRoutePostController(PostController)

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

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	PostRouteController.PostRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
