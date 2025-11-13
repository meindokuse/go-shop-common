package eventsmodels

import "time"

type ProductDescrateSuccessEvent struct {
	IdmKey	      string    `json:"idm_key"`
	UserID      string      `json:"user_id"`
	OrderID       string    `json:"order_id"`
	TotalAmount	  int		`json:"total_amount"`
	Timestamp     time.Time `json:"timestamp"`
}

type ProductDescrateFailedEvent struct {
	IdmKey	      string    `json:"idm_key"`
	OrderID       string    `json:"order_id"`
	Reason        string    `json:"reason"`
	Timestamp     time.Time `json:"timestamp"`
}

