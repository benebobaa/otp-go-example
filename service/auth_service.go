package service

import (
	"context"
	"sent-email-otp/model/web/register"
	"sent-email-otp/model/web/user"
)

type AuthService interface {
	Login(ctx context.Context, request user.UserRequest) user.UserResponse
	Register(ctx context.Context, request register.RegisterRequest) register.RegisterResponse
}
