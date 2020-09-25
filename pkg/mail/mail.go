package mail

import (
	"errors"
	"fmt"
	"net/smtp"
	"strconv"
	"strings"
)

// Mail struct
type Mail struct {
	From    string
	To      []string
	Subject string
	Body    string
}

// Send sends mail.
func Send(from string, to []string, subject, body, smptUser, smtpPassword, smtpHost string, smtpPort int) error {
	mail := Mail{}
	mail.From = from
	mail.To = to
	mail.Subject = subject
	mail.Body = body

	messageBody, err := mail.buildMessage()
	if err != nil {
		return err
	}

	// Envoi avec authentification
	// ---------------------------
	if smptUser != "" && smtpPassword != "" {
		auth := smtp.PlainAuth("", smptUser, smtpPassword, smtpHost)

		err = smtp.SendMail(serverName(smtpHost, smtpPort),
			auth,
			mail.From,
			mail.To,
			[]byte(messageBody))
		if err != nil {
			return err
		}

		return nil
	}

	// Envoi sans authentification
	// ---------------------------
	return sendMail(mail, messageBody, smtpHost, smtpPort)
}

// buildMessage constructs mail.
func (mail *Mail) buildMessage() (string, error) {
	if len(mail.From) == 0 {
		return "", errors.New("From fields empty")
	}

	if len(mail.To) == 0 {
		return "", errors.New("To fields empty")
	}

	header := ""
	header += fmt.Sprintf("From: %s\r\n", mail.From)
	header += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ","))
	header += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	header += fmt.Sprintf("MIME-version: %s\r\n", "1.0")
	header += fmt.Sprintf("Content-Type: %s; charset: %s\r\n", "text/html", "UTF-8")
	header += "\r\n" + mail.Body

	return header, nil
}

// serverName returns SMTP server name from host and port.
func serverName(host string, port int) (s string) {
	s = host

	if port > 0 {
		s += ":" + strconv.Itoa(port)
	}

	return s
}

// sendMail constructs and sends mail.
func sendMail(mail Mail, msg, smtpHost string, smtpPort int) error {
	c, err := smtp.Dial(serverName(smtpHost, smtpPort))
	if err != nil {
		return err
	}
	defer c.Close()

	if err = c.Mail(mail.From); err != nil {
		return err
	}

	for i := range mail.To {
		if err = c.Rcpt(mail.To[i]); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(msg))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}
