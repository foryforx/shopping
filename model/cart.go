package model

type Cart struct {
	ID     int     `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Code   string  `gorm:"not null" form:"code" json:"code"` //username (one user will have one cart code = username)
	Prodid int     `gorm:"not null" form:"prodid" json:"prodid"`
	Name   string  `gorm:"not null" form:"name" json:"name"`
	Price  float64 `gorm:"not null" form:"price" json:"price"`
	Items  int     `gorm:"not null" form:"items" json:"items"`
	Dprice float64 `gorm:"not null" form:"dprice" json:"dprice"`
	// Status bool    `gorm:"not null" form:"status" json:"status"`
}
