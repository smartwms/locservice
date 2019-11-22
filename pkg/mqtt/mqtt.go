package mqtt

import (
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func onMessageReceived(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
}

func Start(stopChan chan struct{}) error {
	errChan := make(chan error)

	opts := mqtt.NewClientOptions().
		AddBroker("tcp://localhost:1883").
		SetClientID("locserver").
		SetCleanSession(true)

	opts.OnConnect = func(c mqtt.Client) {
		if token := c.Subscribe("#", byte(0), onMessageReceived); token.Wait() && token.Error() != nil {
			errChan <- token.Error()
		}
	}

	client := mqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	} else {
		fmt.Printf("Connected to %s\n", "localhost")
	}

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
