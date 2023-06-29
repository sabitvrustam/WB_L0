package natsStreaming

import (
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

	return sc, err
}
