package config

import (
	"encoding/json"
	"testing"

	"github.com/zalando/go-keyring"
)

func TestLayoutModeSerialization(t *testing.T) {
	testCases := []struct {
		name     string
		mode     LayoutMode
		expected string
	}{
		{"Off", LayoutOff, `"off"`},
		{"Vertical", LayoutVertical, `"vertical"`},
		{"Horizontal", LayoutHorizontal, `"horizontal"`},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := json.Marshal(tc.mode)
			if err != nil {
				t.Fatalf("Marshal failed: %v", err)
			}
			if string(data) != tc.expected {
				t.Errorf("Marshal = %s, want %s", string(data), tc.expected)
			}

			var decoded LayoutMode
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatalf("Unmarshal failed: %v", err)
			}
			if decoded != tc.mode {
				t.Errorf("Unmarshal = %s, want %s", decoded, tc.mode)
			}
		})
	}
}

func TestConfigLayoutField(t *testing.T) {
	c := Config{
		Layout: LayoutHorizontal,
	}

	data, err := json.Marshal(c)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded Config
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if decoded.Layout != LayoutHorizontal {
		t.Errorf("Decoded layout = %s, want %s", decoded.Layout, LayoutHorizontal)
	}
}

func TestConfigLayoutRoundTrip(t *testing.T) {
	keyring.MockInit()
	tempDir := t.TempDir()
	t.Setenv("HOME", tempDir)

	expectedConfig := &Config{
		Accounts: []Account{
			{ID: "test", Email: "test@example.com"},
		},
		Layout: LayoutHorizontal,
	}

	if err := SaveConfig(expectedConfig); err != nil {
		t.Fatalf("SaveConfig failed: %v", err)
	}

	loadedConfig, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if loadedConfig.Layout != expectedConfig.Layout {
		t.Errorf("Loaded layout = %s, want %s", loadedConfig.Layout, expectedConfig.Layout)
	}
}
