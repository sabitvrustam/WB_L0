package natsStreaming

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/dgraph-io/ristretto"
	stan "github.com/nats-io/stan.go"
	"github.com/sabitvrustam/WB_L0/pkg/service"
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

func NatsWrit(sc stan.Conn) {
	jsonFile, err := os.Open("pkg/transport/natsStreaming/model.json")
	if err != nil {
		log.Fatal(err)
	}
	order, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	for {
		err := sc.Publish("foo", []byte(order))
		time.Sleep(1 * time.Second)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func NatsRead(sc stan.Conn, db *sql.DB, log *logrus.Logger, cash *ristretto.Cache) {
	s := service.NewService(db, log, cash)
	_, _ = sc.Subscribe("foo", func(m *stan.Msg) {
		s.OrderWrite(string(m.Data))
	}, stan.DeliverAllAvailable())
}
