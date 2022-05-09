package domain

type Profile struct {
	Id      int `gorm:"primaryKey"`
	Name    string
	Address string
}

func (Profile) TableName() string {
	return "customer_profile"
}
