package repository

import (
	"context"
	"database/sql"
	"errors"
	"goelster/helper"
	"goelster/model/domain"
)

type PengajuanBimbinganRepositoryImpl struct{}

func NewPengajuanBimbinganRepository() *PengajuanBimbinganRepositoryImpl {
	return &PengajuanBimbinganRepositoryImpl{}
}

func (repository *PengajuanBimbinganRepositoryImpl) PengajuanBimbingan(ctx context.Context, tx *sql.Tx, bim domain.Bimbingan) domain.Bimbingan {
	query := "INSERT INTO simta_pengajuanbimbingan (id_staf, id_mhs, jadwal_bimbingan, ruang_bimbingan,  status_ajuan) VALUES(?,?,?,?,?);"

	res, err := tx.ExecContext(ctx, query, bim.Id_staf, bim.Id_mhs, bim.JadwalBim, bim.RuangBim, bim.Status)

	helper.PanicIfError(err)

	id, err := res.LastInsertId()
	helper.PanicIfError(err)

	bim.Id_bimbingan = int(id)
	return bim
}

func (repository *PengajuanBimbinganRepositoryImpl) UpdateDataAjuanBimbingan(ctx context.Context, tx *sql.Tx, bimbingan domain.Bimbingan) domain.Bimbingan {
	query := "UPDATE simta_pengajuanbimbingan SET  jadwal_bimbingan=?, ruang_bimbingan=?, status=? WHERE id_bimbingan=?;"
	_, err := tx.ExecContext(ctx, query, bimbingan.JadwalBim, bimbingan.RuangBim, bimbingan.Status, bimbingan.Id_bimbingan)
	helper.PanicIfError(err)
	return bimbingan
}

func (repository *PengajuanBimbinganRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, id_mhs string) ([]domain.Bimbingan, error) {
	query := "SELECT id_bimbingan, id_staf, id_mhs, jadwal_bimbingan, ruang_bimbingan, status_ajuan FROM simta_pengajuanbimbingan WHERE id_mhs = ?;"
	rows, err := tx.QueryContext(ctx, query, id_mhs)
	helper.PanicIfError(err)
	defer rows.Close()

	var newajuan []domain.Bimbingan

	for rows.Next() {
		ajuan := domain.Bimbingan{}
		err := rows.Scan(&ajuan.Id_bimbingan, &ajuan.Id_staf, &ajuan.Id_mhs, &ajuan.JadwalBim, &ajuan.RuangBim, &ajuan.Status)
		helper.PanicIfError(err)
		newajuan = append(newajuan, ajuan)
	}
	var gagal error
	if newajuan == nil {
		gagal = errors.New("data not found")
	} else {
		gagal = nil
	}
	return newajuan, gagal
}

func (repository *PengajuanBimbinganRepositoryImpl) FindByIdDosen(ctx context.Context, tx *sql.Tx, id_dosen string) ([]domain.Bimbingan, error) {
	query := "SELECT id_bimbingan, id_staf, id_mhs, jadwal_bimbingan, ruang_bimbingan, status FROM simta_pengajuanbimbingan WHERE id_staf = ?;"
	rows, err := tx.QueryContext(ctx, query, id_dosen)
	helper.PanicIfError(err)
	defer rows.Close()

	bimbingans := make([]domain.Bimbingan, 0)
	for rows.Next() {
		bimbingan := domain.Bimbingan{}
		err := rows.Scan(&bimbingan.Id_bimbingan, &bimbingan.Id_staf, &bimbingan.Id_mhs, &bimbingan.JadwalBim, &bimbingan.RuangBim, &bimbingan.Status)
		helper.PanicIfError(err)
		bimbingans = append(bimbingans, bimbingan)
	}

	if len(bimbingans) == 0 {
		return nil, errors.New("data not found")
	}
	return bimbingans, nil
}

func (repository *PengajuanBimbinganRepositoryImpl) HasilBim(ctx context.Context, tx *sql.Tx, bimbingan domain.Bimbingan) domain.Bimbingan {
	query := "UPDATE simta_pengajuanbimbingan SET hasil_bimbingan = ? WHERE id_bimbingan = ?;"

	_, err := tx.ExecContext(ctx, query, bimbingan.HasilBim, bimbingan.Id_bimbingan)

	helper.PanicIfError(err)
	return bimbingan
}
