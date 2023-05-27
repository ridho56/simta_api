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
)

type PengajuanUjianTaServiceImpl struct {
	PengajuanUjianTaRepository repository.PengajuanUjianTaRepository
	DB                         *sql.DB
	Validate                   *validator.Validate
}

func NewPengajuanUjianTaService(pengajuanUjianTaRepository repository.PengajuanUjianTaRepository, DB *sql.DB, validate *validator.Validate) *PengajuanUjianTaServiceImpl {
	return &PengajuanUjianTaServiceImpl{
		PengajuanUjianTaRepository: pengajuanUjianTaRepository,
		DB:                         DB,
		Validate:                   validate}
}

func (service *PengajuanUjianTaServiceImpl) PengajuanUjianTa(ctx context.Context, request web.UjianTaRequest) web.UjianTaResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	newUUID := uuid.New().String()

	jadwalujianta := domain.UjianTa{
		Id_ujianta:    newUUID,
		Id_mhs:        request.Id_mhs,
		Id_staf:       request.Id_staf,
		Nama_judul:    request.Nama_judul,
		Abstrak:       request.Abstrak,
		Ruangan:       "",
		Tanggal:       request.Tanggal,
		Proposalakhir: request.Proposalakhir,
		Status_ajuan:  "pending",
	}
	jadwalujianta = service.PengajuanUjianTaRepository.PengajuanUjianTa(ctx, tx, jadwalujianta)
	return helper.ToUjianTaResponse(jadwalujianta)
}

func (service *PengajuanUjianTaServiceImpl) FindById(ctx context.Context, dataujiantaId string) []web.UjianTaResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dataujianta, err := service.PengajuanUjianTaRepository.FindById(ctx, tx, dataujiantaId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToNewUjianTaResponse(dataujianta)
}
