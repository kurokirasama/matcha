# Implementation Plan: Native Toggle Read Status

## Phase 1: Keybinding & Action Definition
- [ ] Task: Add `ToggleRead` to `InboxKeys` and `EmailKeys` in `config/keybinds.go`.
- [ ] Task: Add default `u` mapping for `toggle_read` in `config/default_keybinds.json`.
- [ ] Task: Conductor - User Manual Verification 'Keybinding Setup' (Protocol in workflow.md)

## Phase 2: Core Logic Implementation
- [ ] Task: Implement `ToggleRead` logic in the TUI model (likely `tui/inbox.go` and `tui/email_view.go`).
- [ ] Task: Ensure the daemon is notified of the status change to sync with the backend.
- [ ] Task: Conductor - User Manual Verification 'Core Implementation' (Protocol in workflow.md)

## Phase 3: Verification
- [ ] Task: Verify toggling works in both list and view modes.
- [ ] Task: Verify status persists after refreshing the inbox.
