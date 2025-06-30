package controller

import (
	"blog/internal/model"
	"blog/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentCtl struct {
	CommentService service.CommentService
}

func NewCommnetCtl(db *gorm.DB) *CommentCtl {
	return &CommentCtl{CommentService: *service.NewCommentService(db)}
}

func (ctl *CommentCtl) CreateComment(c *gin.Context) {
	userID := c.MustGet("UserID").(float64)
	idstr := c.Query("id")

	postID, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ParseUint error")
		return
	}
	var input model.CreateComment
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}
	if err := ctl.CommentService.CreateComment(input, uint(postID), uint(userID)); err != nil {
		c.JSON(http.StatusBadRequest, "Failed Create comment")
		return
	}

	c.JSON(http.StatusOK, "Create comment successfully")
}

func (ctl *CommentCtl) GetComment(c *gin.Context) {
	idstr := c.Query("id")

	postID, err := strconv.ParseUint(idstr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, "ParseUint error")
		return
	}
	com, err := ctl.CommentService.GetComment(uint(postID))
	if err != nil {
		c.JSON(http.StatusNotFound, "get error")
		return
	}
	c.JSON(http.StatusOK, com)
}
