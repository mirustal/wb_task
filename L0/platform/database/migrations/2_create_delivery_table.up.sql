CREATE TABLE delivery(
    order_id VARCHAR(30) PRIMARY KEY REFERENCES "order"(order_uid),
    "name" VARCHAR(50) NOT NULL,
    phone VARCHAR(30) NOT NULL,
    zip VARCHAR(10) NOT NULL,
    city VARCHAR(30) NOT NULL,
    "address" VARCHAR(50) NOT NULL,
    region VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL
);