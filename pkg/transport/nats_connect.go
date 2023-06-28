package natsCconnect

import (
	"os"

	stan "github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func natsConnect(log *logrus.Logger) (sc stan.Conn, err error) {

	var clusterID string = os.Getenv("nats_clusterID")
	var clientID string = os.Getenv("nats_clientID")

	sc, err = stan.Connect(clusterID, clientID)
	return sc, err
}
