package gswclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// MooredNow struct
type MooredNow struct {
	Agency          string `json:"agency"`
	CurrentActivity string `json:"current_activity"`
	GrossTonnage    string `json:"gross_tonnage"`
	IDTrip          string `json:"id_trip"`
	Iso3            string `json:"iso3"`
	MooringTime     string `json:"mooring_time"`
	Quay            string `json:"quay"`
	Ship            string `json:"ship"`
	ShipType        string `json:"ship_type"`
	ShippedGoods    string `json:"shipped_goods"`
	ShipsLength     string `json:"ships_length"`
	ShipsWidth      string `json:"ships_width"`
	TsEtd           string `json:"ts_etd"`
}

// FetchData get goship data
func FetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// JSONConvert []byte to JSON
func JSONConvert(data []byte) (interface{}, error) {
	// MooredNowData result container
	var MooredNowData []MooredNow

	err := json.Unmarshal(data, &MooredNowData)

	if err != nil {
		return nil, err
	}

	return MooredNowData, nil
}
