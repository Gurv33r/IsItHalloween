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
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file does not exist")
	}
	PORT := os.Getenv("PORT")
	fmt.Println(PORT)

	router := gin.Default()

	router.Static("/", "./client")

	go receive(router)

	router.Run("localhost:" + PORT)
}

func receive(router *gin.Engine) {
	router.POST("/check", halloweenCheck)
}

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
	ctx.String(http.StatusOK, res)
}
