package webserver

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shipreporting-platform/gswclient"
)

// MooringNow call for currently moored
var MooringNow = func(x *gin.Context) {
	rawData, err := gswclient.FetchData("http://80.88.88.162:8000/moored/27")

	if err != nil {
		log.Println(err)
	}

	jsonData, err := gswclient.JSONConvert(rawData)

	if err != nil {
		log.Println(err)
	}

	x.HTML(http.StatusOK, "mooring_now", gin.H{
		"data": jsonData,
	})
}
