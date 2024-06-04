package repository

import (
	"NotificationOfBirthdays"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SubscriptionPostgres struct {
	db *sqlx.DB
}

func NewSubscriptionPostgres(db *sqlx.DB) *SubscriptionPostgres {
	return &SubscriptionPostgres{db: db}
}

func (r *SubscriptionPostgres) CreateSubscription(userId int, subscription NotificationOfBirthdays.Subscription) (int, error) {
	var subscriptionId int
	queryGetSubscription := fmt.Sprintf("SELECT id FROM %s "+
		"WHERE subscriber_user_id = $1 AND birthday_user_id = $2", subscriptionsTable)
	err := r.db.Get(&subscriptionId, queryGetSubscription, userId, subscription.BirthdayUserId)
	if err == nil {
		return 0, errors.New("You have already subscribed to this user")
	}

	queryCreateSubscription := fmt.Sprintf("INSERT INTO %s (subscriber_user_id, birthday_user_id, birthday_date, notification_date)"+
		" values ($1, $2, $3, $4) RETURNING id", subscriptionsTable)
	row := r.db.QueryRow(queryCreateSubscription, userId, subscription.BirthdayUserId, subscription.BirthdayDate, subscription.NotificationDate)
	if err := row.Scan(&subscriptionId); err != nil {
		return 0, err
	}

	return subscriptionId, nil
}
