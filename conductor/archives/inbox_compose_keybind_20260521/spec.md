# Specification: Native Inbox Compose Keybind

## Overview
Adds a native `c` keybinding to the inbox view to trigger the "Compose" action, providing a faster and more intuitive workflow for starting new emails.

## Functional Requirements
- **Keybinding**: Map `c` in the inbox view to the compose action.
- **View Transition**: Ensure the TUI correctly transitions to the composer view when the key is pressed.

## Acceptance Criteria
- [ ] Pressing `c` in the inbox opens the composer.
- [ ] Help bar shows `c: compose`.
