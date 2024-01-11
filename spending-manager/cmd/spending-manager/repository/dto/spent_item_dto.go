package dto

import "time"

type SpentItemDTO struct {
	Id          int32
	Category    string
	Description string
	Spent       float32
	CreateDate  time.Time
	Username    string
}
