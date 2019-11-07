package webserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shipreporting-platform/utils"
)

// Server todo doc
type Server interface {
	EngineBuild() error
	Route()
	Serve() error
}

// Instance todo doc
type Instance struct {
	Engine *gin.Engine
}

// EngineBuild todo doc
func (i *Instance) EngineBuild() error {
	r := gin.Default()

	r.Static("/assets", utils.TemplateAssets)
	r.Static("/grid", utils.GridAssets)

	//templates architecture mapping
	_, r.HTMLRender = multiRenderer()

	i.Engine = r

	return nil
}

// Serve todo doc - Test in Unit testing
func (i *Instance) Serve() error {
	err := i.Engine.Run()

	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Run todo doc
func Run(s Server) {
	s.EngineBuild()
	s.Route()
	s.Serve()
}
