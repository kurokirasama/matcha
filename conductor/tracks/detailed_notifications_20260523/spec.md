# Specification: Detailed Email Notifications

## Overview
This track introduces an "Expanded Notifications" setting. When enabled, desktop notifications for new emails will display the sender and subject instead of a generic "New email received" message. This provides more context without opening the application, while remaining an opt-in feature for privacy reasons.

## Functional Requirements
- **Settings Toggle**:
    - Add a boolean setting `EnableExpandedNotifications` to the application configuration.
    - Default value: `false` (Disabled by default for privacy).
    - Add a corresponding toggle labeled "Expanded Notifications" in the General Settings menu.
    - Implementation must be strictly additive, preserving all existing config fields.
- **Notification Formatting (`daemon` / `notify`)**:
    - When `EnableExpandedNotifications` is true, the OS notification should be formatted as follows:
        - **Title**: `New email from <Sender>`
        - **Body**: `<Subject>`
    - If `EnableExpandedNotifications` is false, the application must retain the current generic notification behavior.

## Acceptance Criteria
- [ ] New `enable_expanded_notifications` boolean added to configuration and persisted correctly.
- [ ] Setting is available in the General Settings UI with proper localizations.
- [ ] When enabled, new email notifications display the sender in the title and the subject in the body.
- [ ] When disabled, new email notifications display the legacy generic message.
- [ ] The setting defaults to `false`.