package config

import (
	"encoding/json"
	"testing"
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
