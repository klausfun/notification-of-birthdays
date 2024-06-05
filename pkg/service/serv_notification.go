package service

import (
	"NotificationOfBirthdays"
	"NotificationOfBirthdays/pkg/repository"
	"bytes"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	"html/template"
	"os"
	"time"
)

type NotificationService struct {
	repo repository.Subscription
}

func NewNotificationService(repo repository.Subscription) *NotificationService {
	return &NotificationService{repo: repo}
}

func (s *NotificationService) CheckAndSendNotifications() error {
	usersAndSubscriptions, err := s.repo.GetAllSubscriptions()
	if err != nil {
		return err
	}

	for _, usAndSub := range usersAndSubscriptions {
		for _, sub := range usAndSub.Subscriptions {
			if sub.NotificationDate.Format("02-01-2006 15:04") == time.Now().Format("02-01-2006 15:04") {
				err = sendGomail(viper.GetString("notification.pathToTheMessage"),
					usAndSub.User.Name, sub.Name, usAndSub.User.Email, sub.BirthdayDate)
			}
		}
	}

	return err
}

func sendGomail(templatePath, recipientsName, birthdayName, recipientsEmail string,
	birthdayDate NotificationOfBirthdays.CustomTime) error {

	var body bytes.Buffer
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		logrus.Fatalf("error send gmail: %s", err.Error())
	}
	err = t.Execute(&body, struct {
		BirthdayName, RecipientsName, RecipientsEmail string
		BirthdayDate                                  NotificationOfBirthdays.CustomTime
	}{
		BirthdayName:    birthdayName,
		RecipientsName:  recipientsName,
		BirthdayDate:    birthdayDate,
		RecipientsEmail: recipientsEmail,
	})
	if err != nil {
		logrus.Fatalf("error executing template: %s", err.Error())
	}

	m := gomail.NewMessage()
	m.SetHeader("From", viper.GetString("notification.email"))
	m.SetHeader("To", recipientsEmail)
	m.SetHeader("Subject", "Birthday reminder!")
	m.SetBody("text/html", body.String())

	d := gomail.NewDialer("smtp.gmail.com", 587,
		viper.GetString("notification.email"), os.Getenv("GMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	return err
}
