package service

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/sabitvrustam/WB_L0/pkg/database/orders"
	"github.com/sabitvrustam/WB_L0/pkg/types"
	"github.com/sirupsen/logrus"
)

type Service struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewService(db *sql.DB, log *logrus.Logger) *Service {
	return &Service{db: db,
		log: log}
}

func (s *Service) OrderWrite(order string) {
	var rezult types.Order
	err := json.Unmarshal([]byte(order), &rezult)
	if err != nil {
		fmt.Println(err, rezult)
	}
	d := orders.NewOrder(s.db, s.log)
	d.OrdersWrite(&rezult)
}
