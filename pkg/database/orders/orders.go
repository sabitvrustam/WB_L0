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

func (d *Order) OrdersWrite(order *types.Order) (orderId int64, err error) {
	var deliveriId int64
	var paymentId int64
	var itemId int64
	deliverySq := sq.Insert("delivery").
		Columns("name", "phone", "zip", "city", "address", "region", "email").
		Values(order.Name, order.Phone, order.Zip, order.City, order.Address, order.Region, order.Email).
		Suffix("RETURNING \"id\"")
	err = deliverySq.RunWith(d.db).PlaceholderFormat(sq.Dollar).QueryRow().Scan(&deliveriId)
	if err != nil {
		d.log.Error("не удалось записать данные заказа в базу данных", err)
	}

	paymentSq := sq.Insert("payment").
		Columns("transaction", "request_id", "currency", "provider", "amount",
			"payment_dt", "bank", "delivery_cost", "goods_total", "custom_fee").
		Values(order.Transaction, order.RequestId, order.Currency, order.Provider, order.Amount,
			order.PaymentDt, order.Bank, order.DeliveryCost, order.GoodsTotal, order.CustomFee).
		Suffix("RETURNING \"id\"")
	err = paymentSq.RunWith(d.db).PlaceholderFormat(sq.Dollar).QueryRow().Scan(&paymentId)
	if err != nil {
		d.log.Error("не удалось записать данные заказа в базу данных", err)
	}

	orderSq := sq.Insert("orders").
		Columns("order_uid", "track_number", "entry", "delivery_id", "payment_id", "locate",
			"internal_signature", "customer_id", "delivery_service", "shardkey", "sm_id", "date_created", "oof_shard").
		Values(order.OrderUid, order.TrackNumber, order.Entry, deliveriId, paymentId, order.Locale,
			order.InternalSignature, order.CustomerId, order.DeliveryService, order.Shardkey, order.SmId, order.DateCreated, order.OofShard).
		Suffix("RETURNING \"id\"")
	err = orderSq.RunWith(d.db).PlaceholderFormat(sq.Dollar).QueryRow().Scan(&orderId)
	if err != nil {
		d.log.Error("не удалось записать данные заказа в базу данных", err)
	}

	for _, n := range order.Items {
		itemsSq := sq.Insert("items").
			Columns("chrt_id", "track_number", "price", "rid", "name", "sale", "size", "total_price", "nm_id", "brand", "status").
			Values(n.ChrtId, n.TrackNumber, n.Price, n.Rid, n.Name, n.Sale, n.Size, n.TotalPrice, n.NmId, n.Brand, n.Status).
			Suffix("RETURNING \"id\"")
		err = itemsSq.RunWith(d.db).PlaceholderFormat(sq.Dollar).QueryRow().Scan(&itemId)
		if err != nil {
			d.log.Error("не удалось записать данные заказа в базу данных", err)
		}

		orderItemsSq := sq.Insert("orders_items").
			Columns("order_id", "item_id").
			Values(orderId, itemId)
		_, err = orderItemsSq.RunWith(d.db).PlaceholderFormat(sq.Dollar).Exec()
		if err != nil {
			d.log.Error("не удалось записать данные заказа в базу данных", err)
		}
	}
	d.log.Info(orderId, paymentId, deliveriId, itemId)
	return orderId, err
}
func (d *Order) OrderRead(id int64) (order types.Order, err error) {
	sb := sq.Select("o.order_uid", "o.track_number", "o.entry",
		"d.name", "d.phone", "d.zip", "d.city", "d.address", "d.region", "d.email",
		"p.transaction", "p.request_id", "p.currency", "p.provider", "p.amount", "p.payment_dt",
		"p.bank", "p.delivery_cost", "p.goods_total", "p.custom_fee",
		"o.locate", "o.internal_signature", "o.customer_id", "o.delivery_service", "o.shardkey",
		"o.sm_id", "o.date_created", "o.oof_shard").
		From("orders AS o").
		Join("delivery AS d ON o.delivery_id = d.id").
		Join("payment AS p ON o.payment_id = p.id")
	sb = sb.Where(sq.Eq{"o.id": id})

	res, err := sb.RunWith(d.db).PlaceholderFormat(sq.Dollar).Query()
	if err != nil {
		d.log.Error(err, "не удалось считать заказы из базы данных")
	}
	for res.Next() {
		var result types.Order
		err = res.Scan(&result.OrderUid, &result.TrackNumber, &result.Entry,
			&result.Name, &result.Phone, &result.Zip, &result.City, &result.Address, &result.Region, &result.Email,
			&result.Transaction, &result.RequestId, &result.Currency, &result.Provider, &result.Amount, &result.PaymentDt,
			&result.Bank, &result.DeliveryCost, &result.GoodsTotal, &result.CustomFee,
			&result.Locale, &result.InternalSignature, &result.CustomerId, &result.DeliveryService, &result.Shardkey,
			&result.SmId, &result.DateCreated, &result.OofShard)
		if err != nil {
			d.log.Error(err)
		}
		order = result
	}

	return order, err
}
