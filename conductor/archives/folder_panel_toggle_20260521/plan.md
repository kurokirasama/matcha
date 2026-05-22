# Implementation Plan: Native Folder Panel Toggle

## Phase 1: State & Configuration [checkpoint: a777b31]
- [x] Task: Add `hideSidebar bool` to the `FolderInbox` struct in `tui/folder_inbox.go`.
- [x] Task: Define `ToggleSidebar` action in `config/keybinds.go` and map it to `F` in `default_keybinds.json`. (Keep `f` as `filter`).
- [x] Task: Conductor - User Manual Verification 'State Setup' (Protocol in workflow.md)

## Phase 2: Layout Logic [checkpoint: a777b31]
- [x] Task: Modify `View()` in `tui/folder_inbox.go` to conditionally render the sidebar.
- [x] Task: Update width calculation logic to account for hidden sidebar.
- [x] Task: Conductor - User Manual Verification 'Layout Implementation' (Protocol in workflow.md)

## Phase 3: Verification [checkpoint: a777b31]
- [x] Task: Verify toggle works smoothly without visual glitches.
- [x] Task: Verify resize behavior on terminal window resize while panel is hidden.
