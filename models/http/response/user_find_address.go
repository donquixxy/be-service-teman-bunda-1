package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type FindUserShippingAddress struct {
	Id        string  `json:"id"`
	Status    int     `json:"status"`
	Address   string  `json:"address"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    float64 `json:"radius"`
	Note      string  `json:"note"`
}

func ToFindUserShippingAddressResponse(userShippingAddresss []entity.UserShippingAddress) (userShippingAddressResponses []FindUserShippingAddress) {
	for _, userShippingAddress := range userShippingAddresss {
		var userShippingAddressResponse FindUserShippingAddress
		userShippingAddressResponse.Id = userShippingAddress.Id
		userShippingAddressResponse.Status = userShippingAddress.Status
		userShippingAddressResponse.Address = userShippingAddress.Address
		userShippingAddressResponse.Latitude = userShippingAddress.Latitude
		userShippingAddressResponse.Longitude = userShippingAddress.Longitude
		userShippingAddressResponse.Radius = userShippingAddress.Radius
		userShippingAddressResponse.Note = userShippingAddress.Note
		userShippingAddressResponses = append(userShippingAddressResponses, userShippingAddressResponse)
	}
	return userShippingAddressResponses
}
