package web

type UjianTaRequest struct {
	Id_mhs        string `validate:"required" json:"id_mhs"`
	Id_staf       string `validate:"required" json:"id_staf"`
	Nama_judul    string `validate:"required" json:"nama_judul"`
	Abstrak       string `validate:"required" json:"abstrak"`
	Tanggal       string `validate:"required" json:"tanggal"`
	Proposalakhir string `validate:"required" json:"proposalakhir"`
}
