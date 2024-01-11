package otp

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"sent-email-otp/exception"
	"sent-email-otp/helper"
	"sent-email-otp/model/domain"
	"sent-email-otp/model/web/otp"
	otpRepo "sent-email-otp/repository/otp"
	"time"
)

type OtpServiceImpl struct {
	OtpRepository otpRepo.OtpRepository
	DB            *sql.DB
	Validate      *validator.Validate
}

func NewOtpService(otpRepository otpRepo.OtpRepository, DB *sql.DB, validate *validator.Validate) OtpService {
	return &OtpServiceImpl{
		OtpRepository: otpRepository,
		DB:            DB,
		Validate:      validate,
	}
}

func (o *OtpServiceImpl) Create(ctx context.Context, request otp.OtpRequest) otp.OtpResponse {
	err := o.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := o.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	newExpirationTime := time.Now().Add(time.Minute * 1)

	otpData := domain.Otp{
		RefCode:        request.RefCode,
		OtpValue:       request.OtpValue,
		ExpirationTime: newExpirationTime,
		UserRegisId:    request.UserRegisId,
	}
	result := o.OtpRepository.Create(ctx, tx, otpData)

	return helper.ToOtpResponse(result)
}

func (o *OtpServiceImpl) FindAndValidate(ctx context.Context, request otp.OtpValidateRequest) {
	err := o.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := o.DB.Begin()

	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	result, err := o.OtpRepository.FindByRefCode(ctx, tx, request.RefCode)

	if err != nil {
		fmt.Println("otp not found")
		panic(exception.NewOtpError(err.Error()))
	}

	if result.ExpirationTime.Before(time.Now()) {
		fmt.Println("otp expired")
		panic(exception.NewOtpError("OTP expired. Please request a new one."))
	}

	if result.OtpValue != request.OtpValue {
		fmt.Println("otp wrong")
		panic(exception.NewOtpError("Invalid OTP. Please enter a valid one."))
	}

}
