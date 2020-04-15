package mqtt

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"

	ap "github.com/smartwms/locservice/pkg/proto"
)

type RawMeasure struct {
	SensorID  string
	TagID     string
	Channel   int64
	RSSI      int64
	Timestamp int64
}

var dumpFile, _ = os.Create("data.txt")

func sensorMessageHanlder(values chan RawMeasure) mqtt.MessageHandler {
	return func(client mqtt.Client, message mqtt.Message) {
		sensor := strings.Split(message.Topic(), "/")[1]

		log.Debug("Received message by sensor: ", sensor)

		incoming := &ap.AdvertisementPacket{}
		err := proto.Unmarshal(message.Payload(), incoming)

		if err != nil {
			log.Error("unmarshaling error: ", err)
			return
		}

		tagAddr := incoming.GetAddress()

		for i := len(tagAddr)/2 - 1; i >= 0; i-- {
			opp := len(tagAddr) - 1 - i
			tagAddr[i], tagAddr[opp] = tagAddr[opp], tagAddr[i]
		}

		stringAddr := hex.EncodeToString(tagAddr)

		data := fmt.Sprintf("%s\t%s\t%d\t%d\t%d\n",
			sensor,
			stringAddr,
			incoming.GetChannel(),
			incoming.GetRssi(),
			incoming.GetTimestamp(),
		)

		dumpFile.WriteString(data)

		log.Println(data)

		values <- RawMeasure{
			SensorID:  sensor,
			TagID:     stringAddr,
			Channel:   int64(incoming.GetChannel()),
			RSSI:      int64(incoming.GetRssi()),
			Timestamp: int64(incoming.GetTimestamp()),
		}
	}
}

func Start(stopChan chan struct{}, measChan chan RawMeasure) error {
	errChan := make(chan error)

	opts := mqtt.NewClientOptions().
		AddBroker("tcp://34.73.184.46:1883").
		SetClientID("locserver").
		SetCleanSession(true)

	opts.OnConnect = func(c mqtt.Client) {
		token := c.Subscribe(
			"d/+/7005/report_data",
			byte(0),
			sensorMessageHanlder(measChan),
		)

		if token.Wait() && token.Error() != nil {
			errChan <- token.Error()
		}
	}

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	log.Infof("Connected to %v", opts.Servers)

	defer func() {
		client.Disconnect(50)
	}()

	select {
	case err := <-errChan:
		return err
	case <-stopChan:
		break
	}

	return nil
}
