package response

type AuthResponse struct {
	Id           string `json:"id"`
	Username     string `json:"username"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

func ToAuthResponse(id string, username string, token string, refreshToken string) (authResponse AuthResponse) {
	authResponse.Id = id
	authResponse.Username = username
	authResponse.Token = token
	authResponse.RefreshToken = refreshToken
	return authResponse
}
