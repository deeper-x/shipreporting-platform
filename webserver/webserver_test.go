package webserver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var inst = Instance{}

// Most of webserver testing is delegated to integration testing, so %coverage here will not be huge
func TestEngineBuild(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}
}

func TestMultiRenderer(t *testing.T) {
	ok, _ := multiRenderer()

	assert.True(t, ok)
}
