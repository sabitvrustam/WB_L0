package service

import (
	"database/sql"
	"encoding/json"
	"errors"
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

func (s *Service) OrderWrite(order string) error {
	var rezult types.Order
	err := json.Unmarshal([]byte(order), &rezult)
	if err != nil {
		return err
	}

	d := orders.NewOrder(s.db, s.log)
	orderId, err := d.OrdersWrite(&rezult)
	if err != nil {
		return fmt.Errorf("ошибка записи в БД: %w", err)
	}

	ok := s.cache.Set(orderId, order, 1)
	if !ok {
		return errors.New("не удалось записать данные в кэш")
	}
	time.Sleep(10 * time.Millisecond)
	return nil
}

func (s *Service) OrderRead(id int64) (order types.Order, err error) {
	value, found := s.cache.Get(id)
	if !found {
		d := orders.NewOrder(s.db, s.log)
		order, err := d.OrderRead(id)
		ok := s.cache.Set(id, &order, 1)
		if !ok {
			s.log.Error("не удалось записать данные в кэш")
		}
		time.Sleep(10 * time.Millisecond)
		return order, err
	}
	res, ok := value.(string)
	if !ok {
		return order, errors.New("не удалось преобразовать тип данных")
	}
	err = json.Unmarshal([]byte(res), &order)
	if err != nil {
		fmt.Println(err, order)
	}

	s.log.Info(order)

	return order, err
}
