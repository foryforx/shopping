package model

import (
	"fmt"
	"strconv"
)

type Product struct {
	ID    int     `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Name  string  `gorm:"not null" form:"name" json:"name"`
	Price float64 `gorm:"not null" form:"price json:"price"`
	Stock int     `gorm:"not null" form:"stock" json:"stock"`
}

///PRODUCT model function

func GetProducts() []Product {
	db := GetDBInstance().DB
	var products []Product
	// SELECT * FROM users
	db.Find(&products)

	// Display JSON result
	return products
}

func SaveProduct(product *Product) {
	db := GetDBInstance().DB
	db.Create(&product)
}

func DeleteProduct(id string) string {
	db := GetDBInstance().DB

	var product Product
	fmt.Println("Given id", id)
	db.First(&product, id)
	fmt.Println("Product id", product.ID, "Given id", id)
	if strconv.Itoa(product.ID) == id {
		// DELETE FROM users WHERE id = user.Id
		db.Delete(&product)
		// For Display JSON result
		return id
	} else {
		// For Display JSON error
		return "0"
	}

}
