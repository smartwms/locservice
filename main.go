package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/smartwms/locservice/server"
)

func main() {
	log.SetLevel(log.DebugLevel)
	server.Start()
}
