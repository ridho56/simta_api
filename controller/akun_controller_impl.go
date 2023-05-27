package controller

import (
	"github.com/julienschmidt/httprouter"
	"goelster/helper"
	"goelster/model/web"
	"goelster/service"
	"net/http"
	"time"
)

type AkunControllerImpl struct {
	AkunService service.AkunService
}

func NewAkunController(registerService service.AkunService) *AkunControllerImpl {
	return &AkunControllerImpl{
		AkunService: registerService,
	}
}

//func (controller *AkunControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
//	registerRequest := web.RegisterRequest{}
//	helper.ReadFromRequestBody(request, &registerRequest)
//
//	registerResponse := controller.AkunService.Register(request.Context(), registerRequest)
//	webResponse := web.WebResponse{
//		Code:   200,
//		Status: "OK",
//		Data:   registerResponse,
//	}
//	helper.WriteToResponseBody(writer, webResponse)
//}

func (controller *AkunControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web.LoginRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	loginResponse := controller.AkunService.Login(request.Context(), loginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   loginResponse,
	}
	jwt := helper.CreateNewJWT(loginResponse.EmailOrUsername, loginResponse.EmailOrUsername, "secret-key", time.Hour*48)
	helper.WriteToResponseBodyWithJWT(writer, webResponse, jwt)
}

//func (controller *AkunControllerImpl) EmailCheck(writer http.ResponseWriter, r *http.Request, params httprouter.Params) {
//	userEmail := params.ByName("emailUser")
//	EmailCheckResponse := controller.AkunService.EmailCheck(r.Context(), userEmail)
//
//	WebResponse := web.WebResponse{
//		Code:   200,
//		Status: "OK",
//		Data:   EmailCheckResponse,
//	}
//	helper.WriteToResponseBody(writer, WebResponse)
//}
