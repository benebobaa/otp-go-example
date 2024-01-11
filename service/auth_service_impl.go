package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"sent-email-otp/helper"
	"sent-email-otp/model/domain"
	"sent-email-otp/model/web/register"
	"sent-email-otp/model/web/user"
	"sent-email-otp/repository"
	"sent-email-otp/repository/otp"
	"sent-email-otp/utils"
	"strconv"
	"time"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	OtpRepository  otp.OtpRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthService(authRepository repository.AuthRepository, otpRepository otp.OtpRepository, db *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AuthRepository: authRepository,
		OtpRepository:  otpRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (a *AuthServiceImpl) Login(ctx context.Context, request user.UserRequest) user.UserResponse {
	//TODO implement me
	panic("implement me")
}

func (a *AuthServiceImpl) Register(ctx context.Context, request register.RegisterRequest) register.RegisterResponse {
	err := a.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := a.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	userRegisData := domain.UserRegister{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	result := a.AuthRepository.Register(ctx, tx, userRegisData)

	//convert int to nullint
	resultNulId := sql.NullInt32{Int32: int32(result.Id), Valid: true}

	otpData := domain.Otp{
		RefCode:        utils.RandomCombineIntAndString(),
		OtpValue:       strconv.FormatInt(utils.RandomInt(100000, 999999), 10),
		ExpirationTime: time.Now().UTC().Add(time.Minute * 1),
		UserRegisId:    resultNulId,
	}
	resultOTP := a.OtpRepository.Create(ctx, tx, otpData)

	return helper.ToRegisterResponse(resultOTP)
}
