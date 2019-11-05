package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MooringNow call for currently moored
var MooringNow = func(x *gin.Context) {
	x.HTML(http.StatusOK, "mooring_now", gin.H{
		"shipName": "Eurocargo XYZ",
	})
}
