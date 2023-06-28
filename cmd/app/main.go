package main

import (
	"github.com/sabitvrustam/WB_L0/pkg/database"
	"github.com/sabitvrustam/WB_L0/pkg/logger"
)

func main() {
	log := logger.Init()
	_, err := database.SqlConnect(log)
	if err != nil {
		panic(err)
	}

	_, err = natsConnect.natsConnect(log)
	if err != nil {
		panic(err)
	}

}
