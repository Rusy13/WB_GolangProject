package storage

import (
	"context"

	"WB/internal/order/model"
	"WB/internal/order/storage/database/dto"
)

type OrderStorage interface {
	AddOrder(ctx context.Context, order model.Order) (*model.Order, error)
	GetOrderByID(ctx context.Context, orderUID string) (*model.Order, error)
	GetOrderFromCache(orderID string) (*dto.OrderFromCache, error)
	SaveOrderToCache(bannerFromCache dto.OrderFromCache, orderID string) error
}
