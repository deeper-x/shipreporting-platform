package webserver

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/shipreporting-platform/utils"
)

var paths = map[string]string{
	"indexPath":   utils.FullPath("/templates/layout/theme/classic/default/index.html"),
	"mooringPath": utils.FullPath("/templates/tables/mooring_now.html"),
}

// multiRenderer todo doc
func multiRenderer() (bool, multitemplate.Renderer) {
	r := multitemplate.NewRenderer()

	// templates must exists
	checker := PathChecker(paths)
	if !checker {
		return false, nil
	}

	// load one by one
	r.AddFromFiles("mooring_now", paths["indexPath"], paths["mooringPath"])

	return true, r
}

// PathChecker check if fs resource exist
func PathChecker(list map[string]string) bool {
	for i := range list {
		if ok, _ := utils.ResExist(list[i]); !ok {
			return false
		}
	}

	return true
}
