package nats_connect

import (
	"os"

	stan "github.com/nats-io/stan.go"
)

func nats_connect() {

	var clusterID string = os.Getenv("nats_clusterID")
	var clientID string = os.Getenv("nats_clientID")

	sc, _ := stan.Connect(clusterID, clientID)
}
