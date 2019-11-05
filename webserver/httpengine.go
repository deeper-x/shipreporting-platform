package webserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shipreporting-platform/utils"
)

// Server todo doc
type Server interface {
	EngineBuild()
	Route()
	Run()
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
