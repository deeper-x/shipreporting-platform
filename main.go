package main

import (
	"github.com/deeper-x/shipreporting-platform/webserver"
)

var inst = webserver.Instance{}

func main() {
	webserver.Run(&inst)

}
