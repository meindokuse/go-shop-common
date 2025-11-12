package eventsmodels

import "time"

type InventoryReservedEvent struct {
	OrderID       string    `json:"order_id"`
	ReservationID string    `json:"reservation_id"`
	Status        string    `json:"status"` 
	Timestamp     time.Time `json:"timestamp"`
}

type InventoryReleaseEvent struct {
    OrderID       string    `json:"order_id"`
    ReservationID string    `json:"reservation_id"`
    Reason        string    `json:"reason"`
    Timestamp     time.Time `json:"timestamp"`
}