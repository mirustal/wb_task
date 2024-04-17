CREATE TABLE PAYMENT(
    order_id VARCHAR(30) PRIMARY KEY REFERENCES "order"(order_uid),
    "transaction" VARCHAR(50) NOT NULL,
    request_id VARCHAR(30) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    "provider" VARCHAR(30) NOT NULL,
    amount NUMERIC(10, 2) NOT NULL,
    payment_dt BIGINT NOT NULL,
    bank VARCHAR(50) NOT NULL,
    delivery_cost NUMERIC(10, 2) NOT NULL,
    goods_total INTEGER NOT NULL,
    custom_fee NUMERIC(10, 2) NOT NULL
);