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
	"strconv"
	"time"
)

type PengajuanSemproServiceImpl struct {
	PengajuanSemproRepository repository.PengajuanSemproRepository
	DB                        *sql.DB
	Validate                  *validator.Validate
}

func NewPengajuanSemproService(pengajuansemproRepository repository.PengajuanSemproRepository, DB *sql.DB, validate *validator.Validate) *PengajuanSemproServiceImpl {
	return &PengajuanSemproServiceImpl{
		PengajuanSemproRepository: pengajuansemproRepository,
		DB:                        DB,
		Validate:                  validate}
}

func (service *PengajuanSemproServiceImpl) PengajuanSempro(ctx context.Context, request web.SemproRequest) web.SemproResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	newUUID := uuid.New().String()

	scheduledTimeEpoch := time.Now().UnixNano() / int64(time.Millisecond)
	scheduledTimeString := strconv.FormatInt(scheduledTimeEpoch, 10)

	sempro := domain.Sempro{
		Id_sempro:    newUUID,
		Id_mhs:       request.Id_mhs,
		Id_staf:      request.Id_staf,
		Judul_ta:     request.Judul_ta,
		Abstract:     request.Abstract,
		Tanggal:      request.Tanggal,
		Proposal:     request.Proposal,
		Status_ajuan: "pending",
		Created_at:   scheduledTimeString,
	}
	sempro = service.PengajuanSemproRepository.PengajuanSempro(ctx, tx, sempro)
	return helper.ToSemproResponse(sempro)
}

func (service *PengajuanSemproServiceImpl) UpdateAjuanSempro(ctx context.Context, request []web.SemproUpdateRequest) []web.SemproResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var updatedSempros []domain.Sempro

	for _, request := range request {
		err := service.Validate.Struct(request)
		helper.PanicIfError(err)

		updatesempro := domain.Sempro{
			Id_sempro:    request.Id_sempro,
			Id_mhs:       request.Id_mhs,
			Tanggal:      request.Tanggal,
			Ruang_sempro: request.Ruang_sempro,
			Status_ajuan: request.Status,
		}

		updatesempro = service.PengajuanSemproRepository.UpdateAjuanSempro(ctx, tx, updatesempro)

		updatedSempros = append(updatedSempros, updatesempro)
	}

	return helper.ToSemproResponses(updatedSempros)
}

func (service *PengajuanSemproServiceImpl) FindById(ctx context.Context, dataasemproId string) []web.SemproResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	datasempro, err := service.PengajuanSemproRepository.FindById(ctx, tx, dataasemproId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToNewSemproResponse(datasempro)
}
