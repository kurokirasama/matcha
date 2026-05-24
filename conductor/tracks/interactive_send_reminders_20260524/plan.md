# Implementation Plan: Interactive Send Reminders

## Phase 1: Configuration and State
- [ ] Task: Write failing tests for new configuration settings (`EnableSubjectReminder`, `EnableAttachmentReminder`).
- [ ] Task: Implement settings in `Config` structs (`config/config.go`) and persistence logic.
- [ ] Task: Add toggles to the settings UI (`tui/settings_general.go`) and localization files.
- [ ] Task: Conductor - User Manual Verification 'Configuration and State' (Protocol in workflow.md)

## Phase 2: Reminder Logic Extraction
- [ ] Task: Write failing tests for reminder detection logic based on subject content and attachment keywords.
- [ ] Task: Implement `hasEmptySubject()` and `hasMissingAttachment()` evaluation methods in `tui/composer.go`.
- [ ] Task: Conductor - User Manual Verification 'Reminder Logic Extraction' (Protocol in workflow.md)

## Phase 3: Composer UI Integration
- [ ] Task: Update the `View()` logic in `tui/composer.go` to display active warnings statically in the main composer view.
- [ ] Task: Ensure the warnings are styled using the active theme's accent color (avoiding hardcoded reds).
- [ ] Task: Update the enhanced exit dialog view to also include the active warnings if triggered.
- [ ] Task: Conductor - User Manual Verification 'Composer UI Integration' (Protocol in workflow.md)

## Phase 4: Pre-Send Interception & Warning Dialog
- [ ] Task: Add a new `confirmingSendWithWarnings` boolean state to the `Composer` struct.
- [ ] Task: Modify `handleSend()` to intercept the send action and trigger `confirmingSendWithWarnings = true` if warnings exist.
- [ ] Task: Update the `Update()` loop to process `[s]end`, `[a]bort`, `[d]ave`, `[c]ancel` when `confirmingSendWithWarnings` is true. Ensure `[s]` bypasses the checks.
- [ ] Task: Update the `View()` logic to render the pre-send warning dialog (identical options to the exit dialog) when `confirmingSendWithWarnings` is true.
- [ ] Task: Conductor - User Manual Verification 'Pre-Send Interception' (Protocol in workflow.md)

## Phase 5: Final Quality Gate
- [ ] Task: Run `make build` and perform end-to-end manual testing of all triggers, overrides, and dialog paths.
- [ ] Task: Conductor - User Manual Verification 'Final Quality Gate' (Protocol in workflow.md)