package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"isithalloween.com/server"
)

func main() {
	server.KeyboardInterruptHandler()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(".env file does not exist")
	}
	PORT := os.Getenv("PORT")

	router := gin.Default()

	router.Static("/", "./client")

	go receive(router)

	router.Run("localhost:" + PORT)
}

func receive(router *gin.Engine) {
	router.POST("/check", server.HalloweenCheck)
}
