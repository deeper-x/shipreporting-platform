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

	r.GET("/mooring_now", func(x *gin.Context) {
		x.HTML(http.StatusOK, "mooring_now", gin.H{
			"shipName": "Eurocargo XYZ",
		})
	})

	r.Run()
}

func createMultiple() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("mooring_now", "templates/layout/theme/classic/default/index.html", "templates/tables/mooring_now.html")

	return r
}
