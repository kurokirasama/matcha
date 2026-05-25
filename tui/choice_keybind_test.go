package tui

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/floatpane/matcha/config"
)

func TestChoiceMainMenuKeybinds(t *testing.T) {
	tempDir := t.TempDir()
	t.Setenv("HOME", tempDir)

	tests := []struct {
		name     string
		key      string
		enabled  bool
		wantType interface{}
	}{
		{"View (v) Enabled", "v", true, GoToInboxMsg{}},
		{"View (v) Disabled", "v", false, nil},
		{"Compose (c) Enabled", "c", true, GoToSendMsg{}},
		{"Marketplace (p) Enabled", "p", true, GoToMarketplaceMsg{}},
		{"Marketplace (p) Disabled", "p", false, nil},
		{"Settings (s) Enabled", "s", true, GoToSettingsMsg{}},
		{"Settings (s) Disabled", "s", false, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &config.Config{
				EnableMainMenuKeybinds: tt.enabled,
			}
			config.SaveConfig(cfg)
			
			m := NewChoice()
			msg := tea.KeyPressMsg{Text: tt.key} 
			_, cmd := m.Update(msg)
			
			if tt.wantType == nil {
				if cmd != nil {
					res := cmd()
					if res != nil {
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
			case GoToInboxMsg:
				if _, ok := res.(GoToInboxMsg); !ok {
					t.Errorf("Expected GoToInboxMsg, got %T", res)
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
