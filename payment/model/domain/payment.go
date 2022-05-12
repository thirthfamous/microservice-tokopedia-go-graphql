package domain

import "time"

type Payment struct {
	Id       int       `gorm:"primaryKey"`
	OrderId  int       `json:"order_id"`
	PaidDate time.Time `json:"paid_date"`
	Status   int       `json:"status" comment:"1 = menunggu dibayar, 2 = paid/telah dibayar"`
}

func (Payment) TableName() string {
	return "payment"
}
