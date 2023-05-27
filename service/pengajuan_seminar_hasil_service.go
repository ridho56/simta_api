package service

import (
	"context"
	"goelster/model/web"
)

type PengajuanSeminarHasilService interface {
	UplodFile(ctx context.Context, request web.SeminarHasilUploadRequest) web.SeminarHasilResponse
	GetFile(ctx context.Context) []web.SeminarHasilResponse
	PengajuanJadwalSemhas(ctx context.Context, request web.SeminarHasilRequest) web.SeminarHasilResponse
	UpdateAjuanSemhas(ctx context.Context, request web.SeminarHasilUpdateRequest) web.SeminarHasilResponse
	FindById(ctx context.Context, dataasemhasId string) []web.SeminarHasilResponse
}
