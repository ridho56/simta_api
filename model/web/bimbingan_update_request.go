package web

type PengajuanBimbinganUpdateRequest struct {
	Id_bimbingan int    `validate:"required" json:"id_bimbingan"`
	JadwalBim    string `validate:"required" json:"jadwal_bim"`
	RuangBim     string `validate:"required" json:"ruang_bim"`
	Status       string `validate:"required" json:"status"`
}
