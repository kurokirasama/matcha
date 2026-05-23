# Specification: Layout Quick Toggle

**Overview**
This track implements a "Layout Quick Toggle" feature, allowing users to quickly cycle between different email preview layouts (split view vs. full-screen list) using a keyboard shortcut. It also adds a new setting to enable or disable this shortcut.

**Functional Requirements**
- **New Setting: "Enable Layout Quick Toggle"**:
    - Add a boolean setting in the "General" settings menu.
    - Labeled as "Layout Quick Toggle".
    - Default: `false` (Disabled).
- **Keyboard Shortcut: `Shift+L`**:
    - This shortcut is active only when the "Layout Quick Toggle" setting is `true`.
- **Context-Aware Behavior**:
    - **Preview Mode: OFF**:
        - `Shift+L` cycles between:
            1.  Inbox list uses **all** the screen height.
            2.  Inbox list uses **half** the screen height (matching the space it would take in horizontal mode, but with no preview pane active).
    - **Preview Mode: Vertical**:
        - `Shift+L` cycles between:
            1.  Inbox list uses **all** the screen width (preview pane inactive/hidden).
            2.  Inbox list uses **half** the screen width (preview pane active on the right).
    - **Preview Mode: Horizontal**:
        - The Quick Toggle setting must be **disabled** and cannot be turned on.
        - If the toggle was previously on and the user switches to Horizontal mode, it must be automatically and silently turned **OFF**.
- **State Persistence**:
    - When the layout is cycled via `Shift+L`, the main "Split View" setting and the `config.json` must be updated to reflect the new state.
- **Feedback**:
    - When in Horizontal mode, if the user attempts to enable the toggle, a brief notification or hint should explain that it's unavailable in this mode.

**Acceptance Criteria**
- [ ] New "Layout Quick Toggle" setting appears in General settings.
- [ ] `Shift+L` works correctly in "OFF" and "Vertical" preview modes when enabled.
- [ ] `Shift+L` has no effect when disabled.
- [ ] Horizontal mode correctly forces the toggle to `OFF`.
- [ ] Cycling the layout via shortcut updates the menu setting and persists to disk.
