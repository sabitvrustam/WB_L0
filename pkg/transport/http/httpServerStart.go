package http

import (
	"database/sql"
	"net/http"

	"github.com/dgraph-io/ristretto"
	"github.com/gorilla/mux"
	"github.com/sabitvrustam/WB_L0/pkg/transport/http/api"
	"github.com/sirupsen/logrus"
)

func StartHandler(db *sql.DB, log *logrus.Logger, cache *ristretto.Cache) {

	t := NewTemplates(log)
	r := mux.NewRouter()

	OrderAPI := api.NewOrderAPI(db, log, cache)

	r.HandleFunc("/", t.indexPage)

	r.HandleFunc("/api/order/{id:[0-9]+}", OrderAPI.GetOrder).Methods("GET")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	log.Info("Локальный сервер запущен порт 8080")
	http.ListenAndServe(":8081", r)
}
