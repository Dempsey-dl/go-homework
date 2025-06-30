package repoisitory

import (
	"blog/internal/model"

	"gorm.io/gorm"
)

type Commentrepoisitory struct {
	db *gorm.DB
}

func NewCommentrepoisitory(db *gorm.DB) *Commentrepoisitory {
	return &Commentrepoisitory{db: db}
}

func (c *Commentrepoisitory) CreateComment(Comment *model.Comment) error {
	return c.db.Create(&Comment).Error
}

func (c *Commentrepoisitory) FindByUserID(ID uint) (*model.Comment, error) {
	var comment model.Comment
	err := c.db.Preload("User").Preload("Post").First(&comment, ID).Error
	return &comment, err
}
