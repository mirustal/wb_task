CREATE TABLE item(
    id SERIAL PRIMARY KEY,
    order_id VARCHAR(30) REFERENCES "order"(order_uid),
    chrt_id BIGINT NOT NULL,
    track_number VARCHAR(50) NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    rid VARCHAR(50) NOT NULL,
    "name" VARCHAR(50) NOT NULL,
    sale INTEGER CHECK(sale >= 0 AND sale <= 100) NOT NULL,
    size VARCHAR(10) NOT NULL,
    total_price NUMERIC(10, 2) NOT NULL,
    nm_id BIGINT NOT NULL,
    brand VARCHAR(50) NOT NULL,
    "status" INTEGER NOT NULL
);