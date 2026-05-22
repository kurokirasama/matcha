# Implementation Plan: Fuzzy Folder Jump (Native)

## Phase 1: UI Component
- [ ] Task: Create or adapt a searchable list model for the TUI.
- [ ] Task: Implement `matcha.FolderPicker` (internal Go component).
- [ ] Task: Conductor - User Manual Verification 'Component Design' (Protocol in workflow.md)

## Phase 2: Integration
- [ ] Task: Add `FolderJump` action to `config/keybinds.go` and map to `g`.
- [ ] Task: Implement the switch logic in `tui/folder_inbox.go`.
- [ ] Task: Conductor - User Manual Verification 'Integration' (Protocol in workflow.md)

## Phase 3: Verification
- [ ] Task: Verify fuzzy filtering works.
- [ ] Task: Verify successful navigation to selected folder.
