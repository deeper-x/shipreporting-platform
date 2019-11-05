package webserver

import (
	"github.com/gin-contrib/multitemplate"
)

// multiRenderer todo doc
func multiRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("mooring_now", IndexPath, MooringPath)

	return r
}
