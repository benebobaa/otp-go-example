package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"sent-email-otp/helper"
	"sent-email-otp/model/domain"
	"sent-email-otp/model/web/user"
	"sent-email-otp/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, request user.UserRequest) user.UserResponse {
	err := u.Validate.Struct(request)
	helper.PanicIfError(err)
	tx := u.DB
	userData := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	result := u.UserRepository.CreateUser(ctx, tx, userData)

	return helper.ToUserResponse(result)
}
