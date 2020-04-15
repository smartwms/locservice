package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/smartwms/locservice/pkg/db"
	"github.com/smartwms/locservice/pkg/mqtt"
	"github.com/smartwms/locservice/server"
)

func main() {
	logger := log.New()
	logger.SetLevel(log.DebugLevel)

	errChan := make(chan error)
	stopChan := make(chan struct{})
	exitChan := make(chan os.Signal, 1)
	measChan := make(chan mqtt.RawMeasure, 30)
	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)

	// loc := locator.New(logger, db.NewClient())
	// loc.Init()

	go func() {
		errChan <- mqtt.Start(stopChan, measChan)
	}()

	go func() {
		errChan <- DumpRawMeasure(measChan, "5ccf7fdb3643")
	}()

	go func() {
		errChan <- server.Start()
	}()

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-exitChan:
			log.Infof("Captured %v. Exiting...", s)
			close(stopChan)
			time.Sleep(100 * time.Millisecond)
			os.Exit(0)
		}
	}
}

func DumpRawMeasure(measures chan mqtt.RawMeasure, sensor string) error {
	client := db.NewClient()

	for {
		meas := <-measures

		if meas.SensorID != sensor {
			continue
		}

		err := client.AddRawMeasure(
			meas.SensorID,
			meas.TagID,
			meas.Timestamp,
			meas.Channel,
			meas.RSSI,
		)

		if err != nil {
			return err
		}
	}
}
