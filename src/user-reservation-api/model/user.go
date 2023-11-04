package model

type User struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(300); not null"`
	LastName string `gorm:"type:varchar(300); not null"`
	Dni      string `gorm:"type:varchar(8); not null"`
	Email    string `gorm:"type:varchar(300); unique"`
	Password string `gorm:"type:varchar(300); not null"`
	Role     string `gorm:"type:varchar(10); not null"`
}

type Users []User
