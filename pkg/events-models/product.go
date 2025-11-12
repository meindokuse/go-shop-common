package eventsmodels

import "time"

type InventoryReservedEvent struct {
	OrderID       string    `json:"order_id"`
	ReservationID string    `json:"reservation_id"`
	Status        string    `json:"status"` 
	Error         string    `json:"error,omitempty"`
	Timestamp     time.Time `json:"timestamp"`
}

type InventoryReleaseRequestedEvent struct {
    OrderID       string    `json:"order_id"`
    ReservationID string    `json:"reservation_id"`
    Reason        string    `json:"reason"`
    Timestamp     time.Time `json:"timestamp"`
}