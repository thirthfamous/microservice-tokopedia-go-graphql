package domain

type User struct {
	Username  string `gorm:"primaryKey"`
	Password  string
	ProfileId int
	Profile   Profile `gorm:"foreignKey:ProfileId"`
}

func (User) TableName() string {
	return "customer_user"
}
