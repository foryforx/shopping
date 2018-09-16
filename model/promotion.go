package model

type Promotion struct {
	ID       int     `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Sprodid  int     `gorm:"not null" form:"sprodid" json:"sprodid"`
	Sminqty  int     `gorm:"not null" form:"sminqty" json:"sminqty"`
	Dprodid  int     `gorm:"not null" form:"dprodid" json:"dprodid"`
	Dminqty  int     `gorm:"not null" form:"dminqty" json:"dminqty"`
	Disctype string  `gorm:"not null" form:"disctype" json:"disctype"`
	Discount float64 `gorm:"not null" form:"discount" json:"discount"`
	Priority int     `gorm:"not null" form:"priority" json:"priority"`
}
