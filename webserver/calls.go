package webserver

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/deeper-x/shipreporting-platform/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(utils.DotenvFile)
	if err != nil {
		log.Println(err)
	}
}

// Login form presentation
var Login = func(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{
		"data": "ok",
	})
}

// Enter present auth form
var Enter = func(c *gin.Context) {
	status, _ := ProcessAuth(c)

	if status == http.StatusOK {
		created := CreateSession(c)

		if !created {
			Login(c)
		} else {
			Welcome(c)
		}

	} else {
		Login(c)
	}
}

// Welcome render a landing page
var Welcome = func(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome", gin.H{
		"data": "ok",
	})
}

// Logout destroy user session
var Logout = func(c *gin.Context) {
	status, message := DestroySession(c)

	c.HTML(status, "login", gin.H{
		"data": message,
	})
}

// MooringNow call for currently moored
var MooringNow = func(c *gin.Context) {
	session := sessions.Default(c)

	fmt.Println(session.Get(userkey))
	c.HTML(http.StatusOK, "mooring_now", gin.H{
		"SHIPFLOW_SERVER": os.Getenv("SHIPFLOW_SERVER"),
	})
}

// AnchoredNow call for currently anchored
var AnchoredNow = func(c *gin.Context) {
	c.HTML(http.StatusOK, "anchored_now", gin.H{
		"SHIPFLOW_SERVER": os.Getenv("SHIPFLOW_SERVER"),
	})
}

// ArrivalsToday call for arrivals today
var ArrivalsToday = func(c *gin.Context) {
	c.HTML(http.StatusOK, "arrivals_today", gin.H{
		"SHIPFLOW_SERVER": os.Getenv("SHIPFLOW_SERVER"),
	})
}

// DeparturesToday call for departures today
var DeparturesToday = func(c *gin.Context) {
	c.HTML(http.StatusOK, "departures_today", gin.H{
		"SHIPFLOW_SERVER": os.Getenv("SHIPFLOW_SERVER"),
	})
}

// ArrivalPrevisionsNow call for arrival previsions active
var ArrivalPrevisionsNow = func(c *gin.Context) {
	c.HTML(http.StatusOK, "arrival_previsions_now", gin.H{
		"SHIPFLOW_SERVER": os.Getenv("SHIPFLOW_SERVER"),
	})
}

// ShippedGoodsToday todo doc
var ShippedGoodsToday = func(c *gin.Context) {
	c.HTML(http.StatusOK, "shipped_goods_today", gin.H{
		"SHIPFLOW_SERVER": os.Getenv("SHIPFLOW_SERVER"),
	})
}

// TrafficListToday call for today ro/ro pax
var TrafficListToday = func(c *gin.Context) {
	c.HTML(http.StatusOK, "traffic_list_today", gin.H{
		"SHIPFLOW_SERVER": os.Getenv("SHIPFLOW_SERVER"),
	})
}

// ShiftingPrevisionsNow call for arrival previsions active
var ShiftingPrevisionsNow = func(c *gin.Context) {
	c.HTML(http.StatusOK, "shifting_previsions_now", gin.H{
		"SHIPFLOW_SERVER": os.Getenv("SHIPFLOW_SERVER"),
	})
}

// DeparturePrevisionsNow call for arrival previsions active
var DeparturePrevisionsNow = func(c *gin.Context) {
	c.HTML(http.StatusOK, "departure_previsions_now", gin.H{
		"SHIPFLOW_SERVER": os.Getenv("SHIPFLOW_SERVER"),
	})
}
