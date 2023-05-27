package controller

import (
	"github.com/julienschmidt/httprouter"
	"goelster/helper"
	"goelster/model/web"
	"goelster/service"
	"net/http"
)

type PengajuanBimbinganControllerImpl struct {
	PengajuanBimbinganService service.PengajuanBimbinganService
}

func NewPengajuanBimbinganController(pengajuanbimbinganService service.PengajuanBimbinganService) *PengajuanBimbinganControllerImpl {
	return &PengajuanBimbinganControllerImpl{
		PengajuanBimbinganService: pengajuanbimbinganService,
	}
}

func (controller *PengajuanBimbinganControllerImpl) PengajuanBimbingan(writer http.ResponseWriter, r *http.Request, param httprouter.Params) {
	pengajuanRequest := web.PengajuanBimbinganRequest{}
	helper.ReadFromRequestBody(r, &pengajuanRequest)

	pengajuanResponse := controller.PengajuanBimbinganService.PengajuanBimbingan(r.Context(), pengajuanRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   pengajuanResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PengajuanBimbinganControllerImpl) UpdateDataAjuanBimbingan(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	dataAjuanUpdateRequest := web.PengajuanBimbinganUpdateRequest{}
	helper.ReadFromRequestBody(r, &dataAjuanUpdateRequest)

	dataAjuanUpdateResponse := controller.PengajuanBimbinganService.UpdateDataAjuanBimbingan(r.Context(), dataAjuanUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataAjuanUpdateResponse,
	}
	helper.WriteToResponseBody(w, webResponse)
}

func (controller *PengajuanBimbinganControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	dataajuanId := params.ByName("Id")

	categoryResponse := controller.PengajuanBimbinganService.FindById(request.Context(), dataajuanId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PengajuanBimbinganControllerImpl) FindByIdDosen(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	ajuanDosenId := params.ByName("Id")

	categoryResponse := controller.PengajuanBimbinganService.FindByIdDosen(r.Context(), ajuanDosenId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *PengajuanBimbinganControllerImpl) Hasilbim(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	dataRequest := web.TambahHasilBimRequest{}
	helper.ReadFromRequestBody(request, &dataRequest)

	dataAjuanUpdateResponse := controller.PengajuanBimbinganService.HasilBim(request.Context(), dataRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   dataAjuanUpdateResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
