# Specification: Interactive Send Blocking Hooks

## Overview
Enhance the plugin system to allow Lua scripts to intercept and potentially block the email sending process (e.g., for attachment reminders or subject checks).

## Functional Requirements
- **Blocking Hooks**: Update `CallSendHook` to respect a boolean return value from Lua.
- **Interactivity**: Allow plugins to halt the send process and return focus to the composer with a notification.
- **Metadata Exposure**: Pass full email metadata (body, attachments) to the hook.

## Acceptance Criteria
- [ ] A Lua plugin returning `false` from `email_send_before` prevents the email from being sent.
- [ ] Plugins have access to attachment list and body.
