package queries

import (
	"context"
	"task-l0/internal/models"
)




type Storage interface {
	CreateOrder(context.Context, models.OrderDTO) error
	GetOrder(context.Context, string) error
 }