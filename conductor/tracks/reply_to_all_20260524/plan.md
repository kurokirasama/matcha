# Implementation Plan: Reply to All Feature

## Phase 1: Configuration & Settings UI
- [ ] Task: Write failing unit tests for `EnableReplyToAll` in `config/config_test.go`.
- [ ] Task: Implement `EnableReplyToAll` boolean in `Config`, `secureDiskConfig`, and `diskConfig` structs in `config/config.go` with persistence logic.
- [ ] Task: Add localization keys for the new setting in `i18n/locales/en.json` and `es.json`.
- [ ] Task: Implement the toggle in `tui/settings_general.go` within `buildGeneralOptions()` and `updateGeneral()`.
- [ ] Task: Conductor - User Manual Verification 'Configuration & Settings UI' (Protocol in workflow.md)

## Phase 2: Composer Logic Integration
- [ ] Task: Write failing tests for parsing "Reply-To", "To", and "Cc" headers to correctly gather "Reply to All" recipients (excluding the current user).
- [ ] Task: Implement "Reply to All" recipient gathering logic.
- [ ] Task: Conductor - User Manual Verification 'Composer Logic Integration' (Protocol in workflow.md)

## Phase 3: Email View Keybind & UI Integration
- [ ] Task: Update the `Update` loop in `tui/email_view.go` (or wherever the email visualization screen logic resides) to intercept the `A` (`shift+a`) keybind when `EnableReplyToAll` is enabled.
- [ ] Task: Implement the transition to the composer view, pre-populated with the generated "Reply to All" recipients.
- [ ] Task: Update the `View` rendering logic of the email visualization screen to conditionally show the `A: reply all` hint based on the configuration.
- [ ] Task: Conductor - User Manual Verification 'Email View Keybind & UI Integration' (Protocol in workflow.md)

## Phase 4: Final Verification
- [ ] Task: Run `make build` and perform end-to-end manual verification of the "Reply to All" flow.
- [ ] Task: Conductor - User Manual Verification 'Final Quality Gate' (Protocol in workflow.md)