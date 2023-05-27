package web

type SeminarHasilResponse struct {
	Id_mhs        string `json:"id_mhs"`
	Id_staf       string `json:"id_staf"`
	Nama_judul    string `json:"nama_judul"`
	Abstrak       string `json:"abstrak"`
	Jadwal_semhas string `json:"jadwal_semhas"`
	Ruang_semhas  string `json:"Ruang_semhas"`
	Link_file     string `json:"link_file"`
	Status_ajuan  string `json:"status_ajuan"`
}
