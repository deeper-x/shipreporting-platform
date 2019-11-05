package webserver

// Route todo doc
func (i *Instance) Route() {
	i.Engine.GET("/mooring_now", MooringNow)
}
