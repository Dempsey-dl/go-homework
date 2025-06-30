package service

import (
	"blog/internal/model"
	"blog/internal/repoisitory"

	"gorm.io/gorm"
)

type CommentService struct {
	CommentRepo *repoisitory.Commentrepoisitory
}

func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{CommentRepo: repoisitory.NewCommentrepoisitory(db)}
}

func (s *CommentService) CreateComment(input model.CreateComment, postID uint, userID uint) error {
	Comment := model.Comment{
		Content: input.Content,
		UserID:  userID,
		PostID:  postID,
	}

	if err := s.CommentRepo.CreateComment(&Comment); err != nil {
		return err
	}
	return nil
}

func (s *CommentService) GetComment(id uint) (*model.Comment, error) {
	comment, err := s.CommentRepo.FindByUserID(id)
	if err != nil {
		return nil, err
	}
	return comment, nil
}
