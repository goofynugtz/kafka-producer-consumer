package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	routes "github.com/goofynugtz/kafka-producer-consumer/pkg/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	router := gin.Default()
	router.Use(cors.Default())

	public := router.Group("/")
	public.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	routes.PublicRoutes(public)
	router.Run(":" + port)
}
