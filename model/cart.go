package model

import (
	"errors"
)

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

/// CART Model functions
func GetCart(user string) []Cart {
	db := GetDBInstance().DB
	var carts []Cart
	db.Where("code = ?", user).Find(&carts)

	// Display JSON result
	return carts
}

func AddCartItem(cart *Cart) error {
	db := GetDBInstance().DB

	// Check if item exists and with inventory
	var product Product
	db.Where("id = ?", cart.Prodid).Find(&product)
	if len(product.Name) < 2 {
		return errors.New("Product does not exist")
	}
	if product.Stock < cart.Items {
		return errors.New("Product stock not available")
	}
	cart.Price = product.Price
	cart.Name = product.Name
	// Check for already existing item in cart
	// TODO: Future update: to allow update of same cart items
	var cartExist Cart
	// SELECT * FROM product WHERE code = 1;
	db.Where("code = ? AND prodid = ?", cart.Code, cart.Prodid).Find(&cartExist)
	if len(cartExist.Name) > 1 {
		return errors.New("Cart item already exists. Please delete the cart line before adding again")
	} else {

		//Add cart line item
		db.Create(&cart)
	}

	return nil
}

func DeleteCartItem(user string, prodid string) error {
	db := GetDBInstance().DB

	var cart Cart
	// SELECT * FROM product WHERE code = 1;
	db.Where("code = ? AND prodid = ?", user, prodid).Find(&cart)
	if cart.Code == "" {
		return errors.New("Cart item cannot be found. Please try again after sometime")
	}

	// DELETE FROM users WHERE id = user.Id
	db.Where("code = ? AND prodid = ?", user, prodid).Delete(&cart)

	return nil
}
