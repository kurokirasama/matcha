# Implementation Plan: Native Inbox Compose Keybind

## Phase 1: Configuration
- [ ] Task: Add `Compose` to `InboxKeys` in `config/keybinds.go`.
- [ ] Task: Add default `c` mapping in `config/default_keybinds.json`.
- [ ] Task: Conductor - User Manual Verification 'Config' (Protocol in workflow.md)

## Phase 2: TUI Implementation
- [ ] Task: Add keypress handler for `Compose` in `tui/inbox.go`.
- [ ] Task: Update help bindings to display the new shortcut.
- [ ] Task: Conductor - User Manual Verification 'Implementation' (Protocol in workflow.md)

## Phase 3: Verification
- [ ] Task: Confirm `c` triggers the composer.
