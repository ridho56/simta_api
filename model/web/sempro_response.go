package web

type SemproResponse struct {
	Id_mhs       string `json:"id_mhs"`
	Id_staf      string `json:"id_staf"`
	Judul_ta     string `json:"judul_ta"`
	Abstract     string `json:"abstract"`
	Tanggal      string `json:"tanggal"`
	Ruang_sempro string `json:"ruang_sempro"`
	Proposal     string `json:"proposal"`
	Status_ajuan string `json:"status_ajuan"`
}
