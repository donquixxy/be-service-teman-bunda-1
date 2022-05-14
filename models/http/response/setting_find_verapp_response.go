package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type FindSettingVerApp struct {
	Value string `json:"ver_app"`
}

func ToFindSettingVerApp(verApp entity.Settings) (verAppResponse FindSettingVerApp) {
	verAppResponse.Value = verApp.SettingsTitle
	return verAppResponse
}
