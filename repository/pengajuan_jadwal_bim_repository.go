package repository

import (
	"context"
	"database/sql"
	"goelster/model/domain"
)

type PengajuanBimbinganRepository interface {
	PengajuanBimbingan(ctx context.Context, tx *sql.Tx, ta domain.Bimbingan) domain.Bimbingan
	UpdateDataAjuanBimbingan(ctx context.Context, tx *sql.Tx, bimbingan domain.Bimbingan) domain.Bimbingan
	FindById(ctx context.Context, tx *sql.Tx, id_mhs string) ([]domain.Bimbingan, error)
	FindByIdDosen(ctx context.Context, tx *sql.Tx, id_dosen string) ([]domain.Bimbingan, error)
	HasilBim(ctx context.Context, tx *sql.Tx, bimbingan domain.Bimbingan) domain.Bimbingan
}
