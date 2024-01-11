package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sent-email-otp/helper"
	"sent-email-otp/model/web"
	otp2 "sent-email-otp/model/web/otp"
	"sent-email-otp/service/otp"
)

type OtpControllerImpl struct {
	OtpService otp.OtpService
}

func NewOtpController(otpService otp.OtpService) OtpController {
	return &OtpControllerImpl{
		OtpService: otpService,
	}
}

func (o *OtpControllerImpl) FindAndValidate(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	otpValidateRequest := otp2.OtpValidateRequest{}
	helper.ReadFromRequestBody(request, &otpValidateRequest)

	o.OtpService.FindAndValidate(request.Context(), otpValidateRequest)

	webResponse := web.BaseResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
