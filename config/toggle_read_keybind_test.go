package config

import (
	"testing"
)

func TestKeybinds_ToggleReadDefaultMapping(t *testing.T) {
	// Verify that the default mapping for toggle_read is 'u'
	kb := defaultKeybinds()
	
	if kb.Inbox.ToggleRead != "u" {
		t.Errorf("Default Inbox.ToggleRead should be 'u', got %q", kb.Inbox.ToggleRead)
	}
	if kb.Email.ToggleRead != "u" {
		t.Errorf("Default Email.ToggleRead should be 'u', got %q", kb.Email.ToggleRead)
	}
}

func TestKeybinds_ToggleSidebarDefaultMapping(t *testing.T) {
	// Verify that the default mapping for toggle_sidebar is 'F'
	kb := defaultKeybinds()
	
	if kb.Inbox.ToggleSidebar != "F" {
		t.Errorf("Default Inbox.ToggleSidebar should be 'F', got %q", kb.Inbox.ToggleSidebar)
	}
}
