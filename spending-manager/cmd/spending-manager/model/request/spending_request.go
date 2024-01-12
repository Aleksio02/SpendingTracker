package request

import "time"

type SpendingRequest struct {
	Id          int32     `json:"id"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Spent       float32   `json:"spent"`
	CreateDate  time.Time `json:"create_date"`
}
