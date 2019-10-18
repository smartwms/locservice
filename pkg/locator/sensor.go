package locator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/golang/geo/r2"
)

type sensor struct {
	ID                string   `json:"id"`
	Position          r2.Point `json:"position"`
	RefPower          float64  `json:"ref_power"`
	AttenuationFactor float64  `json:"attenuation_factor"`
}

func loadSensors() (map[string]sensor, error) {
	f, err := ioutil.ReadFile("data/sensors.json")

	if err != nil {
		return nil, fmt.Errorf("Couldn't read sensors file: %s", err)
	}

	raw := struct {
		Sensors map[string]sensor
	}{}

	if err := json.Unmarshal(f, &raw); err != nil {
		return nil, fmt.Errorf("Couldn't unmarshal file", err)
	}

	return raw.Sensors, nil
}
