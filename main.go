package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/smartwms/locservice/pkg/db"
	"github.com/smartwms/locservice/pkg/mqtt"
	_ "github.com/smartwms/locservice/server"
)

func main() {
	log.SetLevel(log.DebugLevel)

	errChan := make(chan error)
	stopChan := make(chan struct{})
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		errChan <- mqtt.Start(stopChan)
	}()

	// server.Start()

	client := db.NewClient()
	meas, _ := client.GetTagLastMeasures("dbd93de08ed7")
	log.Infof("%+v", meas)

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
