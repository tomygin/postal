package postal

import (
	"net/smtp"
	"strings"
)

type QQMail struct {
	SendAddr    string
	ReceiveAddr []string
	AuthCode    string

	auth smtp.Auth
	//waitTime time.Duration
}

func (q *QQMail) Init() bool {
	q.auth = smtp.PlainAuth("", q.SendAddr, q.AuthCode, "smtp.qq.com")

	return q.auth != nil
}

func (q *QQMail) Send(title, msg string) bool {
	nickName := "Notify"
	contentType := "Content-Type:text/plain;charset=UTF-8\r\n"
	content := []byte("To: " + strings.Join(q.ReceiveAddr, ",") + "\r\nFrom: " + nickName +
		"<" + q.SendAddr + ">\r\nSubject: " + title + "\r\n" + contentType + "\r\n\r\n" + msg)

	err := smtp.SendMail("smtp.qq.com:25", q.auth, q.SendAddr, q.ReceiveAddr, content)
	return err == nil
}

var _ Msger = (*QQMail)(nil)
