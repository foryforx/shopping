package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karuppaiah/shopping/model"
)

func HandleGetProducts(c *gin.Context) {

	allProducts := model.GetProducts()
	c.JSON(http.StatusOK, gin.H{"Products": allProducts})
}

func HandlePostProducts(c *gin.Context) {
	var product model.Product
	errJ := c.Bind(&product)
	if errJ != nil {
		c.JSON(http.StatusOK, gin.H{"error": errJ.Error()})

	}
	model.SaveProduct(&product)
	c.JSON(http.StatusOK, gin.H{"success": product})
}

func HandleDelProducts(c *gin.Context) {
	id := model.DeleteProduct(c.Query("id"))
	c.JSON(http.StatusOK, gin.H{"status": "User #" + id + " deleted"})
}
