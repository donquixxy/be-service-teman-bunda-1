package response

type AuthResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	VerApp       string `json:"ver_app"`
}

func ToAuthResponse(id string, username string, token string, refreshToken string, verApp string) (authResponse AuthResponse) {
	authResponse.Token = token
	authResponse.RefreshToken = refreshToken
	authResponse.VerApp = verApp
	return authResponse
}
