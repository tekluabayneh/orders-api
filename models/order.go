package model

import (
	"time"
	"uuid"
)

type Order struct {
	OrderID     uint64     `json:"order_id"`
	CustomerID  uuid.UUID  `json:"custsomer_id"`
	lineItems   []LineItem `json:"line_items"`
	OrderStatus string     `json:"orders_status"`
	CreatedAt   *time.Time `json:"create_at"`
	ShippedAT   *time.Time `json:"shipped_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type LineItem struct {
	itemID   uuid.UUID `json:"item_it"`
	Quantity unit      `json:"qulity_unit"`
	Price    unit      `json:"price_unit"`
}
