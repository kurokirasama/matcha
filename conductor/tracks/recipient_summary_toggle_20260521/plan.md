# Implementation Plan: Recipient Summary & Toggle

## Phase 1: State & Helper
- [ ] Task: Add `expandedRecipients bool` to `EmailView` in `tui/email_view.go`.
- [ ] Task: Implement `formatRecipients` helper function.
- [ ] Task: Conductor - User Manual Verification 'Helpers' (Protocol in workflow.md)

## Phase 2: UI Implementation
- [ ] Task: Update `View()` to use the summary/full logic.
- [ ] Task: Handle `v` keybind to toggle state and recalculate layout.
- [ ] Task: Conductor - User Manual Verification 'UI Implementation' (Protocol in workflow.md)

## Phase 3: Verification
- [ ] Task: Test with various numbers of recipients.
- [ ] Task: Verify viewport scrolling works correctly after expansion.
