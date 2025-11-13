package eventsmodels

import "time"

type ProductReservedEvent struct {
	OrderID       string    `json:"order_id"`
	TotalAmount	  int		`json:"total_amount"`
	Timestamp     time.Time `json:"timestamp"`
}

