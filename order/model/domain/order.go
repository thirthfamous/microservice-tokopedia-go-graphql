package domain

import "time"

type Order struct {
	Id          int       `gorm:"primaryKey"`
	ProfileId   int       `json:"profile_id"`
	ProductId   int       `json:"product_id"`
	DateOrdered time.Time `json:"date_ordered"`
	Status      int       `json:"status"`
}

func (Order) TableName() string {
	return "order"
}
