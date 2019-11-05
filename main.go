package main

import (
	"github.com/shipreporting-platform/webserver"
)

var inst = webserver.Instance{}

func main() {
	inst.EngineBuild()
	inst.Route()
	inst.Run()
}
