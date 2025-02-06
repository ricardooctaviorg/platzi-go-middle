package main

import "fmt"

//Solo SMS & Email

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type InotificationFactroy interface {
	SendNotification()
	GetSender() ISender
}

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS")
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification via Email")
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "AWS"
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

func getNotificationFactory(notificationType string) (InotificationFactroy, error) {
	if notificationType == "sms" {
		return &EmailNotification{}, nil
	} else if notificationType == "email" {
		return &EmailNotification{}, nil
	} else {
		return nil, fmt.Errorf("No existe tipo")
	}
}

func sendNotification(factory InotificationFactroy) {
	factory.SendNotification()
}

func getMethod(factory InotificationFactroy) {
	fmt.Println(factory.GetSender().GetSenderMethod())
}

func GetChannel(factory InotificationFactroy) {
	fmt.Println(factory.GetSender().GetSenderChannel())
}

func main() {
	fmt.Println("Hello World")

	smsNotificator, error := getNotificationFactory("hola")

	if error == nil {
		sendNotification(smsNotificator)
		getMethod(smsNotificator)
		GetChannel(smsNotificator)
	} else {
		fmt.Println(error)
	}
}
