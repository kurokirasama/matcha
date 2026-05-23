# Implementation Plan: Layout Quick Toggle

#### Phase 1: Configuration & Settings Integration
- [x] Task: Add `EnableQuickToggle` boolean field to the `Config` struct in `config/config.go`. 7a8dd62
- [x] Task: Update serialization logic in `SaveConfig` and `LoadConfig` in `config/config.go` to handle the new field. f93466c
- [x] Task: Update the settings menu in `tui/settings_general.go` to include the "Layout Quick Toggle" option. 923ac91
- [x] Task: Implement logic in `tui/settings_general.go` to disable/force-off "Layout Quick Toggle" when "Split View" is set to Horizontal. 7c5caea
- [x] Task: Conductor - User Manual Verification 'Configuration Safety' (Protocol in workflow.md)
[checkpoint: e9db627]

#### Phase 2: Keybinding & TUI Logic
- [x] Task: Define the `Shift+L` keybinding in `tui/keys.go` (if not already present). 47a65d3
- [x] Task: Modify `FolderInbox.Update` in `tui/folder_inbox.go` to handle the `Shift+L` message. bd735a4
- [x] Task: Implement the layout cycling logic in `FolderInbox` based on the current `m.layout` and `m.config.EnableQuickToggle`. bd735a4
- [x] Task: Add a new message type `ToggleLayoutMsg` (or similar) to notify the main model when the layout is changed via shortcut. bd735a4
- [x] Task: Update `main.go` to handle layout changes from `FolderInbox`, update the config, and save to disk. bd735a4
- [x] Task: Conductor - User Manual Verification 'Shortcut Interaction' (Protocol in workflow.md)
[checkpoint: a6957c0]

#### Phase 3: Validation & Polish
- [ ] Task: Add unit tests in `tui/folder_inbox_test.go` to verify `Shift+L` behavior in different modes.
- [ ] Task: Add unit tests in `config/config_test.go` for the new config field persistence.
- [ ] Task: Implement the feedback notification when attempting to enable the toggle in Horizontal mode.
- [ ] Task: Run `make build` and perform end-to-end manual verification of all layout modes and the toggle.
- [ ] Task: Conductor - User Manual Verification 'Final Quality Gate' (Protocol in workflow.md)
