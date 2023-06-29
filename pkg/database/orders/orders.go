package orders

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sabitvrustam/WB_L0/pkg/types"
	"github.com/sirupsen/logrus"
)

type Order struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewOrder(db *sql.DB, log *logrus.Logger) *Order {
	return &Order{
		db:  db,
		log: log}
}

func (d *Order) readOrders(id *int64) (orders []*types.Order, err error) {
	sb := sq.Select("o.id", "u.id", "u.f_name", "u.l_name", "u.m_name", "u.n_phone",
		"d.id", "d.type", "d.brand", "d.model", "d.sn",
		"m.id", "m.f_name", "m.l_name", "m.m_name", "m.n_phone", "s.id", "s.o_status").
		From("orders AS o").
		Join("users AS u ON o.id_users = u.id").
		Join("device AS d ON o.id_device = d.id").
		Join("masters AS m ON o.id_masters  = m.id").
		Join("status AS s ON o.id_status  = s.id")
	sb = sb.Where(sq.Eq{"o.id": *id})

	return orders, err
}
