package response

import (
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
)

type CreateUserResponse struct {
	IdUser          string `json:"id_user"`
	IdFamilyMembers string `json:"id_family_members"`
	IdFamily        string `json:"id_family"`
	IdBalancePoint  string `json:"id_balance_point"`
}

func ToUserCreateUserResponse(user entity.User, family entity.Family, familyMembers entity.FamilyMembers, balancePoint entity.BalancePoint) (userResponse CreateUserResponse) {
	userResponse.IdUser = user.Id
	userResponse.IdFamilyMembers = familyMembers.Id
	userResponse.IdFamily = family.Id
	userResponse.IdBalancePoint = balancePoint.Id
	return userResponse
}
