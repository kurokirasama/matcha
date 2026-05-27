package config

import (
	"encoding/json"
	"testing"

	"github.com/zalando/go-keyring"
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

func TestConfigEnableQuickToggleRoundTrip(t *testing.T) {
	keyring.MockInit()
	tempDir := t.TempDir()
	t.Setenv("HOME", tempDir)

	expectedConfig := &Config{
		Accounts: []Account{
			{ID: "test", Email: "test@example.com"},
		},
		EnableQuickToggle: true,
	}

	if err := SaveConfig(expectedConfig); err != nil {
		t.Fatalf("SaveConfig failed: %v", err)
	}

	loadedConfig, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if loadedConfig.EnableQuickToggle != expectedConfig.EnableQuickToggle {
		t.Errorf("Loaded EnableQuickToggle = %v, want %v", loadedConfig.EnableQuickToggle, expectedConfig.EnableQuickToggle)
	}
}

func TestConfigSplitActiveRoundTrip(t *testing.T) {
	keyring.MockInit()
	tempDir := t.TempDir()
	t.Setenv("HOME", tempDir)

	expectedConfig := &Config{
		Accounts: []Account{
			{ID: "test", Email: "test@example.com"},
		},
		SplitActive: true,
	}

	if err := SaveConfig(expectedConfig); err != nil {
		t.Fatalf("SaveConfig failed: %v", err)
	}

	loadedConfig, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if loadedConfig.SplitActive != expectedConfig.SplitActive {
		t.Errorf("Loaded SplitActive = %v, want %v", loadedConfig.SplitActive, expectedConfig.SplitActive)
	}
}
