package response

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
)

type FindUserByIdResponse struct {
	Id                   string  `json:"id"`
	FullName             string  `json:"full_name"`
	Username             string  `json:"username"`
	Email                string  `json:"email"`
	Phone                string  `json:"phone"`
	ReferalCode          string  `json:"referal_code"`
	BalancePoints        float64 `json:"balance_points"`
	ReferalCodeUsedCount int     `json:"referal_code_used_count"`
}

func ToUserFindByIdResponse(user entity.User, userCount int) (userResponse FindUserByIdResponse) {
	userResponse.Id = user.Id
	userResponse.Username = user.Username
	userResponse.FullName = user.FamilyMembers.FullName
	userResponse.Email = user.FamilyMembers.Email
	userResponse.Phone = user.FamilyMembers.Phone
	userResponse.ReferalCode = user.ReferalCode
	userResponse.BalancePoints = user.BalancePoint.BalancePoints
	userResponse.ReferalCodeUsedCount = userCount
	return userResponse
}
