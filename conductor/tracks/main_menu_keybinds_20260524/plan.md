# Implementation Plan: Main Screen Quick Keys

## Phase 1: Configuration & Persistence [checkpoint: 2375984]
- [x] Task: Write failing unit tests for `EnableMainMenuKeybinds` in `config/config_test.go`. d0117ae
- [x] Task: Implement `EnableMainMenuKeybinds` boolean in `Config`, `secureDiskConfig`, and `diskConfig` structs in `config/config.go`. d0117ae
- [x] Task: Update `SaveConfig` and `LoadConfig` in `config/config.go` to handle the new field. d0117ae
- [x] Task: Conductor - User Manual Verification 'Configuration & Persistence' (Protocol in workflow.md) d0117ae

## Phase 2: Settings UI & Localization [checkpoint: d7de6c0]
- [x] Task: Add localization keys for `settings_general.enable_main_menu_keybinds` in `i18n/locales/en.json` and `es.json`. a50e2eb
- [x] Task: Add a new toggle to `buildGeneralOptions()` in `tui/settings_general.go` with the label "Enable Main Menu Keybinds". a50e2eb
- [x] Task: Update `updateGeneral()` in `tui/settings_general.go` to handle the setting index and save the configuration. a50e2eb
- [x] Task: Conductor - User Manual Verification 'Settings UI & Localization' (Protocol in workflow.md) a50e2eb

## Phase 3: Choice Screen Keybind Implementation
- [x] Task: Write failing UI unit tests simulating key presses `v`, `c`, `p`, `s` in the Choice view with the setting enabled/disabled. 91bb31f
- [x] Task: Update `Update()` in `tui/choice.go` to intercept `v`, `c`, `p`, and `s` keys when `EnableMainMenuKeybinds` is true. 91bb31f
- [x] Task: Map the keys to their respective messages. 91bb31f
- [~] Task: Conductor - User Manual Verification 'Choice Screen Keybind Implementation' (Protocol in workflow.md)

## Phase 4: Dynamic Help Bar
- [x] Task: Update `View()` in `tui/choice.go` to conditionally include the quick key hints in the help bar. 91bb31f
- [x] Task: Ensure the hints only appear when `EnableMainMenuKeybinds` is enabled. 91bb31f
- [~] Task: Conductor - User Manual Verification 'Dynamic Help Bar' (Protocol in workflow.md)

## Phase 5: Final Quality Gate
- [ ] Task: Run `make build` and perform end-to-end manual testing of all triggers and persistence.
- [ ] Task: Verify that keybinds are inactive when the setting is OFF or when focus is elsewhere.
- [ ] Task: Conductor - User Manual Verification 'Final Quality Gate' (Protocol in workflow.md)