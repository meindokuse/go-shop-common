package eventsmodels

import (
	"time"

	servicedto "github.com/meindokuse/go-shop-common/pkg/service-dto"
)


type OrderCreatedEvent struct {
	OrderID     string             `json:"order_id"`
	UserID      string             `json:"user_id"`
	IdmKey	      string    `json:"idm_key"`
	Items       []servicedto.OrderItemDTO `json:"items"`
	TotalAmount float64            `json:"total_amount"`
	Timestamp   time.Time          `json:"timestamp"`
}

type OrderReadyToPay struct {
	OrderID   string    `json:"order_id"`
	IdmKey	      string    `json:"idm_key"`
	TotalAmount float64 `json:"total_amount"`
	Timestamp time.Time `json:"timestamp"`
}

type OrderCancelledEvent struct {
	OrderID   string    `json:"order_id"`
	IdmKey	  string    `json:"idm_key"`
	Reason    string    `json:"reason"`
	Timestamp time.Time `json:"timestamp"`
}