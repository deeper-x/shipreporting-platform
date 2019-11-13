package webserver

// Route todo doc - Test delegate to Integration Testing
func (i *Instance) Route() {
	i.Engine.GET("/mooring_now", MooringNow)
	i.Engine.GET("/anchored_now", AnchoredNow)
	i.Engine.GET("/arrivals_today", ArrivalsToday)
	i.Engine.GET("/departures_today", DeparturesToday)
	i.Engine.GET("/arrival_previsions_now", ArrivalPrevisionsNow)
	i.Engine.GET("/shipped_goods_today", ShippedGoodsToday)
	i.Engine.GET("/traffic_list_today", TrafficListToday)
	i.Engine.GET("/shifting_previsions_now", ShiftingPrevisionsNow)
	i.Engine.GET("/departure_previsions_now", DeparturePrevisionsNow)
}
