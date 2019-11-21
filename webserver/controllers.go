package webserver

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

// Route todo doc - Test delegate to Integration Testing
func (i *Instance) Route() {
	store := cookie.NewStore([]byte("secret"))
	i.Engine.Use(sessions.Sessions(userkey, store))

	i.Engine.GET("/login", Login)
	i.Engine.GET("/logout", Logout)
	i.Engine.POST("/enter", Enter)

	private := i.Engine.Group("/application")
	private.Use(AuthRequired)
	{
		private.GET("/welcome", Welcome)
		private.GET("/mooring_now", MooringNow)
		private.GET("/anchored_now", AnchoredNow)
		private.GET("/arrivals_today", ArrivalsToday)
		private.GET("/departures_today", DeparturesToday)
		private.GET("/arrival_previsions_now", ArrivalPrevisionsNow)
		private.GET("/shipped_goods_today", ShippedGoodsToday)
		private.GET("/traffic_list_today", TrafficListToday)
		private.GET("/shifting_previsions_now", ShiftingPrevisionsNow)
		private.GET("/departure_previsions_now", DeparturePrevisionsNow)
	}

}
