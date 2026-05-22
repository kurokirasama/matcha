# Implementation Plan: Synchronize Fork with Upstream

## Phase 1: Upstream Sync (Master) [checkpoint: d075955]
- [x] Task: Fetch latest changes from `upstream` remote.
- [x] Task: Switch to local `master` branch.
- [x] Task: Merge `upstream/master` into local `master`.
- [x] Task: Push updated `master` to `origin/master`.
- [x] Task: Conductor - User Manual Verification 'Master Sync' (Protocol in workflow.md)

## Phase 2: Private Branch Update [checkpoint: 510dfb6]
- [x] Task: Switch to local `private` branch.
- [x] Task: Merge local `master` into `private`.
- [x] Task: Resolve any conflicts in `private` (e.g., in `todos.md` or `conductor/` if accidentally touched).
- [x] Task: Conductor - User Manual Verification 'Private Sync' (Protocol in workflow.md)

## Phase 3: Verification & Compatibility [checkpoint: 510dfb6]
- [x] Task: Verify application builds successfully (`make build`). 20d031c
- [x] Task: Run full test suite (`make test`). 20d031c
- [x] Task: Manually verify the "Native Toggle Read Status" feature (`bin/matcha`). 20d031c
- [x] Task: Conductor - User Manual Verification 'Final Verification' (Protocol in workflow.md) 20d031c
