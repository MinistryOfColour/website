package main

import (
	"crypto/tls"
	"log"
	"time"

	gomail "gopkg.in/gomail.v2"
)

var MailChannel = make(chan *gomail.Message, 64)

func NewMail() *gomail.Message {

	m := gomail.NewMessage()
	m.SetHeader("From", "umpire@wargaming.io")
	return m
}

func InitMailer(c ConfigType) chan *gomail.Message {
	go mailerDaemon(c)
	return MailChannel
}

func mailerDaemon(c ConfigType) {

	log.Println("starting mailer with", c.MailServer, c.MailPort, c.MailUser, c.MailPasswd)
	d := gomail.NewPlainDialer(
		c.MailServer,
		c.MailPort,
		c.MailUser,
		c.MailPasswd)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	var s gomail.SendCloser
	var err error
	open := false
	for {
		select {
		case m, ok := <-MailChannel:
			if !ok {
				return
			}
			if !open {
				if s, err = d.Dial(); err != nil {
					println("Mail Dialer Error: ", err.Error())
				} else {
					open = true
				}
			}
			if open {
				if err := gomail.Send(s, m); err != nil {
					log.Print(err)
				}
			}
		// Close the connection to the SMTP server if no email was sent in
		// the last couple of minutes.
		case <-time.After(120 * time.Second):
			if open {
				if err := s.Close(); err != nil {
					print("Mail Close Error: ", err.Error())
				}
				open = false
			}
		}
	}
}
