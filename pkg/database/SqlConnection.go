package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func SqlConnect(log *logrus.Logger) (db *sql.DB, err error) {

	var dbuser string = os.Getenv("bduser")
	var dbpass string = os.Getenv("bdpass")
	var connStr string = fmt.Sprintf("user=%s password=%s sslmode=disable", dbuser, dbpass)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Error("не удалось подключиться к базе данных", err)
	} else {
		log.Info("Подключение к базе данных")
	}
	return db, err

}
