package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type LocateBody struct {
	DeviceID string `json:"device_id"`
	Measures []struct {
		Sensor string `json:"sensor"`
		RSSI   int64  `json:"rssi"`
		AOA    int64  `json:"aoa"`
	} `json:"measures"`
}

func locateHandler(w http.ResponseWriter, r *http.Request) {
	var data LocateBody

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &data)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
