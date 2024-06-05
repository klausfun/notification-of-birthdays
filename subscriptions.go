package NotificationOfBirthdays

import (
	"encoding/json"
	"fmt"
	"time"
)

type CustomTime struct {
	time.Time
}

const ctLayout = "02-01-2006 15:04"

func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := string(b[1 : len(b)-1])
	parsedTime, err := time.Parse(ctLayout, s)
	if err != nil {
		return fmt.Errorf("error parsing time: %v", err)
	}
	ct.Time = parsedTime
	return nil
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Format(ctLayout))
}

func (ct CustomTime) TimeValue() time.Time {
	return ct.Time
}

func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		*ct = CustomTime{Time: time.Time{}}
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		*ct = CustomTime{Time: v}
		return nil
	case []byte:
		parsedTime, err := time.Parse(time.RFC3339, string(v))
		if err != nil {
			return err
		}
		*ct = CustomTime{Time: parsedTime}
		return nil
	case string:
		parsedTime, err := time.Parse(time.RFC3339, v)
		if err != nil {
			return err
		}
		*ct = CustomTime{Time: parsedTime}
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into CustomTime", value)
	}
}

type Subscription struct {
	Id               int        `json:"-" db:"id"`
	SubscriberUserId int        `json:"subscriber_user_id"`
	BirthdayUserId   int        `json:"birthday_user_id" binding:"required"`
	BirthdayDate     CustomTime `json:"birthday_date" binding:"required"`
	NotificationDate CustomTime `json:"notification_date" binding:"required"`
}

type UserSubscription struct {
	Id               int        `json:"id" db:"birthday_user_id"`
	Name             string     `json:"name" db:"name"`
	Email            string     `json:"email" db:"email"`
	BirthdayDate     CustomTime `json:"birthday_date" db:"birthday_date"`
	NotificationDate CustomTime `json:"notification_date" db:"notification_date"`
}

type UserAndHisSubscriptions struct {
	User          Author             `json:"user"`
	Subscriptions []UserSubscription `json:"subscriptions"`
}
