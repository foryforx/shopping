package http

import (
	"fmt"
	"net/http"
	"strconv"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	cartUcase "github.com/karuppaiah/shopping/cart"
	"github.com/karuppaiah/shopping/middleware"
	"github.com/karuppaiah/shopping/model"
	validator "gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpCartHandler struct {
	AUsecase cartUcase.EUsecase
}

func NewCartHttpHandler(r *gin.Engine, us cartUcase.EUsecase) {
	handler := &HttpCartHandler{
		AUsecase: us,
	}

	auth := r.Group("/auth")
	authMiddleware := middleware.InitMiddleware().AuthMiddleware()
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/cart", handler.Fetch)
		auth.POST("/cart", handler.Store)
		auth.DELETE("/cart", handler.Delete)

	}

}

func (a *HttpCartHandler) Fetch(c *gin.Context) {

	ctx := c
	fmt.Println("In Fetch")
	claims := jwt.ExtractClaims(c)
	// fmt.Println("user:", claims["id"].(string))
	user := claims["id"].(string)
	listC, err := a.AUsecase.Fetch(ctx, user)
	fmt.Println(listC)
	fmt.Println("In Fetch")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	fmt.Println("In Fetch")
	c.JSON(http.StatusOK, gin.H{"Cart": listC})
}
func isRequestValid(m *model.Cart) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *HttpCartHandler) Store(c *gin.Context) {
	var cart model.Cart
	err := c.BindJSON(&cart)
	claims := jwt.ExtractClaims(c)
	// fmt.Println("user:", claims["id"].(string))
	user := claims["id"].(string)
	cart.Code = user

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	}

	if ok, err := isRequestValid(&cart); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx := c
	fmt.Println("Cart:", cart)
	pr, err := a.AUsecase.Store(ctx, &cart)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"success": pr})
}

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
