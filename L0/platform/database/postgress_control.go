package database

import (
	"context"

	"task-l0/internal/cache"
	"task-l0/internal/models"
	"task-l0/internal/models/queries"
)


var _ queries.Storage = (*API)(nil)


type API struct {
    db    Client
    cache *cache.Cache
}


func NewApi(db Client, cache *cache.Cache) *API {
    return &API{
        db:    db,
        cache: cache,
    }
}

func (d *API) CreateOrder(ctx context.Context, orderDTO models.OrderDTO) error {

	orderQuery := `
	INSERT INTO "order" (
		order_uid, track_number, "entry", locale, internal_signature, customer_id, delivery_service, 
		shardkey, sm_id, date_created, oof_shard
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	err := d.db.QueryRowContext(ctx, orderQuery, orderDTO.OrderUid, orderDTO.TrackNumber, orderDTO.Entry, orderDTO.Locale, 
		orderDTO.InternalSignature, orderDTO.CustomerId, orderDTO.DeliveryService, orderDTO.Shardkey, orderDTO.SmId, 
		orderDTO.DateCreated, orderDTO.OofShard)
	// if err != nil {
	// 	return fmt.Errorf("Insert order %v", err)
	// }
	print(err)
	deliveryQuery := `
	INSERT INTO delivery (
		order_id, "name", phone, zip, city, "address", region, email
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`
	err = d.db.QueryRowContext(ctx, deliveryQuery, orderDTO.OrderUid, orderDTO.Delivery.Name, orderDTO.Delivery.Phone, 
	orderDTO.Delivery.Zip, orderDTO.Delivery.City, orderDTO.Delivery.Address, orderDTO.Delivery.Region, orderDTO.Delivery.Email)
	// if err != nil {
	// 	return fmt.Errorf("Insert delivery %v", err)
	// }

	paymentQuery := `
	INSERT INTO payment (
		order_id, "transaction", request_id, currency, "provider", amount, payment_dt, 
		bank, delivery_cost, goods_total, custom_fee
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	err = d.db.QueryRowContext(ctx, paymentQuery, orderDTO.OrderUid, orderDTO.Payment.Transaction, orderDTO.Payment.RequestId, 
	orderDTO.Payment.Currency, orderDTO.Payment.Provider, orderDTO.Payment.Amount, orderDTO.Payment.PaymentDt, 
	orderDTO.Payment.Bank, orderDTO.Payment.DeliveryCost, orderDTO.Payment.GoodsTotal, orderDTO.Payment.CustomFee)
	// if err != nil {
	// 	return fmt.Errorf("Insert payment %v", err)
	// }


	itemQuery := `
	INSERT INTO item (
		order_id, chrt_id, track_number, price, rid, "name", sale, size, total_price, nm_id, brand, "status"
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
	for _, item := range orderDTO.Items {
		err = d.db.QueryRowContext(ctx, itemQuery, orderDTO.OrderUid, item.ChrtId, item.TrackNumber, item.Price, 
		item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmId, item.Brand, item.Status)
		// if err != nil {
		// 	return fmt.Errorf("Insert item %v", err)
		// }
	}

	return nil
}

func (d *API) LoadAndCacheOrders(ctx context.Context) error {
    orderIDsQuery := `SELECT order_uid FROM "order"`
    rows, err := d.db.QueryContext(ctx, orderIDsQuery)
    if err != nil {
        return err
    }
    defer rows.Close()

    var orderUIDs []string
    for rows.Next() {
        var orderUID string
        if err := rows.Scan(&orderUID); err != nil {
            return err
        }
        orderUIDs = append(orderUIDs, orderUID)
    }


    for _, orderUID := range orderUIDs {
        if err := d.GetOrder(ctx, orderUID); err != nil {
            return nil
        }
    }

    return nil
}

func (d *API) GetOrder(ctx context.Context, orderUID string) error {
    orderQuery := `
	SELECT order_uid, track_number, "entry", locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard 
	FROM "order"
	 WHERE order_uid = $1
	 `
    var order models.OrderDTO
    err := d.db.QueryRowContext(ctx, orderQuery, orderUID).Scan(&order.OrderUid, &order.TrackNumber, &order.Entry, &order.Locale, &order.InternalSignature, &order.CustomerId, &order.DeliveryService, &order.Shardkey, &order.SmId, &order.DateCreated, &order.OofShard)
    if err != nil {
        return err
    }


    deliveryQuery := `
	SELECT "name", phone, zip, city, "address", region, email 
	FROM delivery 
	WHERE order_id = $1`
    var delivery models.DeliveryDTO
    err = d.db.QueryRowContext(ctx, deliveryQuery, orderUID).Scan(&delivery.Name, &delivery.Phone, &delivery.Zip, &delivery.City, &delivery.Address, &delivery.Region, &delivery.Email)
    if err != nil {
        return err
    }


    paymentQuery := `
	SELECT "transaction", request_id, currency, "provider", amount, payment_dt, bank, delivery_cost, goods_total, custom_fee 
	FROM payment 
	WHERE order_id = $1
	`
    var payment models.PaymentDTO
    err = d.db.QueryRowContext(ctx, paymentQuery, orderUID).Scan(&payment.Transaction, &payment.RequestId, &payment.Currency, &payment.Provider, &payment.Amount, &payment.PaymentDt, &payment.Bank, &payment.DeliveryCost, &payment.GoodsTotal, &payment.CustomFee)
    if err != nil {
        return err
    }


    itemsQuery := `
	SELECT chrt_id, track_number, price, rid, "name", sale, size, total_price, nm_id, brand, "status" 
	FROM item 
	WHERE order_id = $1`
    rows, err := d.db.QueryContext(ctx, itemsQuery, orderUID)
    if err != nil {
        return err
    }
    defer rows.Close()

    var items []models.ItemDTO
    for rows.Next() {
        var item models.ItemDTO
        if err := rows.Scan(&item.ChrtId, &item.TrackNumber, &item.Price, &item.Rid, &item.Name, &item.Sale, &item.Size, &item.TotalPrice, &item.NmId, &item.Brand, &item.Status); err != nil {
            return err
        }
        items = append(items, item)
    }

	order.Delivery = delivery
	order.Items = items
	order.Payment = payment

    d.cache.Set(orderUID, order)

    return nil
}


