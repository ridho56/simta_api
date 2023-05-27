package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"goelster/exception"
	"goelster/helper"
	"goelster/model/domain"
	"goelster/model/web"
	"goelster/repository"
	"time"
)

type PengajuanSeminarHasilServiceImpl struct {
	PengajuanSeminarHasilRepository repository.PengajuanSeminarHasilRepository
	DB                              *sql.DB
	Validate                        *validator.Validate
}

func NewPengajuanSeminarHasilService(pengajuanSeminarHasilRepository repository.PengajuanSeminarHasilRepository, DB *sql.DB, validate *validator.Validate) *PengajuanSeminarHasilServiceImpl {
	return &PengajuanSeminarHasilServiceImpl{
		PengajuanSeminarHasilRepository: pengajuanSeminarHasilRepository,
		DB:                              DB,
		Validate:                        validate}
}

func (service *PengajuanSeminarHasilServiceImpl) UplodFile(ctx context.Context, request web.SeminarHasilUploadRequest) web.SeminarHasilResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	newUUID := uuid.New().String()

	scheduledTimeEpoch := time.Now().UnixNano() / int64(time.Millisecond)

	seminarhasil := domain.SeminarHasil{
		Id_seminarhasil: newUUID,

		Link_file:  request.Link_file,
		Created_at: string(scheduledTimeEpoch),
	}
	seminarhasil = service.PengajuanSeminarHasilRepository.UploadFile(ctx, tx, seminarhasil)
	return helper.ToSeeminarHasilResponse(seminarhasil)
}

func (service *PengajuanSeminarHasilServiceImpl) GetFile(ctx context.Context) []web.SeminarHasilResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hasil, err := service.PengajuanSeminarHasilRepository.GetFile(ctx, tx)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToNewSeminarHasilResponse(hasil)
}

func (service *PengajuanSeminarHasilServiceImpl) PengajuanJadwalSemhas(ctx context.Context, request web.SeminarHasilRequest) web.SeminarHasilResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	newUUID := uuid.New().String()

	jadwalseminarhasil := domain.SeminarHasil{
		Id_seminarhasil: newUUID,
		Id_mhs:          request.Id_mhs,
		Id_staf:         request.Id_staf,
		Nama_judul:      request.Nama_judul,
		Abstak:          request.Abstrak,
		Jadwal_semhas:   request.Jadwal_semhas,
		Link_file:       request.Link_file,
		Status_ajuan:    "pending",
	}
	jadwalseminarhasil = service.PengajuanSeminarHasilRepository.PengajuanJadwalSemhas(ctx, tx, jadwalseminarhasil)
	return helper.ToSeeminarHasilResponse(jadwalseminarhasil)
}

func (service *PengajuanSeminarHasilServiceImpl) UpdateAjuanSemhas(ctx context.Context, request web.SeminarHasilUpdateRequest) web.SeminarHasilResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	updatesemhas := domain.SeminarHasil{
		Id_seminarhasil: request.Id_semhas,
		Jadwal_semhas:   request.Jadwal_semhas,
		Ruang_semhas:    request.Ruang_semhas,
		Status_ajuan:    "Diterima",
	}

	updatesemhas = service.PengajuanSeminarHasilRepository.UpdateAjuanSemhas(ctx, tx, updatesemhas)

	return helper.ToSeeminarHasilResponse(updatesemhas)
}

func (service *PengajuanSeminarHasilServiceImpl) FindById(ctx context.Context, dataasemhasId string) []web.SeminarHasilResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	datasemhas, err := service.PengajuanSeminarHasilRepository.FindById(ctx, tx, dataasemhasId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToNewSeminarHasilResponse(datasemhas)
}
