package domain

type Product struct {
	Id    int `gorm:"primaryKey"`
	Name  string
	Desc  string
	Price int
	Stock int
}

func (Product) TableName() string {
	return "product"
}
