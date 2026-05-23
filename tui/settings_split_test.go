package tui

import (
	"testing"
	"github.com/floatpane/matcha/config"
)

func TestSettingsSplitViewOption(t *testing.T) {
	cfg := &config.Config{
		Layout: config.LayoutOff,
	}
	settings := NewSettings(cfg)
	
	options := settings.buildGeneralOptions()
	found := false
	for _, opt := range options {
		if opt.labelKey == "settings_general.split_view" {
			found = true
			break
		}
	}
	
	if !found {
		t.Errorf("expected settings_general.split_view option not found")
	}
}

func TestSettingsQuickToggleOption(t *testing.T) {
	cfg := &config.Config{
		EnableQuickToggle: false,
	}
	settings := NewSettings(cfg)
	
	options := settings.buildGeneralOptions()
	found := false
	for _, opt := range options {
		if opt.labelKey == "settings_general.layout_quick_toggle" {
			found = true
			break
		}
	}
	
	if !found {
		t.Errorf("expected settings_general.layout_quick_toggle option not found")
	}
}
