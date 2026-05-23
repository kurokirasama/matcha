# Implementation Plan: Detailed Email Notifications

## Phase 1: Configuration Extension
- [ ] Task: Add `EnableExpandedNotifications` boolean to `Config`, `secureDiskConfig`, and `diskConfig` structs in `config/config.go`.
- [ ] Task: Update `SaveConfig` and `LoadConfig` to serialize/deserialize the new field safely.
- [ ] Task: Write a regression test in `config/config_test.go` to ensure `EnableExpandedNotifications` persists correctly.
- [ ] Task: Conductor - User Manual Verification 'Config Extension' (Protocol in workflow.md)

## Phase 2: Settings UI Integration
- [ ] Task: Add localization keys for `settings_general.enable_expanded_notifications` in `i18n/locales/en.json` and `es.json`.
- [ ] Task: Add a new toggle to `buildGeneralOptions()` in `tui/settings_general.go`.
- [ ] Task: Update `updateGeneral()` in `tui/settings_general.go` to handle the new setting index.
- [ ] Task: Conductor - User Manual Verification 'Settings UI' (Protocol in workflow.md)

## Phase 3: Notification Logic Refactor
- [ ] Task: Identify the notification triggering logic (likely in `daemon/` or `notify/` packages).
- [ ] Task: Update the notification dispatch payload to include `From` and `Subject` details when appropriate.
- [ ] Task: Modify the notification construction logic to conditionally format the title and body based on the `EnableExpandedNotifications` configuration flag.
- [ ] Task: Write or update unit tests to verify the conditional notification formatting.
- [ ] Task: Conductor - User Manual Verification 'Notification Logic' (Protocol in workflow.md)

## Phase 4: Final Verification
- [ ] Task: Run `make build` and perform end-to-end manual verification of notifications in both ON and OFF states.
- [ ] Task: Conductor - User Manual Verification 'Final Quality Gate' (Protocol in workflow.md)