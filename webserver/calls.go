package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MooringNow call for currently moored
var MooringNow = func(x *gin.Context) {
	x.HTML(http.StatusOK, "mooring_now", gin.H{
		"data": "ok",
	})
}

// AnchoredNow call for currently anchored
var AnchoredNow = func(x *gin.Context) {
	x.HTML(http.StatusOK, "anchored_now", gin.H{
		"data": "ok",
	})
}

// ArrivalsToday call for arrivals today
var ArrivalsToday = func(x *gin.Context) {
	x.HTML(http.StatusOK, "arrivals_today", gin.H{
		"data": "ok",
	})
}

// DeparturesToday call for departures today
var DeparturesToday = func(x *gin.Context) {
	x.HTML(http.StatusOK, "departures_today", gin.H{
		"data": "ok",
	})
}

// ArrivalPrevisionsNow call for arrival previsions active
var ArrivalPrevisionsNow = func(x *gin.Context) {
	x.HTML(http.StatusOK, "arrival_previsions_now", gin.H{
		"data": "ok",
	})
}

// ShippedGoodsToday todo doc
var ShippedGoodsToday = func(x *gin.Context) {
	x.HTML(http.StatusOK, "shipped_goods_today", gin.H{
		"data": "ok",
	})
}

// TrafficListToday call for today ro/ro pax
var TrafficListToday = func(x *gin.Context) {
	x.HTML(http.StatusOK, "traffic_list_today", gin.H{
		"data": "ok",
	})
}

// ShiftingPrevisionsNow call for arrival previsions active
var ShiftingPrevisionsNow = func(x *gin.Context) {
	x.HTML(http.StatusOK, "shifting_previsions_now", gin.H{
		"data": "ok",
	})
}

// DeparturePrevisionsNow call for arrival previsions active
var DeparturePrevisionsNow = func(x *gin.Context) {
	x.HTML(http.StatusOK, "departure_previsions_now", gin.H{
		"data": "ok",
	})
}
