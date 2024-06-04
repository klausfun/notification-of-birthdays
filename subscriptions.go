package NotificationOfBirthdays

type Subscription struct {
	Id               int    `json:"-" db:"id"`
	SubscriberUserId int    `json:"subscriber_user_id"`
	BirthdayUserId   int    `json:"birthday_user_id" binding:"required"`
	BirthdayDate     string `json:"birthday_date" binding:"required"`
	NotificationDate string `json:"notification_date" binding:"required"`
}
