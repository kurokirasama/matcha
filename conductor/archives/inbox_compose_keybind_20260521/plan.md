# Implementation Plan: Native Inbox Compose Keybind

## Phase 1: Configuration [checkpoint: 2bf17c9]
- [x] Task: Add `Compose` to `InboxKeys` in `config/keybinds.go`. d7bb401
- [x] Task: Add default `c` mapping in `config/default_keybinds.json`. 60810fb
- [x] Task: Conductor - User Manual Verification 'Config' (Protocol in workflow.md) 60810fb

## Phase 2: TUI Implementation [checkpoint: 96abe46]
- [x] Task: Add keypress handler for `Compose` in `tui/inbox.go`. 2ff1331
- [x] Task: Update help bindings to display the new shortcut. b2676f1
- [x] Task: Conductor - User Manual Verification 'Implementation' (Protocol in workflow.md) b2676f1

## Phase 3: Verification
- [x] Task: Confirm `c` triggers the composer. 96abe46
