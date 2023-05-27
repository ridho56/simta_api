package web

type DospemResponse struct {
	Id_mhs  string `validate:"required" json:"id_mhs"`
	Id_staf string `validate:"required" json:"id_staf"`
	Status  string `validate:"required" json:"status"`
}

type NamaIDMhs struct {
	ID   string `validate:"required" json:"id"`
	Nama string `validate:"required" json:"nama"`
}

type NamaIDDosen struct {
	ID   string `validate:"required" json:"id"`
	Nama string `validate:"required" json:"nama"`
}
