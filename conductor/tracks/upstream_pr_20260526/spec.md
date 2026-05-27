# Track Specification: Upstream Sync and PR

## Overview
Synchronize the local fork with the official upstream repository (`upstream/master`), resolve any merge conflicts while preserving new local features, ensure the GitHub fork (`origin`) is up to date, and create a Pull Request to the official repository.

## Functional Requirements
1. **Upstream Sync**: Fetch the latest changes from `upstream` (`floatpane/matcha`) and merge them into the local public feature branch (e.g., `master`).
2. **Conflict Resolution**: In the event of conflicts, accept upstream changes for existing files but ensure that the new custom features remain perfectly intact.
3. **Remote Sync**: Push the updated, synchronized local branch to the GitHub fork (`origin`) to ensure local and remote are identical.
4. **Pull Request Creation**: Open a Pull Request against the official upstream repository. 
5. **PR Description Content**:
   - Express appreciation for the app.
   - Explain that the new features address workflow needs and aim to help other users.
   - Specifically mention the `u` keybind feature: state that the original plugin did not work, and since modifying it wasn't possible, it was implemented directly.

## Non-Functional Requirements
- Maintain adherence to the Conductor Dual-Branch Architecture: ensure no private/conductor tracking files from the `private` branch leak into the public PR branch.