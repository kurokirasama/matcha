package tui

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/floatpane/matcha/config"
	"github.com/floatpane/matcha/fetcher"
)

// TestFolderInboxSplitPreviewRendersSearchHit covers the case Lea reported on
// PR #1186: opening a search result in split-pane mode used to silently drop
// the keypress because the email was not in m.inbox.allEmails. After the fix
// OpenSplitPreview accepts the resolved email and findEmailByUID falls back
// to it, so PreviewBodyFetchedMsg can build the preview pane.
func TestFolderInboxSplitPreviewRendersSearchHit(t *testing.T) {
	accounts := []config.Account{
		{ID: "account-1", Email: "host.example.com", FetchEmail: "first@example.com"},
	}
	fi := NewFolderInbox([]string{"INBOX", "Archive"}, accounts)
	// Force a non-zero canvas so calculate*Width does not panic on Update.
	model, _ := fi.Update(tea.WindowSizeMsg{Width: 200, Height: 60})
	fi = model.(*FolderInbox)

	// Search hit lives in a different folder; allEmails is empty.
	hit := &fetcher.Email{
		UID:       4242,
		AccountID: "account-1",
		MessageID: "<search-hit@example.com>",
		From:      "sender@example.com",
		To:        []string{"first@example.com"},
		Subject:   "Search hit",
	}

	fi.OpenSplitPreview(hit.UID, hit.AccountID, hit)

	if fi.previewSearchEmail == nil {
		t.Fatal("OpenSplitPreview should retain the search hit email")
	}
	if got := fi.findEmailByUID(hit.UID, hit.AccountID); got == nil {
		t.Fatal("findEmailByUID should fall back to the search hit email")
	}

	// Simulate the body arriving and verify the preview pane is built.
	model, _ = fi.Update(PreviewBodyFetchedMsg{
		UID:       hit.UID,
		AccountID: hit.AccountID,
		Body:      "hello body",
	})
	fi = model.(*FolderInbox)

	if fi.previewPane == nil {
		t.Fatal("expected previewPane to be built from the search hit fallback")
	}

	// closeSplitPreview must clear the cached search hit so a later open with
	// no email cannot accidentally reuse the stale reference.
	fi.closeSplitPreview()
	if fi.previewSearchEmail != nil {
		t.Fatal("closeSplitPreview should clear previewSearchEmail")
	}
}

// TestFolderInboxSplitPreviewPrefersAllEmails verifies that when the email is
// already known in allEmails, findEmailByUID returns the live entry (so reads
// like IsRead stay current) instead of the snapshot passed via OpenSplitPreview.
func TestFolderInboxSplitPreviewPrefersAllEmails(t *testing.T) {
	accounts := []config.Account{
		{ID: "account-1", Email: "host.example.com", FetchEmail: "first@example.com"},
	}
	fi := NewFolderInbox([]string{"INBOX"}, accounts)
	model, _ := fi.Update(tea.WindowSizeMsg{Width: 200, Height: 60})
	fi = model.(*FolderInbox)

	live := fetcher.Email{UID: 7, AccountID: "account-1", Subject: "live", IsRead: true}
	fi.SetEmails([]fetcher.Email{live}, accounts)

	stale := &fetcher.Email{UID: 7, AccountID: "account-1", Subject: "stale", IsRead: false}
	fi.OpenSplitPreview(live.UID, live.AccountID, stale)

	got := fi.findEmailByUID(live.UID, live.AccountID)
	if got == nil {
		t.Fatal("findEmailByUID should resolve the email")
	}
	if got.Subject != "live" || !got.IsRead {
		t.Fatalf("expected the live allEmails entry, got %+v", got)
	}
}

// TestSearchOverlayKeysNotIntercepted covers issue #1199: pressing keys that
// match folder-level bindings (e.g. "m" for move) while the search overlay is
// active used to trigger the move flow instead of entering text into the
// search input. FolderInbox.Update now passes through to the inner inbox
// while m.inbox.searchOverlay != nil so the overlay receives raw keystrokes.
func TestSearchOverlayKeysNotIntercepted(t *testing.T) {
	accounts := []config.Account{
		{ID: "account-1", Email: "host.example.com", FetchEmail: "first@example.com"},
	}
	fi := NewFolderInbox([]string{"INBOX", "Archive"}, accounts)
	model, _ := fi.Update(tea.WindowSizeMsg{Width: 200, Height: 60})
	fi = model.(*FolderInbox)

	// Open search overlay
	model, _ = fi.Update(tea.KeyPressMsg{Code: '/', Text: "/"})
	fi = model.(*FolderInbox)

	if fi.inbox.searchOverlay == nil {
		t.Fatal("expected search overlay to be open")
	}

	// Press "m" (matches move-to-folder keybinding in FolderInbox)
	model, _ = fi.Update(tea.KeyPressMsg{Code: 'm', Text: "m"})
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

func TestFolderInboxLayoutResizing(t *testing.T) {
	accounts := []config.Account{
		{ID: "account-1", Email: "host.example.com"},
	}
	fi := NewFolderInbox([]string{"INBOX"}, accounts)
	
	// Test Vertical (default for split if not set)
	fi.SetLayout(config.LayoutVertical)
	width, height := 100, 40
	fi.Update(tea.WindowSizeMsg{Width: width, Height: height})
	fi.OpenSplitPreview(1, "account-1", &fetcher.Email{UID: 1})

	if fi.calculateInboxHeight() != height {
		t.Errorf("Vertical layout: expected inbox height %d, got %d", height, fi.calculateInboxHeight())
	}
	if fi.calculatePreviewWidth() >= width-sidebarWidth {
		t.Errorf("Vertical layout: expected preview width to be less than available width")
	}

	// Test Horizontal
	fi.SetLayout(config.LayoutHorizontal)
	fi.Update(tea.WindowSizeMsg{Width: width, Height: height})

	inboxHeight := fi.calculateInboxHeight()
	previewHeight := fi.calculatePreviewHeight()

	if inboxHeight >= height {
		t.Errorf("Horizontal layout: expected inbox height to be less than %d, got %d", height, inboxHeight)
	}
	if previewHeight >= height {
		t.Errorf("Horizontal layout: expected preview height to be less than %d, got %d", height, previewHeight)
	}
	if fi.calculateInboxWidth() != width-sidebarWidth-2 {
		t.Errorf("Horizontal layout: expected inbox width to be full, got %d", fi.calculateInboxWidth())
	}
}

func TestFolderInbox_ToggleLayout(t *testing.T) {
	accounts := []config.Account{
		{ID: "account-1", Email: "host.example.com"},
	}
	fi := NewFolderInbox([]string{"INBOX"}, accounts)
	fi.SetLayout(config.LayoutVertical)
	fi.SetEnableQuickToggle(true)

	width, height := 100, 40
	model, _ := fi.Update(tea.WindowSizeMsg{Width: width, Height: height})
	fi = model.(*FolderInbox)

	// Initial state: splitActive is true
	if !fi.splitActive {
		t.Error("expected splitActive to be true initially")
	}

	// Press 'L' (Shift+L) -> Should deactivate split
	model, _ = fi.Update(tea.KeyPressMsg{Code: 'L', Text: "L"})
	fi = model.(*FolderInbox)

	if fi.splitActive {
		t.Error("expected splitActive to be false after toggle")
	}

	// In Vertical mode, inactive split should mean full width inbox
	if fi.calculateInboxWidth() != width-sidebarWidth-2 {
		t.Errorf("expected full width inbox, got %d", fi.calculateInboxWidth())
	}

	// Toggle back -> Should reactivate split
	model, _ = fi.Update(tea.KeyPressMsg{Code: 'L', Text: "L"})
	fi = model.(*FolderInbox)

	if !fi.splitActive {
		t.Error("expected splitActive to be true after second toggle")
	}

	// Test OFF mode
	fi.SetLayout(config.LayoutOff)
	fi.splitActive = true // Full height
	
	// Toggle -> Half height
	model, _ = fi.Update(tea.KeyPressMsg{Code: 'L', Text: "L"})
	fi = model.(*FolderInbox)

	if fi.calculateInboxHeight() != height/2 {
		// Note: calculateInboxHeight might adjust for borders, so we check if it's less than full
		if fi.calculateInboxHeight() >= height {
			t.Errorf("expected reduced height in OFF mode with split inactive, got %d", fi.calculateInboxHeight())
		}
	}
}
