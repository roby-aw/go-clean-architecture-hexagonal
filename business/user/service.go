package user

import (
	"errors"

	"github.com/roby-aw/go-clean-architecture-hexagonal/utils"

	"github.com/go-playground/validator/v10"
)

type Repository interface {
	FindUserByID(id string) (User, error)
	FindUserByEmail(email string) (User, error)
	CreateUser(data Register) error
}
type Service interface {
	Login(auth AuthLogin) (*ResponseLogin, error)
	RegisterUser(data Register) error
	GetUserByID(id string) (User, error)
}

type service struct {
	repository Repository
	validate   *validator.Validate
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
		validate:   validator.New(),
	}
}

func (s *service) Login(auth AuthLogin) (*ResponseLogin, error) {
	err := s.validate.Struct(&auth)
	if err != nil {
		err = utils.HandleErrorValidator(err)
		return nil, err
	}
	Admin, err := s.repository.FindUserByEmail(auth.Email)
	if err != nil {
		return nil, err
	}
	err = utils.VerifyPassword(Admin.Password, auth.Password)
	if err != nil {
		return nil, errors.New("wrong password")
	}

	exp, token, err := utils.GenerateAccessTokenUser(Admin.ID, Admin.Email)
	if err != nil {
		return nil, err
	}
	exprefresh, refreshtoken, err := utils.GenerateRefreshTokenUser(Admin.ID, Admin.Email)
	if err != nil {
		return nil, err
	}
	var restoken = utils.Token{
		AccessToken:         token,
		AccessTokenExpired:  exp,
		RefreshToken:        refreshtoken,
		RefreshTokenExpired: exprefresh,
	}
	return &ResponseLogin{
		Email: Admin.Email,
		Token: restoken,
	}, nil
}
func (s *service) RegisterUser(data Register) error {
	err := s.validate.Struct(&data)
	if err != nil {
		return err
	}
	err = s.repository.CreateUser(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetUserByID(id string) (User, error) {
	return s.repository.FindUserByID(id)
}
