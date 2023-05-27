package repository

import (
	"context"
	"database/sql"
	"goelster/model/domain"
)

type PengajuanUjianTaRepository interface {
	PengajuanUjianTa(ctx context.Context, tx *sql.Tx, ujian domain.UjianTa) domain.UjianTa
	FindById(ctx context.Context, tx *sql.Tx, id_mhs string) ([]domain.UjianTa, error)
}
