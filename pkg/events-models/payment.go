package eventsmodels

import "time"

type PaymentProcessedEvent struct {
    OrderID       string   			`json:"order_id"`
    IdmKey	      string   			`json:"idm_key"`
	Status		  bool				`json:"status"`
	ReservedItems map[string]int	`json:"reserved_ids,omitempty"`
    Reason        string    		`json:"reason,omitempty"`
    Timestamp     time.Time 		`json:"timestamp"`
}

type PaymentRefundEvent struct {
	OrderID       string    `json:"order_id"`
	IdmKey	      string    `json:"idm_key"`
	Status		  bool				`json:"status"`
	Amount        float64   `json:"amount"`
	Timestamp     time.Time `json:"timestamp"`
}