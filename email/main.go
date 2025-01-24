package main

import (
	"fmt"
	"net/smtp"
)

func main() {

	// Информация об отправителе
	//from := "CIT\\Mishchenko.R"
	from := "Mishchenko.R@mintrud.by"
	//password := "пароль"

	// Информация о получателе
	//to := []string{
	//	"romah29ius@gmail.com",
	//}

	// smtp сервер конфигурация
	//smtpHost := "172.16.68.141"
	//smtpPort := "587"

	// Сообщение.
	message := []byte("To: romah28ius@gmail.com\r\n" +
		"Subject: Hi there!\r\n" +
		"Content-Type: text/plain; charset=UTF-8\r\n" +
		"\r\n" +
		"Hi!\r\n")

	to := make([]string, 1)
	// Авторизация.
	//auth := smtp.PlainAuth("", from, "Mrv.789", smtpHost)

	// Отправка почты.
	err := smtp.SendMail("172.16.68.141:587",
		nil, /* this is the optional Auth */
		from, to, message)
	fmt.Println(err)
}
