package model


type UserLogin struct {
	Phone string `form:"phone" json:"phone" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}