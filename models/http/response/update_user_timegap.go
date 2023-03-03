package response

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
)
type CreateUserTimegapEditResponse struct {
	IdUser					string `json:"id_user"`
	TimegapData          	string `json:"timegap_data"`
	Message					string `json:"message"`
}
func ToUserCreateUserTimeGapResponse(user entity.User) (userResponse CreateUserTimegapEditResponse) {
	userResponse.IdUser = user.Id
	userResponse.TimegapData = user.TimegapData
	return userResponse
}
