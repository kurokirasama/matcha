package tui

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/floatpane/matcha/config"
	"github.com/floatpane/matcha/fetcher"
)

func TestInbox_Compose(t *testing.T) {
	accounts := []config.Account{
		{ID: "account-1", Email: "test@example.com"},
	}

	emails := []fetcher.Email{}
	inbox := NewInbox(emails, accounts)

	// Simulate pressing 'c' to compose
	_, cmd := inbox.Update(tea.KeyPressMsg{Code: 'c', Text: "c"})
	
	if cmd == nil {
		t.Fatal("Expected a command for composing, but got nil")
	}

	msg := cmd()
	// Should send GoToSendMsg
	_, ok := msg.(GoToSendMsg)
	if !ok {
		t.Fatalf("Expected GoToSendMsg, got %T", msg)
	}
}
