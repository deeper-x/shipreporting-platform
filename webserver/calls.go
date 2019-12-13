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
	// Step 1: Ask for token to auth webservice
	status, token := ProcessAuth(c)

	// Step 2: Auth is granted?
	if status {
		// Step 2A: try to build the session (w/ data)
		created := CreateSession(c, token)

		// Step 2B: Verify session is built successfully
		if !created {
			// Step 2B-1: Something get wrong, bounce to login page
			Login(c)
		} else {
			// Step 2B-2: Success. Go to welcome page
			ActiveNow(c)
		}

	} else {
		// Step 2C: remote WS denied access. Bounce to login
		Login(c)
	}
}

// Welcome render a landing/access page
var Welcome = func(c *gin.Context) {
	c.HTML(http.StatusOK, "welcome", gin.H{
		"message": "welcome",
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

// FormMooredRegister search form
var FormMooredRegister = func(c *gin.Context) {
	conn := db.Connector()
	repo := db.NewRepository(conn)
	defer repo.Close()

	session := sessions.Default(c)
	portinformer := session.Get("managedPortinformer")

	portinformerName := repo.SelectPortinformerName(portinformer)

	c.HTML(http.StatusOK, "form_moored_register", gin.H{
		"SHIPFLOW_SERVER":      os.Getenv("SHIPFLOW_SERVER"),
		"SHIPREPORTING_SERVER": os.Getenv("SHIPREPORTING_SERVER"),
		"portinformer":         portinformer,
		"pageName":             "Moored register",
		"portinformerName":     portinformerName,
	})
}

// MooredRegister call for register arrivals
var MooredRegister = func(c *gin.Context) {
	conn := db.Connector()
	repo := db.NewRepository(conn)
	defer repo.Close()

	session := sessions.Default(c)
	portinformer := session.Get("managedPortinformer")

	portinformerName := repo.SelectPortinformerName(portinformer)
	tsStart := c.Param("ts_start")
	tsStop := c.Param("ts_stop")

	c.HTML(http.StatusOK, "moored_register", gin.H{
		"SHIPFLOW_SERVER":  os.Getenv("SHIPFLOW_SERVER"),
		"portinformer":     portinformer,
		"pageName":         "Moored register",
		"portinformerName": portinformerName,
		"ts_start":         tsStart,
		"ts_stop":          tsStop,
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

// FormArrivalsRegister search form
var FormArrivalsRegister = func(c *gin.Context) {
	conn := db.Connector()
	repo := db.NewRepository(conn)
	defer repo.Close()

	session := sessions.Default(c)
	portinformer := session.Get("managedPortinformer")

	portinformerName := repo.SelectPortinformerName(portinformer)

	c.HTML(http.StatusOK, "form_arrivals_register", gin.H{
		"SHIPFLOW_SERVER":      os.Getenv("SHIPFLOW_SERVER"),
		"SHIPREPORTING_SERVER": os.Getenv("SHIPREPORTING_SERVER"),
		"portinformer":         portinformer,
		"pageName":             "Arrivals register",
		"portinformerName":     portinformerName,
	})
}

// ArrivalsRegister call for register arrivals
var ArrivalsRegister = func(c *gin.Context) {
	conn := db.Connector()
	repo := db.NewRepository(conn)
	defer repo.Close()

	session := sessions.Default(c)
	portinformer := session.Get("managedPortinformer")

	portinformerName := repo.SelectPortinformerName(portinformer)
	tsStart := c.Param("ts_start")
	tsStop := c.Param("ts_stop")

	c.HTML(http.StatusOK, "arrivals_register", gin.H{
		"SHIPFLOW_SERVER":  os.Getenv("SHIPFLOW_SERVER"),
		"portinformer":     portinformer,
		"pageName":         "Arrivals register",
		"portinformerName": portinformerName,
		"ts_start":         tsStart,
		"ts_stop":          tsStop,
	})
}

// FormDeparturesRegister search form
var FormDeparturesRegister = func(c *gin.Context) {
	conn := db.Connector()
	repo := db.NewRepository(conn)
	defer repo.Close()

	session := sessions.Default(c)
	portinformer := session.Get("managedPortinformer")

	portinformerName := repo.SelectPortinformerName(portinformer)

	c.HTML(http.StatusOK, "form_departures_register", gin.H{
		"SHIPFLOW_SERVER":      os.Getenv("SHIPFLOW_SERVER"),
		"SHIPREPORTING_SERVER": os.Getenv("SHIPREPORTING_SERVER"),
		"portinformer":         portinformer,
		"pageName":             "Departures register",
		"portinformerName":     portinformerName,
	})
}

// DeparturesRegister call for register arrivals
var DeparturesRegister = func(c *gin.Context) {
	conn := db.Connector()
	repo := db.NewRepository(conn)
	defer repo.Close()

	session := sessions.Default(c)
	portinformer := session.Get("managedPortinformer")

	portinformerName := repo.SelectPortinformerName(portinformer)
	tsStart := c.Param("ts_start")
	tsStop := c.Param("ts_stop")

	c.HTML(http.StatusOK, "departures_register", gin.H{
		"SHIPFLOW_SERVER":  os.Getenv("SHIPFLOW_SERVER"),
		"portinformer":     portinformer,
		"pageName":         "Departures register",
		"portinformerName": portinformerName,
		"ts_start":         tsStart,
		"ts_stop":          tsStop,
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
