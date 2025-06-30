package controller

import (
	"blog/internal/model"
	"blog/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserCtl struct {
	UserService *service.UserService
}

func NewUserctl(db *gorm.DB, jwtScrept string) *UserCtl {
	return &UserCtl{UserService: service.NewUserService(db, jwtScrept)}
}

func (uc *UserCtl) Register(c *gin.Context) {
	//绑定参数
	var input model.RegisterUser
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.UserService.Register(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Register successfully",
		"user": gin.H{
			"ID":       user.ID,
			"Username": user.Username,
			"Password": user.Password,
			"Email":    user.Email,
		},
	})
}

func (uc *UserCtl) Login(c *gin.Context) {
	var input model.LoginUser

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := uc.UserService.Login(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successfully",
		"token":   token,
		"user": gin.H{
			"ID":       user.ID,
			"Name":     user.Username,
			"password": user.Password,
		},
	})
}
