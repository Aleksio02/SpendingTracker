package dto

import "time"

type SpentItemDTO struct {
	Id          int32     `db:"id"`
	Category    string    `db:"category"`
	Description string    `db:"description"`
	Spent       float32   `db:"spent"`
	CreateDate  time.Time `db:"create_date"`
	Username    string    `db:"username"`
}
