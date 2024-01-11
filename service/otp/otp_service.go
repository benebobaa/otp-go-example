package otp

import (
	"context"
	"sent-email-otp/model/web/otp"
)

type OtpService interface {
	Create(ctx context.Context, request otp.OtpRequest) otp.OtpResponse
	FindAndValidate(ctx context.Context, request otp.OtpValidateRequest)
}
