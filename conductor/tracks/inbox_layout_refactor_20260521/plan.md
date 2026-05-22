# Implementation Plan: Inbox Layout Refactor & Horizontal Preview

## Phase 1: Layout State
- [ ] Task: Define `LayoutMode` enum and add it to `FolderInbox` state.
- [ ] Task: Add `CycleLayout` action to `config/keybinds.go` and map to `L`.
- [ ] Task: Conductor - User Manual Verification 'State Design' (Protocol in workflow.md)

## Phase 2: Rendering Refactor
- [ ] Task: Modify `View()` in `tui/folder_inbox.go` to support vertical/horizontal/none split logic.
- [ ] Task: Update `SetSize` logic to calculate dimensions dynamically based on mode.
- [ ] Task: Conductor - User Manual Verification 'Rendering Implementation' (Protocol in workflow.md)

## Phase 3: Verification
- [ ] Task: Verify all 3 modes work as expected.
- [ ] Task: Check layout integrity on window resize.
