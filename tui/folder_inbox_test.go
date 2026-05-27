package tui

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/floatpane/matcha/config"
	"github.com/floatpane/matcha/fetcher"
)

func TestFolderInbox_InitialState(t *testing.T) {
	accounts := []config.Account{
		{ID: "account-1", Email: "host.example.com", FetchEmail: "first@example.com"},
	}
	folders := []string{"INBOX", "Sent", "Trash"}
	fi := NewFolderInbox(folders, accounts)

	if fi.currentFolder != "INBOX" {
		t.Errorf("expected current folder to be INBOX, got %s", fi.currentFolder)
	}

	if len(fi.folders) != 3 {
		t.Errorf("expected 3 folders, got %d", len(fi.folders))
	}
}

func TestFolderInbox_SwitchFolder(t *testing.T) {
	accounts := []config.Account{
		{ID: "account-1", Email: "host.example.com", FetchEmail: "first@example.com"},
	}
	folders := []string{"INBOX", "Sent"}
	fi := NewFolderInbox(folders, accounts)

	// Switch to Sent
	model, cmd := fi.Update(tea.KeyPressMsg{Code: tea.KeyTab, Text: "tab"})
	fi = model.(*FolderInbox)

	if fi.currentFolder != "Sent" {
		t.Errorf("expected current folder to be Sent, got %s", fi.currentFolder)
	}

	if cmd == nil {
		t.Fatal("expected a command after switching folder")
	}

	msg := cmd()
	switch m := msg.(type) {
	case SwitchFolderMsg:
		if m.FolderName != "Sent" {
			t.Errorf("expected SwitchFolderMsg with Sent, got %s", m.FolderName)
		}
	default:
		t.Errorf("expected SwitchFolderMsg, got %T", msg)
	}
}

func TestFolderInbox_MoveToFolder(t *testing.T) {
	accounts := []config.Account{
		{ID: "account-1", Email: "host.example.com", FetchEmail: "first@example.com"},
	}
	folders := []string{"INBOX", "Archive"}
	fi := NewFolderInbox(folders, accounts)
	fi.SetEmails([]fetcher.Email{{UID: 1, AccountID: "account-1", Subject: "Test"}}, accounts)

	// Start move
	model, _ := fi.Update(tea.KeyPressMsg{Code: 'm', Text: "m"})
	fi = model.(*FolderInbox)

	if !fi.movingEmail {
		t.Error("expected movingEmail to be true after pressing 'm'")
	}

	// Cancel move
	model, _ = fi.Update(tea.KeyPressMsg{Code: tea.KeyEsc, Text: "esc"})
	fi = model.(*FolderInbox)

	if fi.movingEmail {
		t.Error("expected movingEmail to be false after pressing 'esc'")
	}
}

func TestFolderInbox_KeybindingsDuringFilter(t *testing.T) {
	accounts := []config.Account{
		{ID: "account-1", Email: "host.example.com", FetchEmail: "first@example.com"},
	}
	fi := NewFolderInbox([]string{"INBOX"}, accounts)

	// Mock filtering state
	fi.inbox.list.SetFilteringEnabled(true)
	fi.inbox.list.Update(tea.KeyPressMsg{Code: '/', Text: "/"})

	// Press "m" (matches move-to-folder keybinding in FolderInbox)
	model, _ := fi.Update(tea.KeyPressMsg{Code: 'm', Text: "m"})
	fi = model.(*FolderInbox)

	if fi.movingEmail {
		t.Error("FolderInbox keybinding should not trigger while list is filtering")
	}
}

func TestFolderInbox_KeybindingsDuringSearch(t *testing.T) {
	accounts := []config.Account{
		{ID: "account-1", Email: "host.example.com", FetchEmail: "first@example.com"},
	}
	fi := NewFolderInbox([]string{"INBOX"}, accounts)

	// Mock search overlay active
	fi.inbox.searchOverlay = &SearchOverlay{}

	// Press "m" (matches move-to-folder keybinding in FolderInbox)
	model, _ := fi.Update(tea.KeyPressMsg{Code: 'm', Text: "m"})
	fi = model.(*FolderInbox)

	if fi.movingEmail {
		t.Error("FolderInbox keybinding should not trigger while search overlay is active")
	}
}

func TestFolderInbox_ToggleSidebar(t *testing.T) {
	accounts := []config.Account{
		{ID: "account-1", Email: "host.example.com", FetchEmail: "first@example.com"},
	}
	fi := NewFolderInbox([]string{"INBOX"}, accounts)
	model, _ := fi.Update(tea.WindowSizeMsg{Width: 200, Height: 60})
	fi = model.(*FolderInbox)

	initialInboxWidth := fi.inbox.list.Width()

	// Toggle sidebar hidden
	model, _ = fi.Update(tea.KeyPressMsg{Code: 'F', Text: "F"})
	fi = model.(*FolderInbox)

	if !fi.hideSidebar {
		t.Error("expected sidebar to be hidden after pressing 'F'")
	}

	hiddenSidebarInboxWidth := fi.inbox.list.Width()
	if hiddenSidebarInboxWidth <= initialInboxWidth {
		t.Errorf("expected inbox width to increase when sidebar is hidden, got %d <= %d", hiddenSidebarInboxWidth, initialInboxWidth)
	}

	// Toggle back
	model, _ = fi.Update(tea.KeyPressMsg{Code: 'F', Text: "F"})
	fi = model.(*FolderInbox)

	if fi.hideSidebar {
		t.Error("expected sidebar to be visible after pressing 'F' again")
	}

	if fi.inbox.list.Width() != initialInboxWidth {
		t.Errorf("expected inbox width to return to %d, got %d", initialInboxWidth, fi.inbox.list.Width())
	}
}

func TestFolderInbox_ToggleLayout_Vertical(t *testing.T) {
	accounts := []config.Account{
		{ID: "account-1", Email: "host.example.com"},
	}
	fi := NewFolderInbox([]string{"INBOX"}, accounts)
	fi.SetLayout(config.LayoutVertical)
	fi.SetEnableQuickToggle(true)
	fi.splitActive = false

	// Toggle Shift+L
	fi.Update(tea.KeyPressMsg{Code: 'L', Text: "L"})

	if !fi.splitActive {
		t.Error("expected splitActive to be true after pressing Shift+L in Vertical mode")
	}

	// Toggle back
	fi.Update(tea.KeyPressMsg{Code: 'L', Text: "L"})

	if fi.splitActive {
		t.Error("expected splitActive to be false after pressing Shift+L again")
	}
}

func TestFolderInbox_ToggleLayout_Horizontal(t *testing.T) {
	accounts := []config.Account{
		{ID: "account-1", Email: "host.example.com"},
	}
	fi := NewFolderInbox([]string{"INBOX"}, accounts)
	fi.SetLayout(config.LayoutHorizontal)
	fi.SetEnableQuickToggle(true)
	fi.splitActive = true

	// Toggle Shift+L
	fi.Update(tea.KeyPressMsg{Code: 'L', Text: "L"})

	// Should NOT change splitActive in Horizontal mode
	if !fi.splitActive {
		t.Error("expected splitActive to remain true in Horizontal mode")
	}
}
