package http

import (
	"fmt"
	"net/http"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	loginUcase "github.com/karuppaiah/shopping/login"
	"github.com/karuppaiah/shopping/middleware"
	"github.com/karuppaiah/shopping/model"
	validator "gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HttpCartHandler struct {
	AUsecase loginUcase.EUsecase
}

func NewLoginHttpHandler(r *gin.Engine, us loginUcase.EUsecase) {
	handler := &HttpCartHandler{
		AUsecase: us,
	}
	// JWT auth
	auth := r.Group("/auth")
	authMiddleware := middleware.InitMiddleware().AuthMiddleware()
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/login", handler.Fetch)
		auth.POST("/login", handler.Store)
		auth.DELETE("/login", handler.Delete)

	}

}

//Get logins
func (a *HttpCartHandler) Fetch(c *gin.Context) {

	ctx := c
	claims := jwt.ExtractClaims(c)
	// fmt.Println("user:", claims["id"].(string))
	user := claims["id"].(string)
	if user != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated user"})
		return
	}
	getUsername := c.Query("username")
	listC, err := a.AUsecase.Fetch(ctx, getUsername)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Login": listC})
}

//Validate request
func isRequestValid(m *model.Login) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Add new login item
func (a *HttpCartHandler) Store(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	// fmt.Println("user:", claims["id"].(string))
	user := claims["id"].(string)
	if user != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated user"})
		return
	}
	var login model.Login
	err := c.BindJSON(&login)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if ok, err := isRequestValid(&login); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c
	fmt.Println("Login:", login)
	pr, err := a.AUsecase.Store(ctx, &login)

	if err != nil {
		fmt.Println("handler error:" + err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return

	}
	c.JSON(http.StatusOK, gin.H{"success": pr})
}

// Delete API for login
func (a *HttpCartHandler) Delete(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	// fmt.Println("user:", claims["id"].(string))
	user := claims["id"].(string)
	if user != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthenticated user"})
		return
	}
	idP := c.Query("username")

	ctx := c

	_, err := a.AUsecase.Delete(ctx, idP)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
