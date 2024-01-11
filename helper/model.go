package helper

import (
	"sent-email-otp/model/domain"
	"sent-email-otp/model/web/otp"
	"sent-email-otp/model/web/register"
	"sent-email-otp/model/web/user"
)

func ToUserResponse(userData domain.User) user.UserResponse {
	return user.UserResponse{
		Id:        userData.Id,
		Name:      userData.Name,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}
}

func ToOtpResponse(otpData domain.Otp) otp.OtpResponse {
	return otp.OtpResponse{
		RefCode: otpData.RefCode,
	}
}

func ToRegisterResponse(userRegistData domain.Otp) register.RegisterResponse {
	return register.RegisterResponse{
		RefCode: userRegistData.RefCode,
	}

}
