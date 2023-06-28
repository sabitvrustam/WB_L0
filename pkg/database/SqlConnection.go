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
	var pass string = fmt.Sprintf("%s:%s@tcp(localhost:5432)/wb", dbuser, dbpass)
	db, err = sql.Open("postgres", pass)
	if err != nil {
		log.Error("не удалось подключиться к базе данных", err)
	} else {
		log.Info("Подключение к базе данных")
	}
	return db, err

}
