package web

type UjianTaResponse struct {
	Id_mhs        string `json:"id_mhs"`
	Id_staf       string `json:"id_staf"`
	Nama_judul    string `json:"nama_judul"`
	Abstrak       string `json:"abstrak"`
	Tanggal       string `json:"tanggal"`
	Ruangan       string `json:"ruangan"`
	Proposalakhir string `json:"proposalakhir"`
	Status_ajuan  string `json:"status_ajuan"`
}
