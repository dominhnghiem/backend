package dto

type SignupRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=64"`
	Name     string `json:"name" validate:"omitempty,max=100"`
}

type SignupResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}
