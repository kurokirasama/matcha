# Implementation Plan: Native Folder Panel Toggle

## Phase 1: State & Configuration
- [ ] Task: Add `hideSidebar bool` to the `FolderInbox` struct in `tui/folder_inbox.go`.
- [ ] Task: Define `ToggleSidebar` action in `config/keybinds.go` and map it to `F` in `default_keybinds.json`. (Keep `f` as `filter`).
- [ ] Task: Conductor - User Manual Verification 'State Setup' (Protocol in workflow.md)

## Phase 2: Layout Logic
- [ ] Task: Modify `View()` in `tui/folder_inbox.go` to conditionally render the sidebar.
- [ ] Task: Update width calculation logic to account for hidden sidebar.
- [ ] Task: Conductor - User Manual Verification 'Layout Implementation' (Protocol in workflow.md)

## Phase 3: Verification
- [ ] Task: Verify toggle works smoothly without visual glitches.
- [ ] Task: Verify resize behavior on terminal window resize while panel is hidden.
