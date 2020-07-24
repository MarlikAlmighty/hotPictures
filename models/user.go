package models

type UserCache struct {
	MessageID int   `json:"message_id"`
	FloodTime int64 `json:"up_time"`
}
