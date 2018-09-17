package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karuppaiah/shopping/middleware"
	"github.com/karuppaiah/shopping/model"
	promotionUcase "github.com/karuppaiah/shopping/promotion"
	validator "gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpCartHandler struct {
	AUsecase promotionUcase.EUsecase
}

func NewCartHttpHandler(r *gin.Engine, us promotionUcase.EUsecase) {
	handler := &HttpCartHandler{
		AUsecase: us,
	}
	// JWT auth
	auth := r.Group("/auth")
	authMiddleware := middleware.InitMiddleware().AuthMiddleware()
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/promotion", handler.Fetch)
		auth.POST("/promotion", handler.Store)
		auth.DELETE("/promotion", handler.Delete)

	}

}

//Get promotions
func (a *HttpCartHandler) Fetch(c *gin.Context) {

	ctx := c
	fmt.Println("In Fetch")

	listC, err := a.AUsecase.Fetch(ctx)
	fmt.Println(listC)
	fmt.Println("In Fetch")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	fmt.Println("In Fetch")
	c.JSON(http.StatusOK, gin.H{"Promotion": listC})
}

//Validate request
func isRequestValid(m *model.Promotion) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Add new promotion item
func (a *HttpCartHandler) Store(c *gin.Context) {
	var promotion model.Promotion
	err := c.BindJSON(&promotion)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	if ok, err := isRequestValid(&promotion); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx := c
	fmt.Println("Promotion:", promotion)
	pr, err := a.AUsecase.Store(ctx, &promotion)

	if err != nil {
		fmt.Println("handler error:" + err.Error())
		c.JSON(http.StatusBadRequest, err.Error())

	}
	c.JSON(http.StatusOK, gin.H{"success": pr})
}

// Delete API for promotion
func (a *HttpCartHandler) Delete(c *gin.Context) {
	idP, err := strconv.Atoi(c.Query("id"))
	id := int(idP)
	ctx := c

	_, err = a.AUsecase.Delete(ctx, id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
