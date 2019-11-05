package webserver

import (
	"testing"
)

var inst = Instance{}

func TestEngineBuild(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}
}

func TestRoute(t *testing.T) {
	res := inst.Route()
	if !res {
		t.Fail()
	}
}

func TestRun(t *testing.T) {

	// inst.Run()
	t.Fail()
}
