package web

//	type RegisterRequest struct {
//		Nama     string `validate:"required" json:"nama"`
//		Email    string `validate:"required" json:"email"`
//		Password string `validate:"required" json:"password"`
//		Level    string `validate:"required" json:"level"`
//	}
type LoginRequest struct {
	EmailOrUsername string `validate:"required" json:"email_or_username"`
	Password        string `validate:"required" json:"password"`
}
