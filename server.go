package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"isithalloween.com/server"
)

func main() {
	server.KeyboardInterruptHandler() // clean exit handler
	err := godotenv.Load(".env")      // check for .env file
	var PORT string                   // port
	if err != nil {
		PORT = "5500" // .env not found
	} else {
		PORT = os.Getenv("PORT")
	}
	// initialize gin router
	router := gin.Default()
	// host the frontend
	router.Static("/", "./client")
	// concurrently listen for requests
	go receive(router)
	// run server
	router.Run("localhost:" + PORT)
}

func receive(router *gin.Engine) {
	router.POST("/check", server.HalloweenCheck)
}
