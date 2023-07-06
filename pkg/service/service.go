package service

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/sabitvrustam/WB_L0/pkg/database/orders"
	"github.com/sabitvrustam/WB_L0/pkg/types"
	"github.com/sirupsen/logrus"
)

type Service struct {
	db    *sql.DB
	log   *logrus.Logger
	cache *ristretto.Cache
}

func NewService(db *sql.DB, log *logrus.Logger, cache *ristretto.Cache) *Service {
	return &Service{db: db,
		log:   log,
		cache: cache}
}

func (s *Service) OrderWrite(order string) {
	var rezult types.Order
	err := json.Unmarshal([]byte(order), &rezult)
	if err != nil {
		fmt.Println(err, rezult)
	}
	d := orders.NewOrder(s.db, s.log)
	orderId, err := d.OrdersWrite(&rezult)

	if err != nil {
		s.log.Error("не удалось произвести запись в таблицу", err)
	}

	ok := s.cache.Set(orderId, &rezult, 1)
	if !ok {
		s.log.Error("не удалось записать данные в кэш")
	}
	time.Sleep(10 * time.Millisecond)
}

func (s *Service) OrderRead(id int64) (err error) {
	value, found := s.cache.Get(id)
	if !found {
		panic("missing value")
	}

	s.log.Info(value)

	return err
}
