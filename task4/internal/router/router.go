package router

import (
	"blog/internal/config"
	"blog/internal/controller"
	"blog/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {
	router := gin.Default()

	userctl := controller.NewUserctl(db, cfg.JWT.Secret)
	router.POST("/Register", userctl.Register)
	router.POST("/Login", userctl.Login)

	postctl := controller.NewPostctl(db)
	router.GET("/Post", postctl.GetPostCtl)
	router.GET("/PostList", postctl.GetPostListCtl)
	commentctl := controller.NewCommnetCtl(db)
	router.GET("/Comment", commentctl.GetComment)

	auth := router.Group("/auth")
	auth.Use(middleware.JwtAuth(cfg.JWT.Secret))
	{

		auth.POST("/CreatePost", postctl.CreatePostCtl)
		auth.POST("/UpdatePost", postctl.Updatapost)
		auth.POST("/DeletePost", postctl.DeletePost)
		auth.POST("/CreateComment", commentctl.CreateComment)
	}
	return router
}
