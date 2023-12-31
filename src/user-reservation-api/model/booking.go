package model

type Booking struct {
	Id        int  `gorm:"primaryKey"`
	StartDate int  `gorm:"not null"`
	EndDate   int  `gorm:"not null"`
	Rooms     int  `gorm:"not null"`
	User      User `gorm:"foreignkey:UserId"`
	UserId    int
	// como conecto el id del hotel de mongo con este--> ver!
}
type Bookings []Booking
