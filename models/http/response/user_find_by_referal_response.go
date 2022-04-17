package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type FindUserByReferalResponse struct {
	ReferalCode string `json:"referal_code"`
}

func ToUserFindByReferalResponse(user entity.User) (userResponse FindUserByReferalResponse) {
	userResponse.ReferalCode = user.ReferalCode
	return userResponse
}
