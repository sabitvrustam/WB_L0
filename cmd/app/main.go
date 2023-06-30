package main

import (
	"github.com/sabitvrustam/WB_L0/pkg/database"
	//"github.com/sabitvrustam/WB_L0/pkg/database/migration"
	"github.com/sabitvrustam/WB_L0/pkg/logger"
	"github.com/sabitvrustam/WB_L0/pkg/transport/http"
	"github.com/sabitvrustam/WB_L0/pkg/transport/natsStreaming"
)

func main() {
	log := logger.Init()

	db, err := database.SqlConnect(log)
	if err != nil {
		panic(err)
	}

	//err = migration.Migration(db, log)

	_, err = natsStreaming.NatsConnect(log)
	if err != nil {
		panic(err)
	}

	http.StartHandler(db, log)

}
