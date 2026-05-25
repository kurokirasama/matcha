# Track Specification

## Overview
Currently, when previewing an email in horizontal or vertical split view modes, the entire bottom status bar (which normally displays keybinds like reply, forward, delete, etc.) fails to render. This track will investigate and resolve the issue so that the status bar correctly appears in all layout modes.

## Functional Requirements
- **Status Bar Rendering**:
    - The bottom status bar must render correctly when an email is being previewed in `LayoutVertical` and `LayoutHorizontal` modes.
- **Contextual Keybinds**:
    - The status bar must dynamically display the contextual keybinds relevant to the currently focused pane.
    - When the email preview pane is focused, it should display email-specific actions (e.g., `r: reply`, `f: forward`, `d: delete`, `a: archive`, `esc: back`).
    - When the inbox list pane is focused, it should display inbox-specific actions.

## Acceptance Criteria
- [ ] The status bar is visible at the bottom of the screen when previewing an email in vertical split view.
- [ ] The status bar is visible at the bottom of the screen when previewing an email in horizontal split view.
- [ ] When the email preview is focused in split view, the status bar displays email-specific keybinds.
- [ ] When the focus switches back to the inbox list in split view, the status bar updates to show inbox-specific keybinds.

## Out of Scope
- Adding new keybinds or actions.
- Redesigning the status bar UI component itself.