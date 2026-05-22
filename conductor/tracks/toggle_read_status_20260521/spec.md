# Specification: Native Toggle Read Status

## Overview
Implement a native Go keybinding (`u`) to toggle the read/unread status of the selected email directly in the TUI. This replaces the broken `toggle_read.lua` plugin.

## Functional Requirements
- **Native Keybinding**: Map the `u` key in the inbox and email view contexts to a `toggle_read` action.
- **Go Backend Integration**: Implement the logic to check the current status of an email and flip it (Mark Read <-> Mark Unread).
- **UI Feedback**: Immediately update the visual state (e.g., bold/unbold) in the message list upon toggling.

## Acceptance Criteria
- [ ] Pressing `u` in the inbox toggles the read status of the highlighted email.
- [ ] Pressing `u` while viewing an email toggles its status and updates the UI.
- [ ] The action is configurable in `keybinds.json`.
