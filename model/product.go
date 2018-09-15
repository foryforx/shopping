package model

type Product struct {
	ID    int     `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Name  string  `gorm:"not null" form:"name" json:"name"`
	Price float64 `gorm:"not null" form:"price json:"price"`
	Stock int     `gorm:"not null" form:"stock" json:"stock"`
}
