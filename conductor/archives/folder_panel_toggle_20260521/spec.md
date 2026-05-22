# Specification: Native Folder Panel Toggle

## Overview
Implement a single-key toggle (`F`) to show/hide the left folder panel in the Matcha TUI. This allows users to maximize the space available for the message list and preview.

## Functional Requirements
- **Toggle Keybinding**: Map the `F` key to toggle the sidebar. Keep the existing `f` keybinding for the `filter` action.
- **Layout Persistence**: The TUI should remember the panel visibility state during the session.
- **Responsive Resize**: Hiding the panel should immediately expand the message list to fill the available width.

## Acceptance Criteria
- [ ] Pressing `F` in the inbox toggles the folder panel.
- [ ] Message list expands to full width when the panel is hidden.
- [ ] The keybinding is user-configurable.
