package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sabitvrustam/WB_L0/pkg/database/orders"
	"github.com/sirupsen/logrus"
)

type OrderAPI struct {
	db    *sql.DB
	log   *logrus.Logger
	order *orders.Order
}

func NewOrderAPI(db *sql.DB, log *logrus.Logger) *OrderAPI {
	return &OrderAPI{
		db:    db,
		log:   log,
		order: orders.NewOrder(db, log)}
}

func (a *OrderAPI) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	result, err := a.order.GetOrderByID(id)
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
