package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login form presentation
var Login = func(c *gin.Context) {
	c.HTML(http.StatusOK, "login", gin.H{
		"data": "ok",
	})
}

// Enter present auth form
var Enter = func(c *gin.Context) {
	status, token := ProcessAuth(c)

	if status == http.StatusOK {
		created := CreateSession(c, token)

		if !created {
			c.Redirect(http.StatusMovedPermanently, "/login")
		} else {
			c.Redirect(http.StatusMovedPermanently, "/application/welcome")
		}

	} else {
		c.Redirect(http.StatusMovedPermanently, "/login")
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
	c.HTML(http.StatusOK, "mooring_now", gin.H{
		"data": "ok",
	})
}

// AnchoredNow call for currently anchored
var AnchoredNow = func(c *gin.Context) {
	c.HTML(http.StatusOK, "anchored_now", gin.H{
		"data": "ok",
	})
}

// ArrivalsToday call for arrivals today
var ArrivalsToday = func(c *gin.Context) {
	c.HTML(http.StatusOK, "arrivals_today", gin.H{
		"data": "ok",
	})
}

// DeparturesToday call for departures today
var DeparturesToday = func(c *gin.Context) {
	c.HTML(http.StatusOK, "departures_today", gin.H{
		"data": "ok",
	})
}

// ArrivalPrevisionsNow call for arrival previsions active
var ArrivalPrevisionsNow = func(c *gin.Context) {
	c.HTML(http.StatusOK, "arrival_previsions_now", gin.H{
		"data": "ok",
	})
}

// ShippedGoodsToday todo doc
var ShippedGoodsToday = func(c *gin.Context) {
	c.HTML(http.StatusOK, "shipped_goods_today", gin.H{
		"data": "ok",
	})
}

// TrafficListToday call for today ro/ro pax
var TrafficListToday = func(c *gin.Context) {
	c.HTML(http.StatusOK, "traffic_list_today", gin.H{
		"data": "ok",
	})
}

// ShiftingPrevisionsNow call for arrival previsions active
var ShiftingPrevisionsNow = func(c *gin.Context) {
	c.HTML(http.StatusOK, "shifting_previsions_now", gin.H{
		"data": "ok",
	})
}

// DeparturePrevisionsNow call for arrival previsions active
var DeparturePrevisionsNow = func(c *gin.Context) {
	c.HTML(http.StatusOK, "departure_previsions_now", gin.H{
		"data": "ok",
	})
}
