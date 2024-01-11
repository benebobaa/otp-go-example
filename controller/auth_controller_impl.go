package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sent-email-otp/helper"
	"sent-email-otp/model/web"
	"sent-email-otp/model/web/register"
	"sent-email-otp/service"
)

type AuthControllerImpl struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &AuthControllerImpl{
		AuthService: authService,
	}
}

func (a *AuthControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRegisRequest := register.RegisterRequest{}

	helper.ReadFromRequestBody(request, &userRegisRequest)

	userResponse := a.AuthService.Register(request.Context(), userRegisRequest)

	webResponse := web.BaseResponse{
		Code:   200,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
