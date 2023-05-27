package service

import (
	"context"
	"goelster/model/web"
)

type PengajuanBimbinganService interface {
	PengajuanBimbingan(ctx context.Context, request web.PengajuanBimbinganRequest) web.PengajuanBimbinganResponse
	UpdateDataAjuanBimbingan(ctx context.Context, request web.PengajuanBimbinganUpdateRequest) web.PengajuanBimbinganResponse
	FindById(ctx context.Context, dataajuanId string) []web.PengajuanBimbinganResponse
	FindByIdDosen(ctx context.Context, ajuanDosenId string) []web.PengajuanBimbinganResponse
	HasilBim(ctx context.Context, bim web.TambahHasilBimRequest) web.PengajuanBimbinganResponse
}
