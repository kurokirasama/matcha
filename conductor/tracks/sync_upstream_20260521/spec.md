# Specification: Synchronize Fork with Upstream

## Overview
Synchronize the local fork (on GitHub and local machine) with the official upstream repository (`floatpane/matcha`). This involves fetching upstream changes, merging them into the `master` and `private` branches, and ensuring our custom native features (like the toggle read status) are preserved and functional.

## Functional Requirements
- **Upstream Synchronization**:
    - Fetch the latest changes from the `upstream` remote.
    - Merge `upstream/master` into the local `master` branch.
    - Push the updated `master` branch to the `origin` fork on GitHub.
- **Private Branch Update**:
    - Merge the updated `master` branch into the local `private` branch.
- **Conflict Management**:
    - Resolve merge conflicts favoring `upstream` changes where appropriate, but ensuring our custom application logic is maintained.
    - Perform a manual review of all merge results.
- **Verification**:
    - Execute `make build` and `make test` to ensure stability and compatibility.
    - Manually verify that the native "Toggle Read Status" feature remains functional.

## Acceptance Criteria
- [ ] Local `master` and `origin/master` are in sync with `upstream/master`.
- [ ] Local `private` branch contains all new upstream changes.
- [ ] The application builds successfully (`make build`).
- [ ] The full test suite passes (`make test`).
- [ ] Custom native features are verified as functional.

## Out of Scope
- Implementing new features during the sync process.
- Drastic architectural refactoring to match upstream (unless required for compilation).
