# Implementation Plan: Fix Missing Status Bar in Split Views

## Phase 1: Investigation & Test Setup
- [x] Task: Write failing UI unit tests simulating email preview focus in `LayoutVertical` and `LayoutHorizontal` to assert the presence of the status bar. 66ca4fe
- [x] Task: Investigate the `tui/` rendering logic (likely in `main.go`, `view.go`, or specific split view files) to identify why the status bar is omitted when an email is previewed. 66ca4fe
- [x] Task: Conductor - User Manual Verification 'Investigation & Test Setup' (Protocol in workflow.md) 66ca4fe

## Phase 2: Implementation
- [x] Task: Implement the fix to ensure the status bar is rendered in both split view modes during email preview. 66ca4fe
- [x] Task: Ensure the dynamic context logic passes the correct active pane state to the status bar renderer, displaying email keybinds when the preview is focused, and inbox keybinds when the list is focused. 66ca4fe
- [x] Task: Verify unit tests pass. 66ca4fe
- [x] Task: Conductor - User Manual Verification 'Implementation' (Protocol in workflow.md) 66ca4fe

## Phase 3: Final Verification
- [x] Task: Run `make build` and visually confirm the status bar appears with the correct contextual keybinds in both `Vertical` and `Horizontal` split modes. 66ca4fe
- [x] Task: Conductor - User Manual Verification 'Final Quality Gate' (Protocol in workflow.md) 66ca4fe