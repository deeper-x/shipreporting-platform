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
	Run() error
}

// Instance todo doc
type Instance struct {
	Engine *gin.Engine
}

// EngineBuild todo doc
func (i *Instance) EngineBuild() error {
	r := gin.Default()

	r.Static("/assets", utils.StaticPath)

	//templates architecture mapping
	r.HTMLRender = multiRenderer()

	i.Engine = r

	return nil
}

// Run todo doc - Test in Unit testing
func (i *Instance) Run() error {
	err := i.Engine.Run()

	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Serve todo doc
func Serve(s Server) {
	s.EngineBuild()
	s.Route()
	s.Run()
}
