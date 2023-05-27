package web

type PengajuanBimbinganRequest struct {
	Id_mhs    string `validate:"required" json:"id_mhs"`
	Id_staf   string `validate:"required" json:"id_Staf"`
	JadwalBim string `validate:"required" json:"jadwal_bim"`
}

type TambahHasilBimRequest struct {
	Id_bimbingan int    `validate:"required" json:"id_bimbingan"`
	HasilBim     string `validate:"required" json:"hasil_bim"`
}
