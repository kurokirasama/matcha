package tui

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/floatpane/matcha/config"
	"github.com/floatpane/matcha/fetcher"
)

func TestFolderInboxHorizontalLayout(t *testing.T) {
	accounts := []config.Account{
		{ID: "account-1", Email: "host.example.com"},
	}
	fi := NewFolderInbox([]string{"INBOX"}, accounts)
	fi.SetLayout(config.LayoutHorizontal)
	
	width, height := 100, 40
	fi.Update(tea.WindowSizeMsg{Width: width, Height: height})

	// Open a preview to trigger split calculations
	email := &fetcher.Email{UID: 1, AccountID: "account-1"}
	fi.OpenSplitPreview(email.UID, email.AccountID, email)

	inboxHeight := fi.calculateInboxHeight()
	previewHeight := fi.calculatePreviewHeight()

	// In horizontal mode, they should split the height
	if inboxHeight == height {
		t.Errorf("expected inbox height to be less than total height in horizontal mode, got %d", inboxHeight)
	}
	if previewHeight == height {
		t.Errorf("expected preview height to be less than total height in horizontal mode, got %d", previewHeight)
	}
}
