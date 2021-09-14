package main

import (
	//built-in
	"net/http"
	"os"

	//external or not built-in
	"github.com/Gurv33r/Go-Gin-RESTfulAPI/server"
	"github.com/gin-gonic/gin"
)

func main() {
	// for linux/mac: use curl http://localhost:port/dir
	// for windows: use Invoke-WebRequest -URI http://localhost:port/dir
	// web browsers will work too, just go to http://localhost:port/dir

	// Basics of Go Gin:
	// all APIs, especially web APIs need a router
	// instantiate one with gin.Default()
	router := gin.Default()
	// A Gin router is special, where you can specify handlers to HTTP requests to a specfic subdirectory
	// this is useful because Gin abstracts over the routing issues for you and requires you to code the bare essentials of a RESTful API: create a handler and route a request to a handler
	// One way is to make immediate and specific routes to handler for each request with the Gin.Router's GET(), POST(), and PUT() methods
	// for now since the code only transfers data, set getAssignments to GET request handler for the /hw subdirectory
	router.GET("/hw", getAssignments)
	// so far, all this is is letting someone else access an API endpoint for the hw subdirectory
	// let's add a home endpoint by adding a home page handler and then linking it to the '/' sub directory for GET requests
	router.GET("/", homePagehandler)
	// now, just http://localhost:port will call the home page handler

	// start the router with router.Run("localhost:port")
	// grab the port number from the env file
	// os package has Getenv method that works well
	router.Run("localhost:" + os.Getenv("PORT"))
}

// accessor for assignments slice -> returns assignments slice as a single JSON string
// Gin framework's cornerstone is gin.Context, which is an abstraction over a datastruct that carries HTTP request data and also converts to and from JSON format
// to convert a struct slice to a JSON string, use gin.Context.IndentedJSON, which takes a http.StatusCode and the Struct slice to convert
// to deny requests, set the status code to http.StatusBadRequest and the struct slice to nil
func getAssignments(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, server.Assignments)
}

// accessor for homepage (http://localhost:port/)
// returns "hello, world"
func homePagehandler(ctx *gin.Context) {
	//sending a string gets handled by gin
	ctx.JSON(http.StatusOK, "Hello, World")
}
