package main

import (
	"bytes"
	"context"
	"strings"
	"text/template"

	"github.com/momotaro98/mixlunch-email-notification/emailsender"
)

const (
	sender  = "noreply@mixlunch.com"
	subject = "[Mixlunch] Today’s Lunch Group"
)

type Recipient struct {
	UserId   string
	Email    string
	UserName string
	Party    Party
}

type Party struct {
	Date            string
	TimeFrom        string
	TimeTo          string
	ChatRoomId      string
	MemberUserNames []string
}

var bodyTpl *template.Template

func InitBodyTemplate() {
	bodyTpl = template.Must(template.ParseFiles("./templates/body.html.tpl"))
}

func (r *Recipient) AssembleEmailBody() (string, error) {
	// Modify Times from UTC from JST for Japanese users
	jstTimeFrom, err := ShiftTimeStringUtfToJst(r.Party.TimeFrom)
	if err != nil {
	}
	r.Party.TimeFrom = jstTimeFrom
	jstTimeTo, err := ShiftTimeStringUtfToJst(r.Party.TimeTo)
	if err != nil {
	}
	r.Party.TimeTo = jstTimeTo

	// Modify chatRoomId according to frontend app side specification
	squashedDateString, err := SquashDateString(r.Party.Date)
	if err != nil {
		return "", err
	}
	r.Party.ChatRoomId = squashedDateString

	// Modify Date string
	r.Party.Date = strings.Replace(r.Party.Date, "-", "/", 2)

	// Populate to body template
	var tpl bytes.Buffer
	if err := bodyTpl.Execute(&tpl, *r); err != nil {
		return "", err
	}
	body := tpl.String()
	return body, nil
}

func (r *Recipient) Send(ctx context.Context, emailSender emailsender.EmailSender, body string) error {
	sender := sender
	subject := subject
	recipient := r.Email

	if err := emailSender.Send(ctx, sender, subject, body, recipient); err != nil {
		return err
	}

	return nil
}
