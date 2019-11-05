package main

import (
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./templates/layout/theme/classic/assets/")

	r.HTMLRender = createMultiple()

	r.GET("/home", func(x *gin.Context) {
		x.HTML(http.StatusOK, "home", gin.H{
			"ship_name": "Eurocargo XYZ",
		})
	})

	r.Run()
}

func createMultiple() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("home", "templates/layout/theme/classic/default/index.html", "templates/tables/mooring_now.html")

	return r
}
