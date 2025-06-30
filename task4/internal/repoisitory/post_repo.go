package repoisitory

import (
	"blog/internal/model"

	"gorm.io/gorm"
)

type PostRepoisitory struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepoisitory {
	return &PostRepoisitory{db: db}
}

func (p *PostRepoisitory) Create(post *model.Post) error {
	return p.db.Create(&post).Error
}

func (p *PostRepoisitory) FindByID(ID uint) (*model.Post, error) {
	var Post model.Post
	err := p.db.Preload("User").Preload("Comment.User").First(&Post, ID).Error
	return &Post, err
}

func (p *PostRepoisitory) FindByUserID(UserID uint) (*model.Post, error) {
	var Post model.Post
	err := p.db.Preload("Comment").Where("user_id = ?", UserID).First(&Post).Error
	return &Post, err
}

func (p *PostRepoisitory) FindList() ([]model.Post, error) {
	var Posts []model.Post
	err := p.db.Preload("User").Find(&Posts).Error
	return Posts, err
}

func (p *PostRepoisitory) Update(post *model.Post) error {
	return p.db.Save(post).Error
}

func (p *PostRepoisitory) DeleteByUserId(post *model.Post) error {
	return p.db.Delete(post).Error
}
