package migration

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func Migration(db *sql.DB, log *logrus.Logger) (err error) {

	res, err := db.Query("SELECT datname FROM pg_catalog.pg_database")
	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		var result string
		err := res.Scan(&result)
		if err != nil {
			fmt.Println(err, "не удалось записать данные в переменную")
		}
		if result == "wb" {
			log.Info("база данных WB существует")
		}
	}

	m, err := migrate.New("file://C:/Users/Asus/Documents/WB_L0/pkg/database/migration/", "postgres://postgres:root@localhost:5432/wb?sslmode=disable")
	if err != nil {
		log.Fatal("не удалось подключится ", err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("не удалось запустить миграцию ", err)
	}
	return err
}
