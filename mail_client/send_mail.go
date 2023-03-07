package mail_client

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

func (mail *MailClient) SendEmail(toMail []string, subject string, data map[string]interface{}, templateEmail string) error {

	gmailAuth := smtp.PlainAuth("", mail.config.Username, mail.config.Password, mail.config.Host)
	t, err := template.ParseFiles(templateEmail)
	if err != nil {
		return err
	}

	var body bytes.Buffer
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"utf-8\""
	body.Write([]byte(fmt.Sprintf("Subject: %s\n%s\n\n", subject, headers)))

	t.Execute(&body, data)
	return smtp.SendMail(fmt.Sprintf("%s:%d", mail.config.Host, mail.config.Port), gmailAuth, mail.config.Address, toMail, body.Bytes())
}
