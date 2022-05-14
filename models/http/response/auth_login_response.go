package response

type AuthResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func ToAuthResponse(id string, username string, token string, refreshToken string, verApp string) (authResponse AuthResponse) {
	authResponse.Token = token
	authResponse.RefreshToken = refreshToken
	return authResponse
}
