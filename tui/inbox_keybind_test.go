package tui

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/floatpane/matcha/config"
	"github.com/floatpane/matcha/fetcher"
)

func TestInboxMainMenuKeybinds(t *testing.T) {
	config.Keybinds.Inbox.Open = "enter"
	config.Keybinds.Inbox.Compose = "c"

	accounts := []config.Account{{ID: "acc-1", Email: "test@example.com"}}
	emails := []fetcher.Email{{UID: 1, AccountID: "acc-1", Subject: "Test"}}
	
	m := NewInbox(emails, accounts)

	tests := []struct {
		name     string
		key      string
		enabled  bool
		wantType interface{}
	}{
		{"View (v) Enabled", "v", true, ViewEmailMsg{}},
		{"View (v) Disabled", "v", false, nil},
		{"Compose (c) Enabled", "c", true, GoToSendMsg{}},
		{"Marketplace (p) Enabled", "p", true, GoToMarketplaceMsg{}},
		{"Marketplace (p) Disabled", "p", false, nil},
		{"Settings (s) Enabled", "s", true, GoToSettingsMsg{}},
		{"Settings (s) Disabled", "s", false, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.EnableMainMenuKeybinds = tt.enabled
			msg := tea.KeyPressMsg{Text: tt.key} 
			_, cmd := m.Update(msg)
			
			if tt.wantType == nil {
				if cmd != nil {
					res := cmd()
					if res != nil {
						// Filter out common maintenance messages
						switch res.(type) {
						case FetchMoreEmailsMsg:
							return
						}
						t.Errorf("Expected nil msg for key %s, got %T", tt.key, res)
					}
				}
				return
			}

			if cmd == nil {
				t.Fatalf("Expected cmd for key %s, got nil", tt.key)
			}
			res := cmd()
			if res == nil {
				t.Fatalf("Expected msg from cmd for key %s, got nil", tt.key)
			}

			switch tt.wantType.(type) {
			case ViewEmailMsg:
				if _, ok := res.(ViewEmailMsg); !ok {
					t.Errorf("Expected ViewEmailMsg, got %T", res)
				}
			case GoToSendMsg:
				if _, ok := res.(GoToSendMsg); !ok {
					t.Errorf("Expected GoToSendMsg, got %T", res)
				}
			case GoToMarketplaceMsg:
				if _, ok := res.(GoToMarketplaceMsg); !ok {
					t.Errorf("Expected GoToMarketplaceMsg, got %T", res)
				}
			case GoToSettingsMsg:
				if _, ok := res.(GoToSettingsMsg); !ok {
					t.Errorf("Expected GoToSettingsMsg, got %T", res)
				}
			}
		})
	}
}
