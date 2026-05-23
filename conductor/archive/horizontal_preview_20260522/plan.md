# Implementation Plan: Horizontal Email Preview

## Phase 1: Robust Configuration Extension
- [x] Task: Define `LayoutMode` type and constants (Off, Vertical, Horizontal) in `config/config.go`. 669849c
- [x] Task: Add `Layout` field to `Config` struct using a strictly additive approach. ca4ee5b
- [x] Task: Update serialization logic to ensure all existing fields are preserved during save/load. 1ec68d7
- [x] Task: Add regression tests for config persistence to prevent data loss. e17e428
- [x] Conductor - User Manual Verification 'Config Safety' (Protocol in workflow.md)
[checkpoint: 77e0ecf]

## Phase 2: TUI Refactor & Layout Engine
- [x] Task: Modify `FolderInbox` model in `tui/folder_inbox.go` to handle dynamic split orientation. 8ad71eb
- [x] Task: Update `calculateInboxWidth`, `calculateInboxHeight`, `calculatePreviewWidth`, and `calculatePreviewHeight` in `tui/folder_inbox.go` to be layout-aware. ca5aeff
- [x] Task: Add `rowOffset` and `columnOffset` support to `EmailView` in `tui/email_view.go` for proper image rendering in both horizontal and vertical modes. 265de84
- [x] Task: Write unit tests in `tui/folder_inbox_test.go` to verify window resizing. f318dc3
- [x] Conductor - User Manual Verification 'Layout Rendering' (Protocol in workflow.md)
[checkpoint: c2d0640]

## Phase 3: Settings Menu Integration
- [x] Task: Update the settings TUI in `tui/settings_general.go` to include the "Split View" choice menu, with localized labels. 84584ea
- [x] Task: Implement the message handler in `main.go` to apply layout changes immediately across the application when `ConfigSavedMsg` is received. dd6bea0
- [x] Conductor - User Manual Verification 'Settings Interaction' (Protocol in workflow.md)
[checkpoint: af5d053]

## Phase 4: Verification & Final Polish
- [x] Task: Run `make build` and perform end-to-end manual verification of all layout modes. af5d053
- [x] Conductor - User Manual Verification 'Final Quality Gate' (Protocol in workflow.md)
[checkpoint: 2ff2eb0]

## Phase 5: Layout Refinement & Bug Fixes
- [x] Task: Fix `Inbox.SetSize` to use full height instead of hardcoded 50%.
- [x] Task: Refactor `FolderInbox.View` to keep sidebar full-height in horizontal mode.
- [x] Task: Correct `columnOffset` for images in horizontal preview mode.
- [x] Task: Conductor - User Manual Verification 'Refined Layout' (Protocol in workflow.md)

## Phase: Review Fixes
- [x] Task: Apply review suggestions b82b9d3