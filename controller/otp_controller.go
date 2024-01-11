package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type OtpController interface {
	FindAndValidate(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
