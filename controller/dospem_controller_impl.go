package controller

import (
	"github.com/julienschmidt/httprouter"
	"goelster/helper"
	"goelster/model/web"
	"goelster/service"
	"net/http"
)

type DospemControllerImpl struct {
	DospemService service.DospemService
}

func NewDospemController(dospemService service.DospemService) *DospemControllerImpl {
	return &DospemControllerImpl{
		DospemService: dospemService,
	}
}

func (controller *DospemControllerImpl) SetDospem(writer http.ResponseWriter, r *http.Request, param httprouter.Params) {
	setDospemRequest := web.DospemRequest{}
	helper.ReadFromRequestBody(r, &setDospemRequest)

	setDospemResponse := controller.DospemService.SetDospem(r.Context(), setDospemRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   setDospemResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DospemControllerImpl) GetDospem(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	datadosen := params.ByName("Id")

	categoryResponse := controller.DospemService.GetDospem(request.Context(), datadosen)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DospemControllerImpl) GetMhs(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	setMhsResponse := controller.DospemService.GetMhs(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   setMhsResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *DospemControllerImpl) GetDosen(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	setDosenResponse := controller.DospemService.GetDosen(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   setDosenResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
