package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/dgraph-io/ristretto"
	"github.com/gorilla/mux"
	"github.com/sabitvrustam/WB_L0/pkg/service"
	"github.com/sirupsen/logrus"
)

type OrderAPI struct {
	db      *sql.DB
	log     *logrus.Logger
	service *service.Service
	cache   *ristretto.Cache
}

func NewOrderAPI(db *sql.DB, log *logrus.Logger, cache *ristretto.Cache) *OrderAPI {
	return &OrderAPI{
		db:      db,
		log:     log,
		cache:   cache,
		service: service.NewService(db, log, cache)}
}

func (a *OrderAPI) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	result, err := a.service.OrderRead(id)
	if err != nil {
		a.log.Error(err, "не удалось считать данные ордера из базы данных по ид")
		w.WriteHeader(500)
		return
	}

	marshalResult, err := json.Marshal(result)
	if err != nil {
		a.log.Error(err, "не удалось преобразовать данные Ордера в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(marshalResult)

}
