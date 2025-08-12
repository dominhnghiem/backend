package dto

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
