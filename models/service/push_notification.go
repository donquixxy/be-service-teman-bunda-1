package service

type PushNotificationRequestBody struct {
	ToDeviceToken string           `json:"to"`
	Priority      string           `json:"priority"`
	SoundName     string           `json:"soundname"`
	Notification  NotificationData `json:"notification"`
}

type NotificationData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
