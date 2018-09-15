package model

import (
	"errors"
)

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

/// CART Model functions
func GetPromotion() []Promotion {
	db := GetDBInstance().DB
	var promotions []Promotion

	db.Find(&promotions)

	// Display JSON result
	return promotions
}

func AddPromotionItem(promotion *Promotion) error {
	db := GetDBInstance().DB

	// Check if item exists and with inventory
	var sproduct Product
	db.Where("id = ?", promotion.Sprodid).Find(&sproduct)

	if len(sproduct.Name) < 2 {
		return errors.New("Source Product does not exist")
	}
	var dproduct Product
	db.Where("id = ?", promotion.Dprodid).Find(&dproduct)
	if len(dproduct.Name) < 2 {
		return errors.New("Destination Product does not exist")
	}

	db.Create(&promotion)

	return nil
}

func DeletePromotion(id string) error {
	db := GetDBInstance().DB
	// Get id product

	var promotion Promotion
	// SELECT * FROM product WHERE code = 1;
	db.Where("id = ?", id).Find(&promotion)
	if promotion.ID <= 0 {
		return errors.New("promotion doesnt exist")
	}

	db.Where("id = ?", id).Delete(&promotion)

	return nil
}
