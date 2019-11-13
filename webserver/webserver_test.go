package webserver

import (
	"log"
	"net/http"
	"net/http/httptest"
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

func TestMooringNow(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/arrivals_today", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestAnchoredNow(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/anchored_now", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestArrivalsToday(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/arrivals_today", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestDeparturesToday(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/departures_today", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestArrivalPrevisionsNow(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/arrival_previsions_now", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestShippedGoodsToday(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/shipped_goods_today", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestTrafficListToday(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/traffic_list_today", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestShiftingPrevisionsToday(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/shifting_previsions_now", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestDeparturePrevisionsToday(t *testing.T) {
	err := inst.EngineBuild()

	if err != nil {
		t.Error(err)
	}

	inst.Route()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/departure_previsions_now", nil)

	if err != nil {
		log.Println(err)
	}

	inst.Engine.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
