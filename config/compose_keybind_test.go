package config

import (
	"encoding/json"
	"testing"
)

func TestKeybinds_ComposeExists(t *testing.T) {
	// Verify that Compose exists in InboxKeys struct
	kb := KeybindsConfig{}
	
	// Test JSON unmarshaling to ensure the field is recognized
	data := `{
		"inbox": { "compose": "c" }
	}`
	
	if err := json.Unmarshal([]byte(data), &kb); err != nil {
		t.Fatalf("Failed to unmarshal keybinds with compose: %v", err)
	}
	
	if kb.Inbox.Compose != "c" {
		t.Errorf("Inbox.Compose not set correctly, got %q", kb.Inbox.Compose)
	}
}
