package natsStreaming

import (
	"fmt"
	"os"

	stan "github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func NatsConnect(log *logrus.Logger) (sc stan.Conn, err error) {

	var clusterID string = os.Getenv("nats_clusterID")
	var clientID string = os.Getenv("nats_clientID")

	sc, err = stan.Connect(clusterID, clientID)
	if err != nil {
		log.Error("не удалось подключиться к nats stening", err)
	} else {
		log.Info("Подключение к nats stening")
	}

	sc.Publish("foo", []byte("Hello World")) // does not return until an ack has been received from NATS Streaming

	// Simple Async Subscriber
	sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	sub.Unsubscribe()

	return sc, err
}
