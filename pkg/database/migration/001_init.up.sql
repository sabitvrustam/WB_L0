CREATE TABLE IF NOT EXISTS delivery(
   id serial PRIMARY KEY,
   name VARCHAR (50) NOT NULL,
   phone VARCHAR (50) NOT NULL,
   zip VARCHAR (300) NOT NULL,
   city VARCHAR (50) NOT NULL,
   address VARCHAR (50) NOT NULL,
   region VARCHAR (50) NOT NULL,
   email VARCHAR (50) NOT NULL
);

CREATE TABLE IF NOT EXISTS payment(
   id serial PRIMARY KEY,
   transaction VARCHAR (50) NOT NULL,
   request_id VARCHAR (50),
   currency VARCHAR (50) NOT NULL,
   provider VARCHAR (50) NOT NULL,
   amount INTEGER  NOT NULL,
   payment_dt INTEGER NOT NULL,
   bank VARCHAR (50) NOT NULL,
   delivery_cost INTEGER NOT NULL,
   goods_total INTEGER NOT NULL,
   custom_fee INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS items(
   id serial PRIMARY KEY,
   chrt_id VARCHAR (50) NOT NULL,
   track_number VARCHAR (50) NOT NULL,
   price INTEGER  NOT NULL,
   rid VARCHAR (50) NOT NULL,
   name VARCHAR (50) NOT NULL,
   sale INTEGER  NOT NULL,
   size VARCHAR (50) NOT NULL,
   total_price INTEGER  NOT NULL,
   nm_id INTEGER  NOT NULL,
   brand VARCHAR (50) NOT NULL,
   status INTEGER  NOT NULL
);

CREATE TABLE IF NOT EXISTS orders(
   id serial PRIMARY KEY,
   order_uid VARCHAR (50) NOT NULL,
   track_number VARCHAR (50) NOT NULL,
   entry VARCHAR (300) NOT NULL,
   delivery_id INTEGER NOT NULL,
   payment_id INTEGER NOT NULL,
   locate VARCHAR (50) NOT NULL,
   internal_signature VARCHAR (50) NOT NULL,
   customer_id VARCHAR (50) NOT NULL,
   delivery_service VARCHAR NOT NULL,
   shardkey VARCHAR (50) NOT NULL,
   sm_id INTEGER NOT NULL,
   date_created VARCHAR (50) NOT NULL,
   oof_shard VARCHAR (50) NOT NULL,
   CONSTRAINT fk_delivery
      FOREIGN KEY(delivery_id) 
	  REFERENCES delivery(id)
     ON DELETE RESTRICT ON UPDATE RESTRICT,
   CONSTRAINT fk_payment
      FOREIGN KEY(payment_id) 
	  REFERENCES payment(id)
     ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE TABLE IF NOT EXISTS orders_items(
   id serial PRIMARY KEY,
   order_id INT NOT NULL,
   item_id INT NOT NULL,
   CONSTRAINT fk_order
      FOREIGN KEY(order_id) 
	  REFERENCES orders(id)
     ON DELETE RESTRICT ON UPDATE RESTRICT,
   CONSTRAINT fk_item
      FOREIGN KEY(item_id) 
	  REFERENCES items(id)
     ON DELETE RESTRICT ON UPDATE RESTRICT
);