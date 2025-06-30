package controller

import (
	"blog/internal/model"
	"blog/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Postctl struct {
	PostService service.PostService
}

func NewPostctl(db *gorm.DB) *Postctl {
	return &Postctl{PostService: *service.NewPostService(db)}
}

func (ctl *Postctl) CreatePostCtl(c *gin.Context) {
	userID := c.MustGet("UserID").(float64)

	var input model.CreatePost
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctl.PostService.CreatePost(input, uint(userID)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to created post"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Create post successfully"})
}

func (ctl *Postctl) GetPostCtl(c *gin.Context) {
	ID := c.Query("id")
	id, err := strconv.ParseUint(ID, 10, 64)
	if err != nil {
		// 处理错误，例如返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	post, err := ctl.PostService.GetPost(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "404"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": post})
}

func (ctl *Postctl) GetPostListCtl(c *gin.Context) {

	posts, err := ctl.PostService.GetPostList()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (ctl *Postctl) Updatapost(c *gin.Context) {
	userid := c.MustGet("UserID").(float64)
	idstr := c.Query("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		// 处理错误，例如返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}
	var input model.UpdatePostInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}

	if err := ctl.PostService.UpdatePost(uint(id), input, uint(userid)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to update"})
		return
	}

	c.JSON(http.StatusOK, "update successfully")
}
func (ctl *Postctl) DeletePost(c *gin.Context) {
	userid := c.MustGet("UserID").(float64)
	idstr := c.Query("id")
	id, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		// 处理错误，例如返回错误响应
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id parameter"})
		return
	}

	if err := ctl.PostService.DeletePost(uint(id), uint(userid)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to delete"})
		return
	}

	c.JSON(http.StatusOK, "delete successfully")
}
