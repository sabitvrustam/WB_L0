package main

import (
	"github.com/sabitvrustam/WB_L0/pkg/cashe"
	"github.com/sabitvrustam/WB_L0/pkg/database"
	"github.com/sabitvrustam/WB_L0/pkg/database/migration"
	"github.com/sabitvrustam/WB_L0/pkg/logger"

	//"github.com/sabitvrustam/WB_L0/pkg/service"
	"github.com/sabitvrustam/WB_L0/pkg/transport/http"
	"github.com/sabitvrustam/WB_L0/pkg/transport/natsStreaming"
)

func main() {
	log := logger.Init()

	db, err := database.SqlConnect(log)
	if err != nil {
		panic(err)
	}

	err = migration.Migration(db, log)
	if err != nil {
		log.Fatal(err)
	}

	cache, err := cashe.StartCashe(log)
	if err != nil {
		log.Fatal(err)
	}

	sc, err := natsStreaming.NatsConnect(log)
	if err != nil {
		panic(err)
	}
	go natsStreaming.NatsWrit(sc)
	go natsStreaming.NatsRead(sc, db, log, cache)

	//service := service.NewService(db, log)

	http.StartHandler(db, log, cache)

}
