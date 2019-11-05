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

	// static stuff registration
	staticPath := fmt.Sprintf("%v/templates/layout/theme/classic/assets/", utils.ProjectRoot)

	if exists, err := utils.ResExist(staticPath); !exists {
		return err
	}

	r.Static("/assets", staticPath)

	//templates architecture mapping
	r.HTMLRender = multiRenderer()

	i.Engine = r

	return nil
}

// Run todo doc
func (i *Instance) Run() error {
	err := i.Engine.Run()

	if err != nil {
		fmt.Println(err)
	}
	return err
}
