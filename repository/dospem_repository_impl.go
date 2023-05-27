package repository

import (
	"context"
	"database/sql"
	"errors"
	"goelster/helper"
	"goelster/model/domain"
)

type DospemRepositoryImpl struct{}

func NewDospemRepository() *DospemRepositoryImpl {
	return &DospemRepositoryImpl{}
}

func (repository *DospemRepositoryImpl) SetDospem(ctx context.Context, tx *sql.Tx, dospem domain.SetDospem) domain.SetDospem {
	query := "INSERT INTO simta_dospem (id_staf, id_mhs) VALUES(?, ?);"

	res, err := tx.ExecContext(ctx, query, dospem.Id_staf, dospem.Id_mhs)

	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	dospem.Id_dospem = int(id)
	return dospem
}

func (repository *DospemRepositoryImpl) GetDospem(ctx context.Context, tx *sql.Tx, id_mhs string) ([]domain.SetDospem, error) {
	query := "SELECT id_dospem, id_staf FROM simta_dospem WHERE id_mhs = ?;"
	rows, err := tx.QueryContext(ctx, query, id_mhs)
	helper.PanicIfError(err)
	defer rows.Close()

	dospems := make([]domain.SetDospem, 0)
	for rows.Next() {
		dospem := domain.SetDospem{}
		err := rows.Scan(&dospem.Id_dospem, &dospem.Id_staf)
		helper.PanicIfError(err)
		dospems = append(dospems, dospem)
	}

	if len(dospems) == 0 {
		return nil, errors.New("data not found")
	}
	return dospems, nil
}

func (repository *DospemRepositoryImpl) GetMhs(ctx context.Context, tx *sql.Tx) ([]domain.NamaIDMahasiswa, error) {
	query := "SELECT id_mhs, nama FROM mahasiswa;"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []domain.NamaIDMahasiswa
	for rows.Next() {
		var result domain.NamaIDMahasiswa
		err := rows.Scan(&result.ID, &result.Nama)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func (repository *DospemRepositoryImpl) GetDosen(ctx context.Context, tx *sql.Tx) ([]domain.NamaIDDosen, error) {
	query := "SELECT id_staf, nama FROM staf;"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hasils []domain.NamaIDDosen
	for rows.Next() {
		var hasil domain.NamaIDDosen
		err := rows.Scan(&hasil.ID, &hasil.Nama)
		if err != nil {
			return nil, err
		}
		hasils = append(hasils, hasil)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return hasils, nil
}
