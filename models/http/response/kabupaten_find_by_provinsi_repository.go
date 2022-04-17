package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type FindKabupatenByProvinsiResponse struct {
	IdKabu   int    `json:"idkabu"`
	IdProp   int    `json:"idprop"`
	KdKabu   string `json:"kdkabu"`
	NamaKabu string `json:"nama_kabu"`
}

func ToFindKabupatenByProvinsiResponse(kabupatens []entity.Kabupaten) (kabupatenResponses []FindKabupatenByProvinsiResponse) {
	for _, kabupaten := range kabupatens {
		var kabupatenResponse FindKabupatenByProvinsiResponse
		kabupatenResponse.IdKabu = kabupaten.IdKabu
		kabupatenResponse.IdProp = kabupaten.IdProp
		kabupatenResponse.KdKabu = kabupaten.KdKabu
		kabupatenResponse.NamaKabu = kabupaten.NamaKabu
		kabupatenResponses = append(kabupatenResponses, kabupatenResponse)
	}
	return kabupatenResponses
}
