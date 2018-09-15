package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karuppaiah/shopping/model"
)

// CART functions
func HandleGetPromotion(c *gin.Context) {
	promotion := model.GetPromotion()
	c.JSON(http.StatusOK, gin.H{"Promotions": promotion})
}

func HandlePostPromotion(c *gin.Context) {
	var promotion model.Promotion
	errJ := c.BindJSON(&promotion)
	if errJ != nil {
		c.JSON(http.StatusOK, gin.H{"error": errJ.Error()})

	}
	err := model.AddPromotionItem(&promotion)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": promotion})
	}

}

func HandleDelPromotion(c *gin.Context) {
	id := c.Query("id")
	err := model.DeletePromotion(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "Promotion id: " + id + " was deleted"})
	}

}
