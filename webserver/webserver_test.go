package webserver

import (
	"testing"
)

var inst = Instance{}

// Most of webserver testing is delegated to integration testing, so %coverage here will not be huge
func TestEngineBuild(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}
}
