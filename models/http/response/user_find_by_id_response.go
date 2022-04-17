package response

import (
	"github.com/shopspring/decimal"
	"github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"
)

type FindUserByIdResponse struct {
	Id            string          `json:"id"`
	FullName      string          `json:"full_name"`
	Username      string          `json:"username"`
	Email         string          `json:"email"`
	Address       string          `json:"address"`
	Phone         string          `json:"phone"`
	ReferalCode   string          `json:"referal_code"`
	BalancePoints decimal.Decimal `json:"balance_points"`
	IdProvinsi    int             `json:"idprop"`
	IdKabupaten   int             `json:"idkabu"`
	IdKecamatan   int             `json:"idkeca"`
	IdKelurahan   int             `json:"idkelu"`
}

func ToUserFindByIdResponse(user entity.User) (userResponse FindUserByIdResponse) {
	userResponse.Id = user.Id
	userResponse.Username = user.Username
	userResponse.FullName = user.FamilyMembers.FullName
	userResponse.Email = user.FamilyMembers.Email
	userResponse.Address = user.FamilyMembers.Address
	userResponse.Phone = user.FamilyMembers.Phone
	userResponse.ReferalCode = user.ReferalCode
	userResponse.BalancePoints = user.BalancePoint.BalancePoints
	userResponse.IdProvinsi = user.FamilyMembers.Family.IdProvinsi
	userResponse.IdKabupaten = user.FamilyMembers.Family.IdKabupaten
	userResponse.IdKecamatan = user.FamilyMembers.Family.IdKecamatan
	userResponse.IdKelurahan = user.FamilyMembers.Family.IdKelurahan
	return userResponse
}
