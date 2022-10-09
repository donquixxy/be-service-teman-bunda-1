package response

import "github.com/tensuqiuwulu/be-service-teman-bunda/models/entity"

type FindSettingVerApp struct {
	SettingName string `json:"os_type"`
	Value       string `json:"ver_app"`
}

func ToFindSettingVerApp(verApp entity.Settings) (verAppResponse FindSettingVerApp) {
	verAppResponse.SettingName = verApp.SettingsName
	verAppResponse.Value = verApp.SettingsTitle
	return verAppResponse
}
