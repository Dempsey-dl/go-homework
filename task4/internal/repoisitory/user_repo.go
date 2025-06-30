package repoisitory

import (
	"blog/internal/model"

	"gorm.io/gorm"
)

type UserRepoisitory struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepoisitory {
	return &UserRepoisitory{db: db}
}

func (uRepo *UserRepoisitory) CreateUser(user *model.User) error {
	return uRepo.db.Create(&user).Error
}

func (uRepo *UserRepoisitory) FindByID(userID uint) (*model.User, error) {
	var user model.User
	err := uRepo.db.First(&user, userID).Error
	return &user, err
}
func (uRepo *UserRepoisitory) FindByName(Name string) (*model.User, error) {
	var user model.User
	err := uRepo.db.First(&user, "username = ?", Name).Error
	return &user, err
}

func (uRepo *UserRepoisitory) FindByEmail(Email string) (*model.User, error) {
	var user model.User
	err := uRepo.db.First(&user, "Email = ?", Email).Error
	return &user, err
}
