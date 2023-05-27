package controller

import (
	"github.com/julienschmidt/httprouter"
	"goelster/helper"
	"goelster/model/web"
	"goelster/service"
	"net/http"
)

type PengajuanSemproControllerImpl struct {
	PengajuanSemproService service.PengajuanSemproService
}

func NewPengajuanSemproController(pengajuansemproService service.PengajuanSemproService) *PengajuanSemproControllerImpl {
	return &PengajuanSemproControllerImpl{
		PengajuanSemproService: pengajuansemproService,
	}
}

func (controller *PengajuanSemproControllerImpl) Pengajuansempro(writer http.ResponseWriter, r *http.Request, param httprouter.Params) {
	semproRequest := web.SemproRequest{}
	helper.ReadFromRequestBody(r, &semproRequest)

	semproResponse := controller.PengajuanSemproService.PengajuanSempro(r.Context(), semproRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   semproResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PengajuanSemproControllerImpl) UpdateAjuanSempro(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	semproAjuanUpdateRequest := []web.SemproUpdateRequest{}
	helper.ReadFromRequestBody(r, &semproAjuanUpdateRequest)

	dataAjuanUpdateResponse := controller.PengajuanSemproService.UpdateAjuanSempro(r.Context(), semproAjuanUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataAjuanUpdateResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *PengajuanSemproControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	datasemproId := params.ByName("Id")

	categoryResponse := controller.PengajuanSemproService.FindById(request.Context(), datasemproId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
