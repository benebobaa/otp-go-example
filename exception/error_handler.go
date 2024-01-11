package exception

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"sent-email-otp/helper"
	"sent-email-otp/model/web"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, exception interface{}) {

	if validationErrors(writer, request, exception) {
		return

	}

	if otpInvalidErrors(writer, request, exception) {
		return
	}

	internalServerError(writer, request, exception)
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		BaseResponse := web.BaseResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}

		helper.WriteToResponseBody(writer, BaseResponse)
		return true
	} else {
		return false
	}
}

func otpInvalidErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(OtpError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		BaseResponse := web.BaseResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error,
		}

		helper.WriteToResponseBody(writer, BaseResponse)
		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	BaseResponse := web.BaseResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	helper.WriteToResponseBody(writer, BaseResponse)
}
