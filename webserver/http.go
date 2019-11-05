package webserver

import (
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

// Runner todo doc
type Runner interface {
	Run()
}

// Instance todo doc
type Instance struct {
	engine *gin.Engine
}

// EngineBuild todo doc
func (i Instance) EngineBuild() {
	r := gin.Default()

	r.Static("/assets", "./templates/layout/theme/classic/assets/")

	r.HTMLRender = multiRenderer()

	r.GET("/mooring_now", func(x *gin.Context) {
		x.HTML(http.StatusOK, "mooring_now", gin.H{
			"shipName": "Eurocargo XYZ",
		})
	})

	i.engine = r
}

// Run todo doc
func (i Instance) Run() error {
	err := i.engine.Run()

	return err
}

//
func multiRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("mooring_now", "templates/layout/theme/classic/default/index.html", "templates/tables/mooring_now.html")

	return r
}
