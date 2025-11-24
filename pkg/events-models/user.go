package eventsmodels

import "time"

type UserCreated struct {
	IdmKey string `json:"idm_key"`
	UserID string `json:"user_id"`
	Timestamp   time.Time `json:"timestamp"`
}

type UserUpdateName struct {
	UserID string `json:"user_id"`
	NewName string `json:"new_name"`
}
