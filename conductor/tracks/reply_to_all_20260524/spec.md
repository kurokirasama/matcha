# Track Specification

## Overview
Add a "Reply to All" keybind (`A` / `shift+a`) to the email visualization screen. This feature allows users to quickly reply to the sender and all other recipients of an email. The functionality is toggleable via a new configuration setting and includes UI feedback in the status bar when enabled.

## Functional Requirements
- **Keybind**:
  - Add a new keybind `A` (`shift+a`) to the email visualization screen.
  - Pressing `A` opens the composer with "Reply to All" logic applied.
- **Recipient Population**:
  - The "Reply to All" logic must populate the composer's fields as follows:
    - **To**: Original Sender + all original `To` recipients (excluding the current user's email).
    - **Cc**: All original `Cc` recipients (excluding the current user's email).
- **Configuration**:
  - Add a new boolean setting named `Enable Reply to All` (`EnableReplyToAll` in config structs).
  - Default state: `false` (Disabled).
  - Add a toggle for this setting in the General Settings UI.
- **UI Feedback**:
  - When `Enable Reply to All` is enabled, the bottom help/status bar in the email visualization screen must display the `A: reply all` hint alongside existing hints.
  - The hint should not be visible when the setting is disabled.

## Acceptance Criteria
- [ ] The `Enable Reply to All` setting is available in the configuration and General Settings UI, defaulting to `false`.
- [ ] Pressing `A` in the email visualization screen opens the composer with "Reply to All" behavior only when the setting is enabled.
- [ ] The "Reply to All" composer correctly populates `To` and `Cc` fields, excluding the user's own email address.
- [ ] The help menu in the email visualization screen dynamically shows the `A: reply all` hint when the setting is enabled.

## Out of Scope
- Modifying the behavior of the standard `r` (reply) keybind.
- Implementing "Reply to All" from the inbox list view.