package main

import (
	"github.com/sabitvrustam/WB L0/pkg/logger"
)

func main() {
	log := logger.Init()
	db, err := database.MySqlConnect(log)
	if err != nil {
		panic(err)
	}

	telegramm := telegram.NewTelegram(db, log)
	database.Migrate(db, log)
	go telegramm.Tgbot()
	http.StartHandler(db, log)

}
