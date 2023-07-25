package dumplings

import (
	"context"
)

//go:generate mockgen -source=$GOFILE -package=mock -destination=mock/store.gen.go Store

// Store represents abstract dumplings store
type Store interface {
	// ListProducts returns all available dumplings products
	ListProducts(ctx context.Context) ([]Product, error)
	// CreateOrder stores new order into store
	CreateOrder(ctx context.Context, items ...OrderItem) (id int64, err error)
}
