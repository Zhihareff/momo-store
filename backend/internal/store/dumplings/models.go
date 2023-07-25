package dumplings

// Product represents one pack of dumplings
type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

// OrderItem represents one item in user order
type OrderItem struct {
	Pack  Product
	Count uint32
}
