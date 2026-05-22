# Implementation Plan: Interactive Send Blocking Hooks

## Phase 1: Hook Refactor
- [ ] Task: Modify `plugin/hooks.go` to handle boolean return values.
- [ ] Task: Expand the metadata passed to the Lua VM in the send hook.
- [ ] Task: Conductor - User Manual Verification 'Hook Logic' (Protocol in workflow.md)

## Phase 2: Main Loop Integration
- [ ] Task: Update `main.go` to abort the send sequence if the hook returns `false`.
- [ ] Task: Ensure appropriate error/notification messages are shown.
- [ ] Task: Conductor - User Manual Verification 'Integration' (Protocol in workflow.md)

## Phase 3: Verification
- [ ] Task: Verify an "Attachment Reminder" plugin can successfully block a send.
