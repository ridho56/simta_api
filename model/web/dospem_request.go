package web

type DospemRequest struct {
	Id_mhs  string `validate:"required" json:"id_mhs"`
	Id_staf string `validate:"required" json:"id_staf"`
}
