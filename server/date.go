package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Date struct {
	Str string `json:"currentDate"`
}

func HalloweenCheck(ctx *gin.Context) {
	reqbody, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	reqdate := &Date{Str: ""}
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
