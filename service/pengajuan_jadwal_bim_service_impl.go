package service

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"goelster/exception"
	"goelster/helper"
	"goelster/model/domain"
	"goelster/model/web"
	"goelster/repository"
)

type PengajuanBimbinganServiceImpl struct {
	PengajuanBimbinganRepository repository.PengajuanBimbinganRepository
	DB                           *sql.DB
	Validate                     *validator.Validate
}

func NewPegajuanBimbinganService(pengajuanbimbinganRepository repository.PengajuanBimbinganRepository, DB *sql.DB, validate *validator.Validate) *PengajuanBimbinganServiceImpl {
	return &PengajuanBimbinganServiceImpl{
		PengajuanBimbinganRepository: pengajuanbimbinganRepository,
		DB:                           DB,
		Validate:                     validate}
}

func (service *PengajuanBimbinganServiceImpl) PengajuanBimbingan(ctx context.Context, request web.PengajuanBimbinganRequest) web.PengajuanBimbinganResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	bimbingan := domain.Bimbingan{
		Id_mhs:    request.Id_mhs,
		Id_staf:   request.Id_staf,
		JadwalBim: request.JadwalBim,
		RuangBim:  "",
		Status:    "pending",
	}
	bimbingan = service.PengajuanBimbinganRepository.PengajuanBimbingan(ctx, tx, bimbingan)
	return helper.ToBimbinganResponse(bimbingan)
}

func (service *PengajuanBimbinganServiceImpl) UpdateDataAjuanBimbingan(ctx context.Context, request web.PengajuanBimbinganUpdateRequest) web.PengajuanBimbinganResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	updatebim := domain.Bimbingan{
		Id_bimbingan: request.Id_bimbingan,
		JadwalBim:    request.JadwalBim,
		RuangBim:     request.RuangBim,
		Status:       request.Status,
	}

	updatebim = service.PengajuanBimbinganRepository.UpdateDataAjuanBimbingan(ctx, tx, updatebim)

	return helper.ToBimbinganResponse(updatebim)
}

func (service *PengajuanBimbinganServiceImpl) FindById(ctx context.Context, dataajuanId string) []web.PengajuanBimbinganResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dataajuan, err := service.PengajuanBimbinganRepository.FindById(ctx, tx, dataajuanId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToNewBimbinganResponse(dataajuan)
}

func (service *PengajuanBimbinganServiceImpl) FindByIdDosen(ctx context.Context, ajuanDosenId string) []web.PengajuanBimbinganResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	ajuanDosen, err := service.PengajuanBimbinganRepository.FindByIdDosen(ctx, tx, ajuanDosenId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToNewBimbinganResponse(ajuanDosen)
}

func (service *PengajuanBimbinganServiceImpl) HasilBim(ctx context.Context, bim web.TambahHasilBimRequest) web.PengajuanBimbinganResponse {
	err := service.Validate.Struct(bim)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	hasil := domain.Bimbingan{
		Id_bimbingan: bim.Id_bimbingan,
		HasilBim:     bim.HasilBim,
	}

	hasil = service.PengajuanBimbinganRepository.HasilBim(ctx, tx, hasil)

	return helper.ToBimbinganResponse(hasil)
}
