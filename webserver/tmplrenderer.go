package webserver

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/shipreporting-platform/utils"
)

var indexPath = utils.FullPath("/templates/layout/theme/classic/default/index.html")
var mooringPath = utils.FullPath("/templates/tables/mooring_now.html")

// multiRenderer todo doc
func multiRenderer() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	r.AddFromFiles("mooring_now", indexPath, mooringPath)

	return r
}
