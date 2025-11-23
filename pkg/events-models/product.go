package eventsmodels

import "time"

type PaymentProcessed struct {
	IdmKey	      string    `json:"idm_key"`
	Status		 bool		`json:"status"`
	UserID      string      `json:"user_id"`
	OrderID       string    `json:"order_id"`
	TotalAmount	  float64	`json:"total_amount"`
	Reason        string    `json:"reason,omitempty"`
	ReservedItems map[string]int	`json:"reserved_ids,omitempty"`
	Timestamp     time.Time `json:"timestamp"`
}

type ProductRelease struct {
	IdmKey	      string    `json:"idm_key"`
	OrderID       string    `json:"order_id"`
	ReleasedItems map[string]int	`json:"released_ids"`
	Timestamp     time.Time `json:"timestamp"`
}

