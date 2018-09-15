package model

type Login struct {
	Username string `gorm:"not null" form:"username" json:"username" binding:"required"`
	Password string `gorm:"not null" form:"password" json:"password" binding:"required"`
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}
