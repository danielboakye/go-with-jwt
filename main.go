package main

import (
	"log"
	"os"

	"github.com/danielboakye/go-with-jwt/middleware"
	routes "github.com/danielboakye/go-with-jwt/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	router := gin.New()

	// set localost only for development
	router.SetTrustedProxies([]string{"localhost:9000"})

	router.Use(gin.Logger())

	router.Use(middleware.ErrorHandler)

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	err := router.Run(":" + port)
	if err != nil {
		log.Panic(err)
	}
}
