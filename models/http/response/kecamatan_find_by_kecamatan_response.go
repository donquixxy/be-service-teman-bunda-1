package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type FindKecamatanByKabupatenResponse struct {
	IdKeca   int    `json:"idkeca"`
	IdKabu   int    `json:"kdkabu"`
	IdProp   int    `json:"idprop"`
	KdKeca   string `json:"kdkeca"`
	NamaKeca string `json:"nama_keca"`
}

func ToFindKecamatanByKabuaptenResponse(kecamatans []entity.Kecamatan) (kecamatanResponses []FindKecamatanByKabupatenResponse) {
	for _, kecamatan := range kecamatans {
		var kecamatanResponse FindKecamatanByKabupatenResponse
		kecamatanResponse.IdKeca = kecamatan.IdKeca
		kecamatanResponse.IdKabu = kecamatan.IdKabu
		kecamatanResponse.IdProp = kecamatan.IdProp
		kecamatanResponse.KdKeca = kecamatan.KdKeca
		kecamatanResponse.NamaKeca = kecamatan.NamaKeca
		kecamatanResponses = append(kecamatanResponses, kecamatanResponse)
	}
	return kecamatanResponses
}
