package http

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sabitvrustam/WB_L0/pkg/transport/http/api"
	"github.com/sirupsen/logrus"
)

func StartHandler(db *sql.DB, log *logrus.Logger) {

	t := NewTemplates(log)
	r := mux.NewRouter()

	OrderAPI := api.NewOrderAPI(db, log)

	r.HandleFunc("/", t.indexPage)

	r.HandleFunc("/api/users", OrderAPI.GetOrder).Methods("GET")

	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	log.Info("Локальный сервер запущен порт 8080")
	http.ListenAndServe(":8080", r)
}
