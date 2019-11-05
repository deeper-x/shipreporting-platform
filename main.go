package main

import (
	"github.com/shipreporting-platform/webserver"
)

var inst = webserver.Instance{}

func main() {
	webserver.Run(&inst)
}
