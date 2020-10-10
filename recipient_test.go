package main

import (
	"context"
	"os"
	"testing"
)

var fullRecipient Recipient

func TestAssembleEmailBody_FullRecipient_NoIssue(t *testing.T) {
	// Arrange
	copiedFullRecipient := fullRecipient
	r := &copiedFullRecipient
	// Act
	_, err := r.AssembleEmailBody()
	// Assert
	if err != nil {
		t.Errorf("Test failed. Expected: nil, Actual: %v", err)
	}
}

type stubCorrectEmailSender struct{}

func (s *stubCorrectEmailSender) Send(ctx context.Context, sender string, subject string, body string, recipient string) error {
	return nil
}

func TestSend_CorrectEmailSender_NoIssue(t *testing.T) {
	// Arrange
	copiedFullRecipient := fullRecipient
	r := &copiedFullRecipient
	stubEmailSender := &stubCorrectEmailSender{}
	ctx := context.Background()
	body := "Test Body"
	// Act
	err := r.Send(ctx, stubEmailSender, body)
	// Assert
	if err != nil {
		t.Errorf("Test failed. Expected: nil, Actual: %v", err)
	}
}

func setup() {
	InitBodyTemplate()

	party := Party{
		Date:            "2019-05-01",
		TimeFrom:        "3:00", // UTC timezone
		TimeTo:          "4:00",
		ChatRoomId:      "chat-room-id",
		MemberUserNames: []string{"Michel", "John", "Ash"},
	}
	recipient := &Recipient{
		UserId:   "user-id",
		Email:    "test@example.com",
		UserName: "Michel",
		Party:    party,
	}
	fullRecipient = *recipient
}

func teardown() {}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}
