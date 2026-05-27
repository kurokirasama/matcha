package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"github.com/floatpane/matcha/config"
	"github.com/floatpane/matcha/i18n"
)

type generalOption struct {
	labelKey     string
	value        string
	tip          string
	disabled     bool
	isAccountSig bool
	accountID    string
}

func getLayoutLabel(mode config.LayoutMode) string {
	switch mode {
	case config.LayoutOff:
		return t("settings_general.split_view_off")
	case config.LayoutVertical:
		return t("settings_general.split_view_vertical")
	case config.LayoutHorizontal:
		return t("settings_general.split_view_horizontal")
	default:
		return t("settings_general.split_view_off")
	}
}

func (m *Settings) buildGeneralOptions() []generalOption {
	opts := []generalOption{
		{"settings_general.disable_images", onOff(m.cfg.DisableImages), "Prevent images from loading automatically in emails.", false, false, ""},
		{"settings_general.hide_tips", onOff(m.cfg.HideTips), "Hide helpful hints displayed at the bottom of the screen.", false, false, ""},
		{"settings_general.disable_notifications", onOff(m.cfg.DisableNotifications), "Turn off desktop notifications for new mail.", false, false, ""},
		{"settings_general.split_view", getLayoutLabel(m.cfg.Layout), "Orientation of the email preview pane.", false, false, ""},
		{"settings_general.layout_quick_toggle", onOff(m.cfg.EnableQuickToggle), "Enable Shift+L shortcut to cycle layout modes.", m.cfg.Layout == config.LayoutHorizontal, false, ""},
		{"settings_general.enable_threaded", onOff(m.cfg.EnableThreaded), "Group emails into conversations by reply chain. Per-folder overrides are kept.", false, false, ""},
		{"settings_general.enable_detailed_dates", onOff(m.cfg.EnableDetailedDates), "Show detailed inbox dates.", false, false, ""},
		{"settings_general.enable_main_menu_keybinds", onOff(m.cfg.EnableMainMenuKeybinds), "Enable single-key shortcuts (v, c, p, s) on the main screen.", false, false, ""},
		{"settings_general.enable_enhanced_composer_exit", onOff(m.cfg.EnableEnhancedComposerExit), "Show a rich confirmation dialog with quick keys when exiting the composer.", false, false, ""},
		{"settings_general.spellcheck", onOff(!m.cfg.DisableSpellcheck), "Underline misspelled words while composing.", false, false, ""},
		{"settings_general.spell_suggestions", onOff(!m.cfg.DisableSpellSuggestions), "Show suggestion popup for misspelled words.", false, false, ""},
		{"settings_general.date_format", getDateFormatLabel(m.cfg.DateFormat), "Change how dates and times are displayed.", false, false, ""},
		{"settings_general.language", getLanguageLabel(m.cfg.GetLanguage()), "Change the interface language. Changes apply instantly.", false, false, ""},
		{"settings_general.signature", getSignatureStatus(), "Configure the global signature appended to your outgoing emails.", false, false, ""},
	}

	for _, acc := range m.cfg.Accounts {
		status := t("settings_general.signature_not_configured")
		accCopy := acc // capture for pointer safety
		if config.HasAccountSignature(&accCopy) {
			status = t("settings_general.signature_configured")
		}
		opts = append(opts, generalOption{
			labelKey:     fmt.Sprintf("Signature (%s)", acc.Email),
			value:        status,
			tip:          fmt.Sprintf("Configure the signature for %s", acc.Email),
			isAccountSig: true,
			accountID:    acc.ID,
		})
	}

	return opts
}

func (m *Settings) updateGeneral(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	opts := m.buildGeneralOptions()

	switch msg.String() {
	case "up", "k":
		m.generalCursor = (m.generalCursor - 1 + len(opts)) % len(opts)
	case keyDown, "j":
		m.generalCursor = (m.generalCursor + 1) % len(opts)
	case keyEnter, "space", keyRight, "l":
		if m.generalCursor < len(opts) {
			saved := false
			switch m.generalCursor {
			case 0: // Image Display
				m.cfg.DisableImages = !m.cfg.DisableImages
				_ = config.SaveConfig(m.cfg)
				saved = true
			case 1: // Contextual Tips
				m.cfg.HideTips = !m.cfg.HideTips
				_ = config.SaveConfig(m.cfg)
				saved = true
			case 2: // Desktop Notifications
				m.cfg.DisableNotifications = !m.cfg.DisableNotifications
				_ = config.SaveConfig(m.cfg)
				saved = true
			case 3: // Split View
				switch m.cfg.Layout {
				case config.LayoutOff:
					m.cfg.Layout = config.LayoutVertical
					m.cfg.EnableSplitPane = true
				case config.LayoutVertical:
					m.cfg.Layout = config.LayoutHorizontal
					m.cfg.EnableSplitPane = true
					// Force off quick toggle when entering horizontal
					m.cfg.EnableQuickToggle = false
				case config.LayoutHorizontal:
					m.cfg.Layout = config.LayoutOff
					m.cfg.EnableSplitPane = false
				default:
					m.cfg.Layout = config.LayoutOff
					m.cfg.EnableSplitPane = false
				}
				_ = config.SaveConfig(m.cfg)
				saved = true
			case 4: // Layout Quick Toggle
				if m.cfg.Layout != config.LayoutHorizontal {
					m.cfg.EnableQuickToggle = !m.cfg.EnableQuickToggle
					_ = config.SaveConfig(m.cfg)
					saved = true
				} else {
					return m, func() tea.Msg {
						return PluginNotifyMsg{
							Message:  t("settings_general.quick_toggle_unavailable_horizontal"),
							Duration: 2,
						}
					}
				}
			case 5: // Threaded Conversation View
				m.cfg.EnableThreaded = !m.cfg.EnableThreaded
				_ = config.SaveConfig(m.cfg)
				saved = true
			case 6: // Detailed Dates
				m.cfg.EnableDetailedDates = !m.cfg.EnableDetailedDates
				_ = config.SaveConfig(m.cfg)
				saved = true
			case 7: // Main Menu Keybinds
				m.cfg.EnableMainMenuKeybinds = !m.cfg.EnableMainMenuKeybinds
				_ = config.SaveConfig(m.cfg)
				saved = true
			case 8: // Enhanced Composer Exit
				m.cfg.EnableEnhancedComposerExit = !m.cfg.EnableEnhancedComposerExit
				_ = config.SaveConfig(m.cfg)
				saved = true
			case 9: // Spellcheck
				m.cfg.DisableSpellcheck = !m.cfg.DisableSpellcheck
				_ = config.SaveConfig(m.cfg)
				saved = true
			case 10: // Spell Suggestions
				m.cfg.DisableSpellSuggestions = !m.cfg.DisableSpellSuggestions
				_ = config.SaveConfig(m.cfg)
				saved = true
			case 11: // Date Format
				switch m.cfg.DateFormat {
				case config.DateFormatEU:
					m.cfg.DateFormat = config.DateFormatUS
				case config.DateFormatUS:
					m.cfg.DateFormat = config.DateFormatISO
				default: // or ISO
					m.cfg.DateFormat = config.DateFormatEU
				}
				_ = config.SaveConfig(m.cfg)
				saved = true
			case 12: // Language
				// Cycle through available languages
				langs := i18n.LanguageCodes()
				currentLang := m.cfg.GetLanguage()
				currentIdx := -1
				for i, lang := range langs {
					if lang == currentLang {
						currentIdx = i
						break
					}
				}
				nextIdx := (currentIdx + 1) % len(langs)
				m.cfg.Language = langs[nextIdx]
				_ = config.SaveConfig(m.cfg)
				// Apply language change immediately
				i18n.GetManager().SetLanguage(m.cfg.Language) //nolint:errcheck,gosec
				// Trigger full UI rebuild
				return m, tea.Batch(
					func() tea.Msg { return ConfigSavedMsg{} },
					func() tea.Msg { return LanguageChangedMsg{} },
				)
			case 13: // Edit Signature
				if msg.String() == "enter" || msg.String() == "right" || msg.String() == "l" {
					return m, func() tea.Msg { return GoToSignatureEditorMsg{} }
				}
			default:
				// Check for per-account signatures
				opt := opts[m.generalCursor]
				if opt.isAccountSig {
					if msg.String() == "enter" || msg.String() == "right" || msg.String() == "l" {
						return m, func() tea.Msg {
							return GoToSignatureEditorMsg{AccountID: opt.accountID}
						}
					}
				}
			}
			if saved {
				return m, func() tea.Msg { return ConfigSavedMsg{} }
			}
		}
	}
	return m, nil
}

func (m *Settings) viewGeneral() string {
	var b strings.Builder

	b.WriteString(titleStyle.Render("General Settings") + "\n\n")

	options := m.buildGeneralOptions()

	for i, opt := range options {
		cursor := "  "
		style := accountItemStyle
		if m.generalCursor == i {
			cursor = "> "
			style = selectedAccountItemStyle
		}

		label := opt.labelKey
		if !opt.isAccountSig {
			label = t(opt.labelKey)
		}

		val := opt.value
		if opt.disabled {
			val = fmt.Sprintf("\x1b[90m%s (Disabled)\x1b[m", val)
		}

		text := fmt.Sprintf("%s: %s", label, val)
		if opt.labelKey == "settings_general.signature" || opt.isAccountSig {
			text = fmt.Sprintf("%s (%s)", label, val)
		}

		b.WriteString(style.Render(cursor+text) + "\n")
	}

	b.WriteString("\n\n")

	if !m.cfg.HideTips && m.generalCursor < len(options) {
		b.WriteString(TipStyle.Render("Tip: " + options[m.generalCursor].tip))
	}

	return b.String()
}

func onOff(b bool) string {
	if b {
		return t("settings_general.on")
	}
	return t("settings_general.off")
}

func getDateFormatLabel(f string) string {
	if f == "" {
		f = config.DateFormatEU
	}
	switch f {
	case config.DateFormatUS:
		return "US (MM/DD/YYYY hh:MM AM)"
	case config.DateFormatISO:
		return "ISO (YYYY-MM-DD HH:MM)"
	default:
		return "EU (DD/MM/YYYY HH:MM)"
	}
}

func getSignatureStatus() string {
	if config.HasSignature() {
		return t("settings_general.signature_configured")
	}
	return t("settings_general.signature_not_configured")
}

func getLanguageLabel(langCode string) string {
	if locale, ok := i18n.GetLanguage(langCode); ok {
		return fmt.Sprintf("%s (%s)", locale.NativeName, locale.Code)
	}
	return langCode
}
