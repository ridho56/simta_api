package controller

import (
	"github.com/julienschmidt/httprouter"
	"goelster/helper"
	"goelster/model/web"
	"goelster/service"
	"net/http"
)

type PengajuanUjianTaControllerImpl struct {
	PengajuanUjianTaService service.PengajuanUjianTaService
}

func NewPengajuanUjianTaController(ujianService service.PengajuanUjianTaService) *PengajuanUjianTaControllerImpl {
	return &PengajuanUjianTaControllerImpl{
		PengajuanUjianTaService: ujianService,
	}
}

func (controller *PengajuanUjianTaControllerImpl) PengajuanUjianTa(writer http.ResponseWriter, r *http.Request, param httprouter.Params) {
	jadwalUjianTaRequest := web.UjianTaRequest{}
	helper.ReadFromRequestBody(r, &jadwalUjianTaRequest)

	semproResponse := controller.PengajuanUjianTaService.PengajuanUjianTa(r.Context(), jadwalUjianTaRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   semproResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PengajuanUjianTaControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	dataujiantaId := params.ByName("Id")

	categoryResponse := controller.PengajuanUjianTaService.FindById(request.Context(), dataujiantaId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
