package api

import (
	"database/sql"
	// "encoding/json"
	// "net/http"
	// "strconv"

	// "github.com/gorilla/mux"
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
