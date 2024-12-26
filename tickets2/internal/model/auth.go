package model

type AccessTokenPayload struct {
	ID    string
	Email string
}

type VerifyUserRequest struct {
	AccessToken string `json:"access_token,omitempty"`
}
