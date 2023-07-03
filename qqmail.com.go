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
}

func (q *QQMail) Init() bool {
	q.auth = smtp.PlainAuth("", q.SendAddr, q.AuthCode, "smtp.qq.com")

	return q.auth != nil
}

func (q *QQMail) Msg(title, msg string) (c chan struct{}) {
	nickName := "Notify"
	contentType := "Content-Type:text/plain;charset=UTF-8\r\n"
	content := []byte("To: " + strings.Join(q.ReceiveAddr, ",") + "\r\nFrom: " + nickName +
		"<" + q.SendAddr + ">\r\nSubject: " + title + "\r\n" + contentType + "\r\n\r\n" + msg)

	smtp.SendMail("smtp.qq.com:25", q.auth, q.SendAddr, q.ReceiveAddr, content)
	c = make(chan struct{}, 1)
	c <- struct{}{}
	return

}

var _ Msger = (*QQMail)(nil)
