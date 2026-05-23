package config

import (
	"encoding/json"
	"testing"
)

func TestConfigEnableQuickToggleField(t *testing.T) {
	c := Config{
		EnableQuickToggle: true,
	}

	data, err := json.Marshal(c)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded Config
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if !decoded.EnableQuickToggle {
		t.Errorf("Decoded EnableQuickToggle = %v, want true", decoded.EnableQuickToggle)
	}
}
