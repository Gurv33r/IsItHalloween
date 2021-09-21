package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"isithalloween.com/server"
)

func main() {
	// for linux/mac: use curl http://localhost:port/dir
	// for windows: use Invoke-WebRequest -URI http://localhost:port/dir
	// web browsers will work too, just go to http://localhost:port/dir

	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file does not exist")
	}
	//initialize the port
	PORT := os.Getenv("PORT")
	fmt.Println(PORT)

	// Basics of Go Gin:
	// all APIs, especially web APIs need a router
	// instantiate one with gin.Default()
	// A Gin router is special, where you can specify handlers to HTTP requests to a specfic subdirectory
	// this is useful because Gin abstracts over the routing issues for you and requires you to code the bare essentials of a RESTful API: create a handler and route a request to a handler
	// One way is to make immediate and specific routes to handler for each request with the Gin.Router's GET(), POST(), and PUT() methods
	router := gin.Default()

	// router.StaticFile("/", "./client/index.html")
	router.Static("/", "./client")

	go receive(router)

	// now, just http://localhost:port will call the home page handler
	// start the router with router.Run("localhost:port")
	// grab the port number from the env file
	// os package has Getenv method that works well

	router.Run("localhost:" + PORT)
}

func receive(router *gin.Engine) {
	router.POST("/check", halloweenCheck)
}

// accessor for homepage (http://localhost:port/)
func halloweenCheck(ctx *gin.Context) {
	reqbody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	reqdate := &server.Date{Str: ""}
	if err := json.Unmarshal(reqbody, &reqdate); err != nil {
		log.Fatal(err)
	}
	log.Println("Data =", reqdate.Str)
	monthAndDay := strings.Split(reqdate.Str, "-")[1:]
	res := "No"
	if strings.Compare(monthAndDay[0], "10") == 0 && strings.Compare(monthAndDay[1], "31") == 0 {
		res = "Yes"
	}
	//sending a string gets handled by gin
	ctx.String(http.StatusOK, res)
}
