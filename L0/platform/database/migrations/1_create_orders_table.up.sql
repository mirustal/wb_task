CREATE TABLE "order" (
    id SERIAL PRIMARY KEY,
    order_uid VARCHAR(30) NOT NULL UNIQUE,
    track_number VARCHAR(30) NOT NULL ,
    "entry" VARCHAR(30) NOT NULL,
    locale VARCHAR(2) NOT NULL,
    internal_signature VARCHAR(50) NOT NULL,
    customer_id VARCHAR(30) NOT NULL,
    delivery_service VARCHAR(30) NOT NULL,
    shardkey VARCHAR(30) NOT NULL,
    sm_id BIGINT NOT NULL,
    date_created TIMESTAMP NOT NULL,
    oof_shard VARCHAR(30) NOT NULL
);
