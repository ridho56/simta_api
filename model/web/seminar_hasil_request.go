package web

type SeminarHasilRequest struct {
	Id_mhs        string `validate:"required" json:"id_mhs"`
	Id_staf       string `validate:"required" json:"id_staf"`
	Nama_judul    string `validate:"required" json:"Nama_judul"`
	Abstrak       string `validate:"required" json:"abstrak"`
	Jadwal_semhas string `validate:"required" json:"jadwal_semhas"`
	Link_file     string `validate:"required" json:"link_file"`
}

type SeminarHasilUpdateRequest struct {
	Id_semhas     string `validate:"required" json:"id_semhas"`
	Jadwal_semhas string `validate:"required" json:"jadwal_semhas"`
	Ruang_semhas  string `validate:"required" json:"ruang_semhas"`
}

type SeminarHasilUploadRequest struct {
	Nama_file string `validate:"required" json:"nama_file"`
	Link_file string `validate:"required" json:"link_file"`
}
