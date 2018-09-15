package controller

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/karuppaiah/shopping/model"
)

// CART functions
func HandleGetCart(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	// fmt.Println("user:", claims["id"].(string))
	user := claims["id"].(string)
	cart := model.GetCart(user)
	c.JSON(http.StatusOK, gin.H{"Cart": cart})
}

func HandlePostCart(c *gin.Context) {
	// fmt.Println(c.PostForm("name"))
	var cart model.Cart
	errJ := c.BindJSON(&cart)

	if errJ != nil {
		c.JSON(http.StatusOK, gin.H{"error": errJ.Error()})

	}
	claims := jwt.ExtractClaims(c)
	// fmt.Println("user:", claims["id"].(string))
	cart.Code = claims["id"].(string)
	cart.Dprice = 0
	err := model.AddCartItem(&cart)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": cart})
	}

}

func HandleDelCart(c *gin.Context) {
	prodid := c.Query("prodid")
	claims := jwt.ExtractClaims(c)
	code := claims["id"].(string)
	err := model.DeleteCartItem(code, prodid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Cart user: " + code + " and product id :" + prodid + " was deleted"})
	}

}
