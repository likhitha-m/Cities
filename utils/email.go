package utils

import (
	"crypto/tls"
	// "log"
	// "net/smtp"
	"strings"

	// "errors"
	"fmt"
	"os"

	gomail "gopkg.in/mail.v2"

	logger "github.com/sirupsen/logrus"
)

//Request struct
type SendEmailRequest struct {
	from    string
	to      string
	subject string
	body    string
}

func NewSendEmailRequest(from string, to string, subject, body string) *SendEmailRequest {
	return &SendEmailRequest{
		from:    from,
		to:      to,
		subject: subject,
		body:    body,
	}
}


type SendTemplateEmailRequest struct {
	from         string
	to           string
	temapletName string
	temapletData string
}
type Mail struct {
    Sender  string
    To      []string
    Subject string
    Body    string
}


func NewSendTemplateEmailReciever(from string, to string, temaplet_name, temaplet_data string) *SendTemplateEmailRequest {
	return &SendTemplateEmailRequest{
		from:         from,
		to:           to,
		temapletName: temaplet_name,
		temapletData: temaplet_data,
	}
}

func (r *SendTemplateEmailRequest) SendTemplateEmail() error {

	m := gomail.NewMessage()

  // Set E-Mail sender
  m.SetHeader("From", os.Getenv("EMAIL"))

  // Set E-Mail receivers
  m.SetHeader("To", r.to)

  // Set E-Mail subject
  m.SetHeader("Subject", r.temapletName)

  // Set E-Mail body. You can set plain text or html with text/html
  m.SetBody("text/html", r.temapletData)
fmt.Println("usrname--->>",os.Getenv("EMAIL"),os.Getenv("PASSWORD") )
  // Settings for SMTP server
  d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"),os.Getenv("PASSWORD"))

  // This is only needed when SSL/TLS certificate is not valid on server.
  // In production this should be set to false.
  d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

  // Now send E-Mail
  if err := d.DialAndSend(m); err != nil {
    fmt.Println(err)
    panic(err)
  }
	
// sender := "progolang1@gmail.com"

// to := []string{
// 	r.to,
// }

// user := "progolang1@gmail.com"
// password := "Admin@123"

// subject := "Simple HTML mail"
// body := `<p>An old <b>falcon</b> in the sky.</p>`

// request := Mail{
// 	Sender:  sender,
// 	To:      to,
// 	Subject: subject,
// 	Body:    body,
// }

// addr := "smtp.mailtrap.io:2525"
// host := "smtp.mailtrap.io"

// msg := BuildMessage(request)
// auth := smtp.PlainAuth("", user, password, host)
// err := smtp.SendMail(addr, auth, sender, to, []byte(msg))

// if err != nil {
// 	log.Fatal(err)
// }

fmt.Println("Email sent successfully")

	

	logger.Info("Template Email Sent for template : "+r.temapletName+"!!! and result is: ", )
	return nil
}
func BuildMessage(mail Mail) string {
    msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
    msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
    msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
    msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
    msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

    return msg
}

type SendRawEmailRequest struct {
	From string   `json:"from" validate:"required"`
	To   []string `json:"to" validate:"required"`
	Body string   `json:"body" validate:"required"`
}



func NewSendRawEmailRequest(from string, to []string, body string) *SendRawEmailRequest {
	return &SendRawEmailRequest{
		From: from,
		To:   to,
		Body: body,
	}
}
