package service

import (
	"context"
	"sent-email-otp/model/web/user"
)

type UserService interface {
	CreateUser(ctx context.Context, request user.UserRequest) user.UserResponse
}
