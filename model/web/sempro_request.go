package web

type SemproRequest struct {
	Id_mhs   string `validate:"required" json:"id_mhs"`
	Id_staf  string `validate:"required" json:"id_staf"`
	Judul_ta string `validate:"required" json:"judul_ta"`
	Abstract string `validate:"required" json:"abstract"`
	Tanggal  string `validate:"required" json:"tanggal"`
	Proposal string `validate:"required" json:"proposal"`
}
