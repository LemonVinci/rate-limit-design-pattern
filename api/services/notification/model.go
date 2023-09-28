package notification

type Notification struct {
	ID      int16  `json:id`
	Type    string `json:type`
	UserId  string `json:userId`
	Message string `json:message`
}
