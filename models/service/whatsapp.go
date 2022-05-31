package service

type SendWhatsappRequest struct {
	ToNumber        string             `json:"to_number"`
	ToName          string             `json:"to_name"`
	MssgTemplateId  string             `json:"message_template_id"`
	ChIntegrationId string             `json:"channel_integration_id"`
	Language        WhatsappLanguage   `json:"language"`
	Parameters      WhatsappParameters `json:"parameters"`
}

type WhatsappLanguage struct {
	Code string `json:"code"`
}

type WhatsappParameters struct {
	Bodys []WhatsappBody `json:"body"`
}

type WhatsappBody struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	ValueText string `json:"value_text"`
}
