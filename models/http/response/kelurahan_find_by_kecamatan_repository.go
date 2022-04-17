package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type FindKelurahanByKecamatanResponse struct {
	IdKelu   int    `json:"idkelu"`
	IdKeca   int    `json:"kdkeca"`
	IdKabu   int    `json:"idkabu"`
	IdProp   int    `json:"idprop"`
	KdKelu   string `json:"kdkelu"`
	NamaKelu string `json:"nama_kelu"`
}

func ToFindKelurahanByKecamatanResponse(kelurahans []entity.Kelurahan) (kelurahanResponses []FindKelurahanByKecamatanResponse) {
	for _, kelurahan := range kelurahans {
		var kelurahanResponse FindKelurahanByKecamatanResponse
		kelurahanResponse.IdKelu = kelurahan.IdKelu
		kelurahanResponse.IdKeca = kelurahan.IdKeca
		kelurahanResponse.IdKabu = kelurahan.IdKabu
		kelurahanResponse.IdProp = kelurahan.IdProp
		kelurahanResponse.KdKelu = kelurahan.KdKelu
		kelurahanResponse.NamaKelu = kelurahan.NamaKelu
		kelurahanResponses = append(kelurahanResponses, kelurahanResponse)
	}
	return kelurahanResponses
}
