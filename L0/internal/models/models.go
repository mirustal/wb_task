package models

import "time"


type Delivery struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Zip     string `json:"zip" validate:"required"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
}

type Payment struct {
	OrderId      string  `gorm:"primaryKey"`
	Transaction  string  `json:"transaction" validate:"required"`
	RequestId    string `json:"request_id"`
	Currency     string  `json:"currency" validate:"required"`
	Provider     string  `json:"provider" validate:"required"`
	Amount       float64 `json:"amount" validate:"required"`
	PaymentDt    int     `json:"payment_dt" validate:"required"`
	Bank         string  `json:"bank" validate:"required"`
	DeliveryCost float64 `json:"delivery_cost" validate:"min=0"`
	GoodsTotal   int     `json:"goods_total" validate:"min=0"`
	CustomFee    float64 `json:"custom_fee" validate:"min=0"`
}

type Item struct {
	Id          int `gorm:"primaryKey"`
	ChrtId      int     `json:"chrt_id" validate:"min=0"`
	TrackNumber string  `json:"track_number" validate:"required"`
	Price       float64 `json:"price" validate:"min=0"`
	Rid         string  `json:"rid" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Sale        int     `json:"sale" validate:"min=0"`
	Size        string  `json:"size" validate:"required"`
	TotalPrice  float64 `json:"total_price" validate:"min=0"`
	NmId        int     `json:"nm_id" validate:"required"`
	Brand       string  `json:"brand" validate:"required"`
	Status      int     `json:"status" validate:"min=0"`
}

type Order struct {
	OrderUid          string    `gorm:"primaryKey" json:"order_uid" validate:"required"`
	TrackNumber       string    `json:"track_number" validate:"required"`
	Entry             string    `json:"entry" validate:"required"`
	Delivery          Delivery  `json:"delivery" validate:"required"`
	Payment           Payment   `json:"payment" validate:"required"`
	Items             []Item    `json:"items" validate:"required"`
	Locale            string    `json:"locale" validate:"required"`
	InternalSignature string   `json:"internal_signature"`
	CustomerId        string    `json:"customer_id" validate:"required"`
	DeliveryService   string    `json:"delivery_service" validate:"required"`
	Shardkey          string    `json:"shardkey" validate:"required"`
	SmId              int       `json:"sm_id" validate:"required"`
	DateCreated       time.Time `json:"date_created" validate:"required"`
	OofShard          string    `json:"oof_shard" validate:"required"`
}

type OrderComplete struct {
    Order    OrderDTO
    Delivery DeliveryDTO
    Payment  PaymentDTO
    Items    []ItemDTO
}

