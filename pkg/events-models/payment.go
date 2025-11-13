package eventsmodels

import "time"

type PaymentProcessedEvent struct {
	OrderID       string    `json:"order_id"`
	TransactionID string    `json:"transaction_id"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	Timestamp     time.Time `json:"timestamp"`
}

type PaymentRefundEvent struct {
    OrderID       string    `json:"order_id"`
    TransactionID string    `json:"transaction_id"`
    Reason        string    `json:"reason"`
    Timestamp     time.Time `json:"timestamp"`
}