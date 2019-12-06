package webserver

import (
	"log"
	"net/http"
	"os"

	"github.com/deeper-x/shipreporting-platform/db"
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
		"loginMessage": "Welcome",
	})
}

// Enter present auth form
var Enter = func(c *gin.Context) {
	status, token := ProcessAuth(c)

	if status {
		created := CreateSession(c, token)

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
		"loginMessage": "Please insert user&pass",
	})
}

// Logout destroy user session
var Logout = func(c *gin.Context) {
	status, message := DestroySession(c)

	c.HTML(status, "logout", gin.H{
		"data": message,
	})
}

// MooringNow call for currently moored
var MooringNow = func(c *gin.Context) {
	conn := db.Connector()
	repo := db.NewRepository(conn)
	defer repo.Close()

	session := sessions.Default(c)
	portinformer := session.Get("managedPortinformer")

	portinformerName := repo.SelectPortinformerName(portinformer)

	c.HTML(http.StatusOK, "mooring_now", gin.H{
		"SHIPFLOW_SERVER":  os.Getenv("SHIPFLOW_SERVER"),
		"portinformer":     portinformer,
		"pageName":         "Currently moored",
		"portinformerName": portinformerName,
	})
}

// AnchoredNow call for currently anchored
var AnchoredNow = func(c *gin.Context) {
	conn := db.Connector()
	repo := db.NewRepository(conn)
	defer repo.Close()

	session := sessions.Default(c)
	portinformer := session.Get("managedPortinformer")

	portinformerName := repo.SelectPortinformerName(portinformer)

	c.HTML(http.StatusOK, "anchored_now", gin.H{
		"SHIPFLOW_SERVER":  os.Getenv("SHIPFLOW_SERVER"),
		"portinformer":     portinformer,
		"pageName":         "Currently anchored",
		"portinformerName": portinformerName,
	})
}

// ActiveNow currently open trips
var ActiveNow = func(c *gin.Context) {
	conn := db.Connector()
	repo := db.NewRepository(conn)
	defer repo.Close()

	session := sessions.Default(c)
	portinformer := session.Get("managedPortinformer")
	portinformerName := repo.SelectPortinformerName(portinformer)

	c.HTML(http.StatusOK, "active_now", gin.H{
		"SHIPFLOW_SERVER":  os.Getenv("SHIPFLOW_SERVER"),
		"portinformer":     portinformer,
		"pageName":         "Open trips",
		"portinformerName": portinformerName,
	})
}

// ArrivalsToday call for arrivals today
var ArrivalsToday = func(c *gin.Context) {
	conn := db.Connector()
	repo := db.NewRepository(conn)
	defer repo.Close()

	session := sessions.Default(c)
	portinformer := session.Get("managedPortinformer")

	portinformerName := repo.SelectPortinformerName(portinformer)

	c.HTML(http.StatusOK, "arrivals_today", gin.H{
		"SHIPFLOW_SERVER":  os.Getenv("SHIPFLOW_SERVER"),
		"portinformer":     portinformer,
		"pageName":         "Arrivals, today",
		"portinformerName": portinformerName,
	})
}

// DeparturesToday call for departures today
var DeparturesToday = func(c *gin.Context) {
	conn := db.Connector()
	repo := db.NewRepository(conn)
	defer repo.Close()

	session := sessions.Default(c)
	portinformer := session.Get("managedPortinformer")

	portinformerName := repo.SelectPortinformerName(portinformer)

	c.HTML(http.StatusOK, "departures_today", gin.H{
		"SHIPFLOW_SERVER":  os.Getenv("SHIPFLOW_SERVER"),
		"portinformer":     portinformer,
		"pageName":         "Departures, today",
		"portinformerName": portinformerName,
	})
}

// ArrivalPrevisionsNow call for arrival previsions active
var ArrivalPrevisionsNow = func(c *gin.Context) {
	conn := db.Connector()
	repo := db.NewRepository(conn)
	defer repo.Close()

	session := sessions.Default(c)
	portinformer := session.Get("managedPortinformer")

	portinformerName := repo.SelectPortinformerName(portinformer)

	c.HTML(http.StatusOK, "arrival_previsions_now", gin.H{
		"SHIPFLOW_SERVER":  os.Getenv("SHIPFLOW_SERVER"),
		"portinformer":     portinformer,
		"pageName":         "Arrival previsions",
		"portinformerName": portinformerName,
	})
}

// ShippedGoodsToday todo doc
var ShippedGoodsToday = func(c *gin.Context) {
	conn := db.Connector()
	repo := db.NewRepository(conn)
	defer repo.Close()

	session := sessions.Default(c)
	portinformer := session.Get("managedPortinformer")

	portinformerName := repo.SelectPortinformerName(portinformer)

	c.HTML(http.StatusOK, "shipped_goods_today", gin.H{
		"SHIPFLOW_SERVER":  os.Getenv("SHIPFLOW_SERVER"),
		"portinformer":     portinformer,
		"pageName":         "Shipped goods",
		"portinformerName": portinformerName,
	})
}

// TrafficListToday call for today ro/ro pax
var TrafficListToday = func(c *gin.Context) {
	conn := db.Connector()
	repo := db.NewRepository(conn)
	defer repo.Close()

	session := sessions.Default(c)
	portinformer := session.Get("managedPortinformer")

	portinformerName := repo.SelectPortinformerName(portinformer)

	c.HTML(http.StatusOK, "traffic_list_today", gin.H{
		"SHIPFLOW_SERVER":  os.Getenv("SHIPFLOW_SERVER"),
		"portinformer":     portinformer,
		"pageName":         "Traffic list",
		"portinformerName": portinformerName,
	})
}

// ShiftingPrevisionsNow call for arrival previsions active
var ShiftingPrevisionsNow = func(c *gin.Context) {
	conn := db.Connector()
	repo := db.NewRepository(conn)
	defer repo.Close()

	session := sessions.Default(c)
	portinformer := session.Get("managedPortinformer")

	portinformerName := repo.SelectPortinformerName(portinformer)

	c.HTML(http.StatusOK, "shifting_previsions_now", gin.H{
		"SHIPFLOW_SERVER":  os.Getenv("SHIPFLOW_SERVER"),
		"portinformer":     portinformer,
		"pageName":         "Shifting previsions",
		"portinformerName": portinformerName,
	})
}

// DeparturePrevisionsNow call for arrival previsions active
var DeparturePrevisionsNow = func(c *gin.Context) {
	conn := db.Connector()
	repo := db.NewRepository(conn)
	defer repo.Close()

	session := sessions.Default(c)
	portinformer := session.Get("managedPortinformer")

	portinformerName := repo.SelectPortinformerName(portinformer)

	c.HTML(http.StatusOK, "departure_previsions_now", gin.H{
		"SHIPFLOW_SERVER":  os.Getenv("SHIPFLOW_SERVER"),
		"portinformer":     portinformer,
		"pageName":         "Departure previsions",
		"portinformerName": portinformerName,
	})
}
