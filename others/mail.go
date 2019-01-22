package main

import (
	"fmt"
	"strings"
	"net/smtp"
)

func SendMail(user, password, host, to, subject, body, nickname, mailtype string) error {
	ipPort := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, ipPort[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/html; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + nickname + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";") //多个邮箱
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func main() {
	user := ""
	password := ""
	host := "smtp.exmail.qq.com:25"
	to := "962691478@qq.com;571684122@qq.com"
	subject := "这是通过Golang自动发送的邮件"
	body := `
		<html>
		<body>
		<h3>
		"This email is from 皮皮虾"
		</h3>
		</body>
		</html>
		`
	nickname := "皮皮虾"
	err := SendMail(user, password, host, to, subject, body, nickname, "html")
	if err != nil {
		fmt.Println("Send mail error,err = ", err)
		return
	}
	fmt.Println("Send mail success!")
}
