# Implementation Plan: Native Toggle Read Status

## Phase 1: Keybinding & Action Definition [checkpoint: 1721657]
- [x] Task: Add `ToggleRead` to `InboxKeys` and `EmailKeys` in `config/keybinds.go`. 76d3677
- [x] Task: Add default `u` mapping for `toggle_read` in `config/default_keybinds.json`. 76d3677
- [x] Task: Conductor - User Manual Verification 'Keybinding Setup' (Protocol in workflow.md) 76d3677

## Phase 2: Core Logic Implementation [checkpoint: a38c3cd]
- [x] Task: Implement `ToggleRead` logic in the TUI model (likely `tui/inbox.go` and `tui/email_view.go`). a38c3cd
- [x] Task: Ensure the daemon is notified of the status change to sync with the backend. a38c3cd
- [x] Task: Conductor - User Manual Verification 'Core Implementation' (Protocol in workflow.md) a38c3cd

## Phase 3: UI Refinement & Verification [checkpoint: 148f7b2]
- [x] Task: Update help bar in inbox and email views to show `u` keybinding. 148f7b2
- [x] Task: Implement exit-to-inbox behavior when toggling read status in email view. 148f7b2
- [x] Task: Verify toggling works in both list and view modes. 148f7b2
- [x] Task: Verify status persists after refreshing the inbox. 148f7b2
- [x] Task: Conductor - User Manual Verification 'UI Refinement' (Protocol in workflow.md) 148f7b2

## Phase: Review Fixes
- [x] Task: Apply review suggestions f3ae6a0
