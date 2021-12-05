package base

type SignupBody struct {
	Email string `json:"email"`
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
