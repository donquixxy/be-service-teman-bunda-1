package service

type PushNotificationRequestBody struct {
	ToDeviceToken []string         `json:"registration_ids"`
	Notification  NotificationData `json:"notification"`
}

type NotificationData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
