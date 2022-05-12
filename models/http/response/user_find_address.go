package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type FindUserShippingAddress struct {
	IdUser    string  `json:"id_user"`
	Status    int     `json:"status"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    float64 `json:"radius"`
}

func ToFindUserShippingAddressResponse(userShippingAddresss []entity.UserShippingAddress) (userShippingAddressResponses []FindUserShippingAddress) {
	for _, userShippingAddress := range userShippingAddresss {
		var userShippingAddressResponse FindUserShippingAddress
		userShippingAddressResponse.IdUser = userShippingAddress.Id
		userShippingAddressResponse.Status = userShippingAddress.Status
		userShippingAddressResponse.Address = userShippingAddress.Address
		userShippingAddressResponse.Latitude = userShippingAddress.Latitude
		userShippingAddressResponse.Longitude = userShippingAddress.Longitude
		userShippingAddressResponse.Radius = userShippingAddress.Radius
		userShippingAddressResponses = append(userShippingAddressResponses, userShippingAddressResponse)
	}
	return userShippingAddressResponses
}
