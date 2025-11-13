package eventsmodels

import "time"

type ReservedEvent struct {
	OrderID       string    `json:"order_id"`
	ReservationID string    `json:"reservation_id"`
	Status        string    `json:"status"` 
	Timestamp     time.Time `json:"timestamp"`
}

