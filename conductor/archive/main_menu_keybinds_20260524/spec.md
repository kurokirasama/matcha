# Track Specification

## Overview
Add quick access keybinds to the main screen (Inbox) to easily navigate to primary application views. These keybinds are controlled by a new configuration setting and will dynamically update the help bar when enabled.

## Functional Requirements
- **Configuration Setting**:
  - Add a new boolean setting named `EnableMainMenuKeybinds` to the configuration.
  - The default value must be `false` (disabled).
  - Add a toggle for this setting in the General Settings UI, labeled exactly as: `"Enable Main Menu Keybinds"`.
- **Keybinds**:
  - When `EnableMainMenuKeybinds` is enabled, the following single-key shortcuts must be active *only* when the Inbox list is focused:
    - `v`: View email (equivalent to entering the email preview/view mode).
    - `c`: Compose a new email.
    - `p`: Open the Plugin Marketplace.
    - `s`: Open Settings.
- **UI Feedback**:
  - When `EnableMainMenuKeybinds` is true, the bottom help bar on the main screen must display the active shortcuts: `v: view • c: compose • p: plugins • s: settings`.
  - When false, these specific hints should be omitted.

## Acceptance Criteria
- [ ] `EnableMainMenuKeybinds` setting exists, defaults to `false`, and can be toggled via the General Settings UI using the label "Enable Main Menu Keybinds".
- [ ] Pressing `v`, `c`, `p`, or `s` in the Inbox view executes the corresponding action if the setting is enabled.
- [ ] The keybinds do not trigger these actions if the setting is disabled, or if the user is focused on a different pane (e.g., the folder sidebar or composer).
- [ ] The help bar dynamically displays the shortcuts only when the setting is enabled.

## Out of Scope
- Making these keybinds globally accessible across all screens.
- Modifying the actual functionality of the target views (Compose, Settings, etc.).