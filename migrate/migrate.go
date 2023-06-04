package main

import (
	"fmt"
	"log"

	"github.com/aishmn/golang-testcrud/initializers"
	"github.com/aishmn/golang-testcrud/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load env file", err)
	}
	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	fmt.Println("? Migration Complete")
}
