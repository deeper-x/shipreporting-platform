package webserver

// Route todo doc - Test delegate to Integration Testing
func (i *Instance) Route() {
	i.Engine.GET("/mooring_now", MooringNow)
}
