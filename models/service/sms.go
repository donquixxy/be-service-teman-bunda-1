package service

type SmsBody struct {
	UserKey string `json:"userkey"`
	PassKey string `json:"passkey"`
	To      string `json:"to"`
	Message string `json:"message"`
}
