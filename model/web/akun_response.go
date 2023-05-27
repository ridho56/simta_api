package web

//type RegisterResponse struct {
//	Email  string `json:"email"`
//	Status string `json:"status"`
//}

type LoginResponse struct {
	Id              int    `json:"id"`
	Id_mhs          string `json:"id_mhs"`
	Id_staf         string `json:"id_staf"`
	NamaIDDosen     string `json:"nama_id_dosen"`
	EmailOrUsername string `json:"email_or_username"`
	Status          string `json:"status"`
}

//type EmailCheckResponse struct {
//	Status string `json:"status"`
//}
