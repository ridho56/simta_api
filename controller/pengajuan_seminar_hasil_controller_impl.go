package controller

import (
	"github.com/julienschmidt/httprouter"
	"goelster/helper"
	"goelster/model/web"
	"goelster/service"
	"net/http"
)

type PengajuanSeminarHasilControllerImpl struct {
	PengajuanSeminarHasilService service.PengajuanSeminarHasilService
}

func NewPengajuanSeminarHasilController(hasilService service.PengajuanSeminarHasilService) *PengajuanSeminarHasilControllerImpl {
	return &PengajuanSeminarHasilControllerImpl{
		PengajuanSeminarHasilService: hasilService,
	}
}

func (controller *PengajuanSeminarHasilControllerImpl) UploadFile(writer http.ResponseWriter, r *http.Request, param httprouter.Params) {
	seminarHasilRequest := web.SeminarHasilUploadRequest{}
	helper.ReadFromRequestBody(r, &seminarHasilRequest)

	semproResponse := controller.PengajuanSeminarHasilService.UplodFile(r.Context(), seminarHasilRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   semproResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PengajuanSeminarHasilControllerImpl) GetFile(writer http.ResponseWriter, r *http.Request, param httprouter.Params) {
	getfileResponse := controller.PengajuanSeminarHasilService.GetFile(r.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   getfileResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PengajuanSeminarHasilControllerImpl) PengajuanJadwalSemhas(writer http.ResponseWriter, r *http.Request, param httprouter.Params) {
	jadwalseminarHasilRequest := web.SeminarHasilRequest{}
	helper.ReadFromRequestBody(r, &jadwalseminarHasilRequest)

	semproResponse := controller.PengajuanSeminarHasilService.PengajuanJadwalSemhas(r.Context(), jadwalseminarHasilRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   semproResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PengajuanSeminarHasilControllerImpl) UpdateAjuanSemhas(writer http.ResponseWriter, r *http.Request, param httprouter.Params) {
	updatejadwalseminarHasilRequest := web.SeminarHasilUpdateRequest{}
	helper.ReadFromRequestBody(r, &updatejadwalseminarHasilRequest)

	semproResponse := controller.PengajuanSeminarHasilService.UpdateAjuanSemhas(r.Context(), updatejadwalseminarHasilRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   semproResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PengajuanSeminarHasilControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	datasemhasId := params.ByName("Id")

	categoryResponse := controller.PengajuanSeminarHasilService.FindById(request.Context(), datasemhasId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
