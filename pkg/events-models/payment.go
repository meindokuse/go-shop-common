package eventsmodels

import "time"

type PaymentProcessedEvent struct {
	OrderID       string    `json:"order_id"`
	IdmKey	      string    `json:"idm_key"`
	Amount        float64   `json:"amount"`
	Timestamp     time.Time `json:"timestamp"`
}

type PaymentFailedEvent struct {
    OrderID       string    `json:"order_id"`
    IdmKey	      string    `json:"idm_key"`
    Reason        string    `json:"reason"`
    Timestamp     time.Time `json:"timestamp"`
}