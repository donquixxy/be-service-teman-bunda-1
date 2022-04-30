package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type FindAllBannerResponse struct {
	Id          string `json:"id"`
	BannerTitle string `json:"banner_title"`
	BannerUrl   string `json:"banner_url"`
}

func ToFindAllBannerResponse(banners []entity.Banner) (bannerResponses []FindAllBannerResponse) {
	for _, banner := range banners {
		var bannerResponse FindAllBannerResponse
		bannerResponse.Id = banner.Id
		bannerResponse.BannerTitle = banner.BannerTitle
		bannerResponse.BannerUrl = banner.BannerUrl
		bannerResponses = append(bannerResponses, bannerResponse)
	}
	return bannerResponses
}
