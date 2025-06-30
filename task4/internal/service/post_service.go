package service

import (
	"blog/internal/model"
	"blog/internal/repoisitory"
	"errors"

	"gorm.io/gorm"
)

type PostService struct {
	PostRepo *repoisitory.PostRepoisitory
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{PostRepo: repoisitory.NewPostRepository(db)}
}

func (s *PostService) CreatePost(input model.CreatePost, userID uint) error {
	post := model.Post{
		Title:   input.Title,
		Content: input.Content,
		UserID:  userID,
	}
	if err := s.PostRepo.Create(&post); err != nil {
		return err
	}
	return nil
}

func (s *PostService) GetPost(id uint) (*model.Post, error) {
	return s.PostRepo.FindByID(id)
}

func (s *PostService) GetPostList() ([]model.Post, error) {
	return s.PostRepo.FindList()
}

func (s *PostService) DeletePost(id uint, userID uint) error {
	post, err := s.PostRepo.FindByID(id)
	if err != nil {
		return err
	}
	if post.UserID != userID {
		return errors.New("authrization: you can only update your own posts")
	}
	if err := s.PostRepo.DeleteByUserId(post); err != nil {
		return err
	}
	return nil
}

func (s *PostService) UpdatePost(id uint, input model.UpdatePostInput, userID uint) error {
	post, err := s.PostRepo.FindByID(id)
	if err != nil {
		return err
	}
	if post.UserID != userID {
		return errors.New("authrization: you can only update your own posts")
	}
	post.Title = input.Title
	post.Content = input.Content
	if err := s.PostRepo.Update(post); err != nil {
		return err
	}
	return nil
}
