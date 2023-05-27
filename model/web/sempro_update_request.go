package web

type SemproUpdateRequest struct {
	Id_sempro    string `validate:"required" json:"id_sempro"`
	Id_mhs       string `validate:"required" json:"id_mhs"`
	Tanggal      string `validate:"required" json:"tanggal"`
	Ruang_sempro string `validate:"required" json:"ruang_sempro"`
	Status       string `validate:"required" json:"status"`
}
