package fake

import (
	"context"
	"math/rand"
	"sync/atomic"
	"time"

	"gitlab.praktikum-services.ru/Stasyan/momo-store/internal/store/dumplings"
)

// Store is a fake in-memory implementation of dumplings.Store
type Store struct {
	rand              *rand.Rand
	orderID           int64
	availableProducts []dumplings.Product
}

func NewStore() *Store {
	return &Store{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *Store) SetAvailablePacks(products ...dumplings.Product) {
	s.availableProducts = products
}

func (s *Store) ListProducts(_ context.Context) ([]dumplings.Product, error) {
	return s.availableProducts, nil
}

// CreateOrder fakes order creation by incrementing internal id
func (s *Store) CreateOrder(_ context.Context, _ ...dumplings.OrderItem) (id int64, err error) {
	return atomic.AddInt64(&s.orderID, 1), nil
}
