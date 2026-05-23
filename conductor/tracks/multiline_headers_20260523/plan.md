# Implementation Plan: Multi-line Email Headers

## Phase 1: Configuration Extension
- [ ] Task: Add `EnableMultilineHeaders` boolean to `Config`, `secureDiskConfig`, and `diskConfig` structs in `config/config.go`.
- [ ] Task: Update `SaveConfig` and `LoadConfig` to serialize/deserialize the new field.
- [ ] Task: Write a regression test in `config/config_test.go` to ensure `EnableMultilineHeaders` persists correctly across saves without data loss.
- [ ] Task: Conductor - User Manual Verification 'Config Extension' (Protocol in workflow.md)

## Phase 2: Settings UI Integration
- [ ] Task: Add localization keys for `settings_general.enable_multiline_headers` in `i18n/locales/en.json` and `es.json`.
- [ ] Task: Add a new option to `buildGeneralOptions()` in `tui/settings_general.go` to toggle `EnableMultilineHeaders`.
- [ ] Task: Update the switch statement in `updateGeneral()` to handle the new option index and trigger `ConfigSavedMsg`.
- [ ] Task: Conductor - User Manual Verification 'Settings UI' (Protocol in workflow.md)

## Phase 3: EmailView Rendering Refactor
- [ ] Task: Extract the current single-line header rendering logic in `tui/email_view.go` into a `renderSingleLineHeader()` helper.
- [ ] Task: Implement a new `renderMultiLineHeader()` helper that formats From, To, Cc, Bcc, and Subject on separate lines, skipping empty fields.
- [ ] Task: Update the `View()` method in `EmailView` to dynamically choose between the two rendering helpers based on `m.config.EnableMultilineHeaders` (passing config via constructor or messages if needed).
- [ ] Task: Conductor - User Manual Verification 'Header Rendering' (Protocol in workflow.md)

## Phase 4: Final Verification
- [ ] Task: Run `make build` and perform end-to-end manual verification of both ON and OFF states.
- [ ] Task: Verify that resizing the window recalculates the header block height correctly if multi-line is enabled.
- [ ] Task: Conductor - User Manual Verification 'Final Quality Gate' (Protocol in workflow.md)
