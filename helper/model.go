package helper

import (
	"goelster/model/domain"
	"goelster/model/web"
)

//	func ToRegisterResponse(user domain.User) web.RegisterResponse {
//		return web.RegisterResponse{
//			Email:  user.Email,
//			Status: "Berhasil",
//		}
//	}
func ToLoginResponse(user domain.User) web.LoginResponse {
	name := GetStafname(user.Id_staf)
	return web.LoginResponse{
		Id:              user.Id,
		Id_mhs:          user.Id_mhs,
		Id_staf:         user.Id_staf,
		NamaIDDosen:     name,
		EmailOrUsername: user.Email,
		Status:          "login Berhasil",
	}
}

//func ToEmailCheckResponse(string2 string) web.EmailCheckResponse {
//	return web.EmailCheckResponse{Status: string2}
//}

func ToDospemResponse(dospem domain.SetDospem) web.DospemResponse {
	return web.DospemResponse{
		Id_mhs:  dospem.Id_mhs,
		Id_staf: dospem.Id_staf,
	}
}

func ToNewDosenResponse(dosen []domain.NamaIDDosen) []web.NamaIDDosen {
	var newdosen []web.NamaIDDosen
	for _, dosen := range dosen {
		newdosen = append(newdosen, ToDosenResponse(dosen))
	}
	return newdosen
}

func ToMhsResponse(namaid domain.NamaIDMahasiswa) web.NamaIDMhs {
	return web.NamaIDMhs{
		ID:   namaid.ID,
		Nama: namaid.Nama,
	}
}

func ToNewMhsResponse(mhs []domain.NamaIDMahasiswa) []web.NamaIDMhs {
	var newmhs []web.NamaIDMhs
	for _, mhs := range mhs {
		newmhs = append(newmhs, ToMhsResponse(mhs))
	}
	return newmhs
}

func ToDosenResponse(namaid domain.NamaIDDosen) web.NamaIDDosen {
	return web.NamaIDDosen{
		ID:   namaid.ID,
		Nama: namaid.Nama,
	}
}

func ToNewDospemResponse(dospem []domain.SetDospem) []web.DospemResponse {
	var newResponse []web.DospemResponse
	for _, dospem := range dospem {
		newResponse = append(newResponse, ToDospemResponse(dospem))
	}
	return newResponse
}

func ToBimbinganResponse(bim domain.Bimbingan) web.PengajuanBimbinganResponse {
	nama := GetMahasiswaname(bim.Id_mhs)
	return web.PengajuanBimbinganResponse{
		Id_mhs:       bim.Id_mhs,
		Id_staf:      bim.Id_staf,
		NamaMhs:      nama,
		JadwalBim:    bim.JadwalBim,
		RuangBim:     bim.RuangBim,
		HasilBim:     bim.HasilBim,
		Status_ajuan: bim.Status,
	}
}

func ToNewBimbinganResponse(bimbingan []domain.Bimbingan) []web.PengajuanBimbinganResponse {
	var newResponse []web.PengajuanBimbinganResponse
	for _, bimbigan := range bimbingan {
		newResponse = append(newResponse, ToBimbinganResponse(bimbigan))
	}
	return newResponse
}

func ToSemproResponse(sempro domain.Sempro) web.SemproResponse {
	return web.SemproResponse{
		Id_mhs:       sempro.Id_mhs,
		Id_staf:      sempro.Id_staf,
		Judul_ta:     sempro.Judul_ta,
		Abstract:     sempro.Abstract,
		Tanggal:      sempro.Tanggal,
		Ruang_sempro: sempro.Ruang_sempro,
		Proposal:     sempro.Proposal,
		Status_ajuan: sempro.Status_ajuan,
	}
}

func ToNewSemproResponse(sempro []domain.Sempro) []web.SemproResponse {
	var newsempro []web.SemproResponse
	for _, sempro := range sempro {
		newsempro = append(newsempro, ToSemproResponse(sempro))
	}
	return newsempro
}

func ToSemproResponses(sempros []domain.Sempro) []web.SemproResponse {
	var responses []web.SemproResponse
	for _, sempro := range sempros {
		response := web.SemproResponse{
			Tanggal:      sempro.Tanggal,
			Ruang_sempro: sempro.Ruang_sempro,
			Status_ajuan: sempro.Status_ajuan,
		}
		responses = append(responses, response)
	}
	return responses
}

func ToSeeminarHasilResponse(hasil domain.SeminarHasil) web.SeminarHasilResponse {
	return web.SeminarHasilResponse{
		Id_mhs:        hasil.Id_mhs,
		Id_staf:       hasil.Id_staf,
		Nama_judul:    hasil.Nama_judul,
		Abstrak:       hasil.Abstak,
		Jadwal_semhas: hasil.Jadwal_semhas,
		Ruang_semhas:  hasil.Ruang_semhas,
		Link_file:     hasil.Link_file,
		Status_ajuan:  hasil.Status_ajuan,
	}
}

func ToNewSeminarHasilResponse(seminar []domain.SeminarHasil) []web.SeminarHasilResponse {
	var newseminar []web.SeminarHasilResponse
	for _, seminar := range seminar {
		newseminar = append(newseminar, ToSeeminarHasilResponse(seminar))
	}
	return newseminar
}

func ToUjianTaResponse(ujian domain.UjianTa) web.UjianTaResponse {
	return web.UjianTaResponse{
		Id_mhs:        ujian.Id_mhs,
		Id_staf:       ujian.Id_staf,
		Nama_judul:    ujian.Nama_judul,
		Abstrak:       ujian.Abstrak,
		Tanggal:       ujian.Tanggal,
		Ruangan:       ujian.Ruangan,
		Proposalakhir: ujian.Proposalakhir,
		Status_ajuan:  ujian.Status_ajuan,
	}
}

func ToNewUjianTaResponse(ujian []domain.UjianTa) []web.UjianTaResponse {
	var newujianta []web.UjianTaResponse
	for _, ujian := range ujian {
		newujianta = append(newujianta, ToUjianTaResponse(ujian))
	}
	return newujianta
}
