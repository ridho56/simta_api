package repository

import (
	"context"
	"database/sql"
	"goelster/model/domain"
)

type DospemRepository interface {
	SetDospem(ctx context.Context, tx *sql.Tx, dospem domain.SetDospem) domain.SetDospem
	GetDospem(ctx context.Context, tx *sql.Tx, id_mhs string) ([]domain.SetDospem, error)
	GetMhs(ctx context.Context, tx *sql.Tx) ([]domain.NamaIDMahasiswa, error)
	GetDosen(ctx context.Context, tx *sql.Tx) ([]domain.NamaIDDosen, error)
}
