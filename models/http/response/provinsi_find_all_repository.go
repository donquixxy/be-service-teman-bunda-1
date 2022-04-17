package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type FindProvinsiAllResponse struct {
	IdProp   int    `json:"idprop"`
	KdProp   string `json:"kdprop"`
	NamaProp string `json:"nama_prop"`
	KodeArea string `json:"kode_area"`
}

func ToProvinsiFindAllResponse(provinsis []entity.Provinsi) (provinsiResponses []FindProvinsiAllResponse) {
	for _, provinsi := range provinsis {
		var provinsiResponse FindProvinsiAllResponse
		provinsiResponse.IdProp = provinsi.IdProp
		provinsiResponse.KdProp = provinsi.KdProp
		provinsiResponse.NamaProp = provinsi.NamaProp
		provinsiResponse.KodeArea = provinsi.KodeArea
		provinsiResponses = append(provinsiResponses, provinsiResponse)
	}
	return provinsiResponses
}
