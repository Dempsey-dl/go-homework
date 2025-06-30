package service

import (
	"blog/internal/model"
	"blog/internal/repoisitory"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	UserRepo  *repoisitory.UserRepoisitory
	JwtSCrept string
	JwtExp    time.Duration
}

func NewUserService(db *gorm.DB, jwtScrept string) *UserService {
	return &UserService{
		UserRepo:  repoisitory.NewUserRepository(db),
		JwtSCrept: jwtScrept,
	}
}

func (s *UserService) Register(input *model.RegisterUser) (*model.User, error) {
	//用户名是否重复

	if _, err := s.UserRepo.FindByName(input.Username); err == nil {
		return nil, errors.New("username aready existy")
	}
	//邮箱是否重复
	if _, err := s.UserRepo.FindByEmail(input.Email); err == nil {
		return nil, errors.New("email aready existy")
	}
	//加密
	hashedpwd, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	//赋值
	user := model.User{
		Username: input.Username,
		Password: string(hashedpwd),
		Email:    input.Email,
	}

	//创建
	if err := s.UserRepo.CreateUser(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserService) Login(input *model.LoginUser) (string, *model.User, error) {
	//查找用户
	user, err := s.UserRepo.FindByName(input.Username)
	if err != nil {
		return "", nil, err
	}
	//比较加密密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return "", nil, err
	}

	//生成JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	fmt.Println(s.JwtSCrept)
	tokenstr, err := token.SignedString([]byte(s.JwtSCrept))
	if err != nil {
		return "", nil, err
	}
	fmt.Println(s.JwtSCrept)
	return tokenstr, user, nil
}
