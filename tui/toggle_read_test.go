package tui

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/floatpane/matcha/config"
	"github.com/floatpane/matcha/fetcher"
)

func TestInbox_ToggleRead(t *testing.T) {
	accounts := []config.Account{
		{ID: "account-1", Email: "test@example.com"},
	}

	emails := []fetcher.Email{
		{UID: 1, From: "a@example.com", Subject: "Email 1", AccountID: "account-1", IsRead: false},
	}

	inbox := NewInbox(emails, accounts)

	// Simulate pressing 'u' to toggle read status
	_, cmd := inbox.Update(tea.KeyPressMsg{Code: 'u', Text: "u"})

	if cmd == nil {
		t.Fatal("Expected a command for toggling read status, but got nil")
	}

	msg := cmd()
	markReadMsg, ok := msg.(MarkEmailAsReadMsg)
	if !ok {
		t.Fatalf("Expected MarkEmailAsReadMsg, got %T", msg)
	}

	if markReadMsg.UID != 1 {
		t.Errorf("Expected UID 1, got %d", markReadMsg.UID)
	}

	// Now simulate it being read
	emails[0].IsRead = true
	inbox.SetEmails(emails, accounts)

	_, cmd = inbox.Update(tea.KeyPressMsg{Code: 'u', Text: "u"})
	if cmd == nil {
		t.Fatal("Expected a command for toggling unread status, but got nil")
	}

	msg = cmd()
	markUnreadMsg, ok := msg.(MarkEmailAsUnreadMsg)
	if !ok {
		t.Fatalf("Expected MarkEmailAsUnreadMsg, got %T", msg)
	}

	if markUnreadMsg.UID != 1 {
		t.Errorf("Expected UID 1, got %d", markUnreadMsg.UID)
	}
}

func TestEmailView_ToggleRead(t *testing.T) {
	email := fetcher.Email{
		UID:       42,
		From:      "sender@example.com",
		Subject:   "Toggle Test",
		AccountID: "account-1",
		IsRead:    false,
	}

	view := NewEmailView(email, 0, 80, 24, MailboxInbox, "INBOX", false)

	// Simulate pressing 'u' to toggle read status
	_, cmd := view.Update(tea.KeyPressMsg{Code: 'u', Text: "u"})

	if cmd == nil {
		t.Fatal("Expected a command for toggling read status in EmailView, but got nil")
	}

	msg := cmd()
	batch, ok := msg.(tea.BatchMsg)
	if !ok {
		t.Fatalf("Expected tea.BatchMsg, got %T", msg)
	}

	foundMarkRead := false
	foundBackToMailbox := false

	for _, m := range batch {
		if subMsg := m(); subMsg != nil {
			switch res := subMsg.(type) {
			case MarkEmailAsReadMsg:
				foundMarkRead = true
				if res.UID != 42 {
					t.Errorf("Expected UID 42, got %d", res.UID)
				}
			case BackToMailboxMsg:
				foundBackToMailbox = true
			}
		}
	}

	if !foundMarkRead {
		t.Error("Did not find MarkEmailAsReadMsg in batch")
	}
	if !foundBackToMailbox {
		t.Error("Did not find BackToMailboxMsg in batch")
	}
}
