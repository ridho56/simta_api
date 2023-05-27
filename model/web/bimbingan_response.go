package web

type PengajuanBimbinganResponse struct {
	Id_mhs       string `json:"id_mhs"`
	Id_staf      string `json:"id_Staf"`
	NamaMhs      string `json:"nama_mhs"`
	JadwalBim    string `json:"jadwal_bim"`
	RuangBim     string `json:"ruang_bim"`
	HasilBim     string `json:"hasil_bim"`
	Status_ajuan string `json:"status_ajuan"`
}
