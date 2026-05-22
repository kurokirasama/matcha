# Specification: Native Folder Panel Toggle

## Overview
Implement a single-key toggle (`f`) to show/hide the left folder panel in the Matcha TUI. This allows users to maximize the space available for the message list and preview.

## Functional Requirements
- **Toggle Keybinding**: Map the `f` key (moving existing `filter` if necessary, or choosing a conflict-free key if preferred) to toggle the sidebar. Note: `f` is currently `filter` in inbox, we will move filter to `F` or `/`.
- **Layout Persistence**: The TUI should remember the panel visibility state during the session.
- **Responsive Resize**: Hiding the panel should immediately expand the message list to fill the available width.

## Acceptance Criteria
- [ ] Pressing `f` in the inbox toggles the folder panel.
- [ ] Message list expands to full width when the panel is hidden.
- [ ] The keybinding is user-configurable.
