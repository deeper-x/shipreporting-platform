package webserver

import (
	"github.com/deeper-x/shipreporting-platform/utils"
	"github.com/gin-contrib/multitemplate"
)

var paths = map[string]string{
	"loginPath":             utils.FullPath("/templates/layout/theme/classic/assets/Login/login.html"),
	"logoutPath":            utils.FullPath("/templates/layout/theme/classic/assets/Login/logout.html"),
	"indexPath":             utils.FullPath("/templates/layout/theme/classic/default/index.html"),
	"welcomePath":           utils.FullPath("/templates/layout/theme/classic/default/welcome.html"),
	"mooringPath":           utils.FullPath("/templates/tables/mooring_now.html"),
	"anchoredPath":          utils.FullPath("/templates/tables/anchored_now.html"),
	"arrivalsTodayPath":     utils.FullPath("/templates/tables/arrivals_today.html"),
	"departuresTodayPath":   utils.FullPath("/templates/tables/departures_today.html"),
	"arrivalPrevNowPath":    utils.FullPath("/templates/tables/arrival_previsions_now.html"),
	"shippedGoodsTodayPath": utils.FullPath("/templates/tables/shipped_goods_today.html"),
	"trafficListTodayPath":  utils.FullPath("/templates/tables/traffic_list_today.html"),
	"shiftingPrevNowPath":   utils.FullPath("/templates/tables/shifting_previsions_now.html"),
	"departurePrevNowPath":  utils.FullPath("/templates/tables/departure_previsions_now.html"),
	"activeNowPath":         utils.FullPath("/templates/tables/active_now.html"),
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
	r.AddFromFiles("login", paths["loginPath"])
	r.AddFromFiles("logout", paths["logoutPath"])
	r.AddFromFiles("welcome", paths["indexPath"], paths["welcomePath"])
	r.AddFromFiles("mooring_now", paths["indexPath"], paths["mooringPath"])
	r.AddFromFiles("anchored_now", paths["indexPath"], paths["anchoredPath"])
	r.AddFromFiles("arrivals_today", paths["indexPath"], paths["arrivalsTodayPath"])
	r.AddFromFiles("departures_today", paths["indexPath"], paths["departuresTodayPath"])
	r.AddFromFiles("arrival_previsions_now", paths["indexPath"], paths["arrivalPrevNowPath"])
	r.AddFromFiles("shipped_goods_today", paths["indexPath"], paths["shippedGoodsTodayPath"])
	r.AddFromFiles("traffic_list_today", paths["indexPath"], paths["trafficListTodayPath"])
	r.AddFromFiles("shifting_previsions_now", paths["indexPath"], paths["shiftingPrevNowPath"])
	r.AddFromFiles("departure_previsions_now", paths["indexPath"], paths["departurePrevNowPath"])
	r.AddFromFiles("active_now", paths["indexPath"], paths["activeNowPath"])
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
