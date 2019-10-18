package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/smartwms/locservice/pkg/locator"
)

type server struct {
	router *mux.Router
	logger *log.Logger
}

func (s *server) loggingMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f(w, r)

		s.logger.Debugf("%s %s", r.Method, r.URL.Path)
	})
}

func (s *server) handleLocate() http.HandlerFunc {
	type request struct {
		DeviceID string `json:"device_id"`
		Measures []struct {
			Sensor string `json:"sensor"`
			RSSI   int64  `json:"rssi"`
			AOA    int64  `json:"aoa"`
		} `json:"measures"`
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var data request

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

		ld := []locator.LocationInputData{}

		for _, d := range data.Measures {
			ld = append(ld, locator.LocationInputData{
				Sensor: d.Sensor,
				RSSI:   float64(d.RSSI),
				AOA:    float64(d.AOA),
			})
		}

		locator := locator.New(s.logger)
		loc := locator.Locate(ld)

		respData, _ := json.Marshal(map[string]float64{
			"X": loc.X,
			"Y": loc.Y,
		})

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})
}

func Start() {
	l := log.New()
	l.SetLevel(log.DebugLevel)

	log.Info("Starting server...")

	s := server{
		router: mux.NewRouter(),
		logger: l,
	}

	s.routes()

	http.Handle("/", s.router)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", http.DefaultServeMux))
}
