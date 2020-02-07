#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE TABLE orders (
        order_id SERIAL,
        customer_name TEXT,
        item_description TEXT,
        item_price FLOAT,
        quantity INT,
        merchant_name TEXT,
        merchant_address TEXT
    );
EOSQL
