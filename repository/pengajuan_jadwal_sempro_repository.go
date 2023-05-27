package repository

import (
	"context"
	"database/sql"
	"goelster/model/domain"
)

type PengajuanSemproRepository interface {
	PengajuanSempro(ctx context.Context, tx *sql.Tx, sempro domain.Sempro) domain.Sempro
	UpdateAjuanSempro(ctx context.Context, tx *sql.Tx, sempro domain.Sempro) domain.Sempro
	FindById(ctx context.Context, tx *sql.Tx, id_mhs string) ([]domain.Sempro, error)
}
